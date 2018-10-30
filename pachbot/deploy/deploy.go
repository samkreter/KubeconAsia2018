package deploy

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	git "gopkg.in/src-d/go-git.v4"
)

// Deployer holds the needed data for a deployer
type Deployer struct {
	c        *github.Client
	workdir  string
	cloneURL string
}

// New creates a new deployer
func New(cloneURL, githubPAT, workdir string) (*Deployer, error) {
	if workdir == "" {
		return nil, fmt.Errorf("workdir can not be empty")
	}

	if cloneURL == "" {
		return nil, fmt.Errorf("cloneURL can not be empty")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubPAT},
	)
	tc := oauth2.NewClient(ctx, ts)

	c := github.NewClient(tc)

	return &Deployer{
		c:        c,
		cloneURL: cloneURL,
		workdir:  workdir,
	}, nil
}

// NewDeployment creates a new deployment in the cluster
func (d *Deployer) NewDeployment(imageName, tag string) {
	if err := d.updateRepo(); err != nil {
		log.Printf("ERROR: failed to update repository with error: %v", err)
	}
}

// UpdateRepo clones a repo or pulls latest changes if it doesn't exist
func (d *Deployer) updateRepo() error {
	log.Printf("Cloning Repository %s into folder: %s", d.cloneURL, d.workdir)
	_, err := git.PlainClone(d.workdir, false, &git.CloneOptions{
		URL:      d.cloneURL,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Printf("%v: Attempting to pull latest changes", err)
		return d.pullLatest()
	}

	return nil
}

// PullLatest Pulls the latest changes from a repository
func (d *Deployer) pullLatest() error {
	r, err := git.PlainOpen(d.workdir)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin", SingleBranch: true})
	if err != nil {
		return err
	}

	return nil
}
