package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/samkreter/KubeconAsia2018/pachbot/deploy"

	"github.com/gorilla/mux"
)

type ACRResponse struct {
	ID        string     `json:"id"`
	Timestamp string     `json:"timestamp"`
	Action    string     `json:"action"`
	Target    ACRTarget  `json:"target"`
	Request   ACRRequest `json:"request"`
}

type ACRRequest struct {
	ID        string `json:"id"`
	Host      string `json:"host"`
	Method    string `json:"method"`
	Useragent string `json:"useragent"`
}

type ACRTarget struct {
	MediaType  string `json:"mediaType"`
	Size       int64  `json:"size"`
	Digest     string `json:"digest"`
	Length     int64  `json:"length"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	githubPAT := getEnv("GITHUB_TOKEN")
	cloneURL := getEnv("CLONE_URL")
	workDir := getEnvWithDefault("WORK_DIR", "./workrepo")

	port := getEnvWithDefault("SERVER_PORT", "8080")

	githubClient, err := github.NewClient(cloneURL, githubPAT, workDir)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/acrpush", acrPushHandler(githubClient)).
		Methods("POST")

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Listening on: %s", addr)

	log.Fatal(srv.ListenAndServe())
}

// func test() {
// 	cmd := exec.Command("pachctl", "list-repo")
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("test: %s", string(out))
// }

type JSONResp struct {
	message    string `json"message"`
	statusCode int    `json"statusCode"`
}

func acrPushHandler(deployer *deploy.Deployer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var acrResp ACRResponse
		err := json.NewDecoder(r.Body).Decode(&acrResp)
		if err != nil {
			log.Printf("ERROR: acrPush: %v", err)
			return
		}

		msg := fmt.Sprintf("Deploying new image to staging: %s/%s", acrResp.Target.Repository, acrResp.Target.Tag)

		log.Println(msg)

		go deployer.NewDeployment(acrResp.Target.Repository, acrResp.Target.Tag)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(JSONResp{
			message:    msg,
			statusCode: http.StatusOK,
		})
	}
}

func getEnv(envName string) string {
	val, ok := os.LookupEnv(envName)
	if !ok {
		log.Fatalf("%s must be set", envName)
	}
	return val
}

func getEnvWithDefault(envName, defaultVal string) string {
	val, ok := os.LookupEnv(envName)
	if !ok {
		return defaultVal
	}
	return val
}
