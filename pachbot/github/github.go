package github

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	git "gopkg.in/src-d/go-git.v4"
)

const (
	defaultWorkRepo = "./workrepo"
)

// Client holds the needed info for a git client
type Client struct {
	c        *github.Client
	workdir  string
	cloneURL string
}

// NewClient creates a new git client
func NewClient(cloneURL, githubPAT, workdir string) (*Client, error) {
	if workdir == "" {
		workdir = defaultWorkRepo
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubPAT},
	)
	tc := oauth2.NewClient(ctx, ts)

	c := github.NewClient(tc)

	return &Client{
		c:        c,
		cloneURL: cloneURL,
		workdir:  workdir,
	}, nil
}

// UpdateRepo clones a repo or pulls latest changes if it doesn't exist
func (c *Client) UpdateRepo() error {
	log.Printf("Cloning Repository %s into folder: %s", c.cloneURL, c.workdir)
	_, err := git.PlainClone(c.workdir, false, &git.CloneOptions{
		URL:      c.cloneURL,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Printf("%v: Attempting to pull latest changes", err)
		return c.PullLatest()
	}

	return nil
}

// PullLatest Pulls the latest changes from a repository
func (c *Client) PullLatest() error {
	r, err := git.PlainOpen(c.workdir)
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
