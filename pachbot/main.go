package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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

func acrPushHandler(w http.ResponseWriter, r *http.Request) {
	var resp ACRResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		log.Printf("ERROR: acrPush: %v", err)
		return
	}

	log.Printf("Got request: %s:%s", resp.Target.Repository, resp.Target.Tag)

	json.NewEncoder(w).Encode(http.StatusOK)
}

func main() {
	// acrIdentity, ok := os.LookupEnv("ACR_IDENTITY")
	// if !ok {
	// 	log.Println("WARN: ACR_IDENTITY not set.")
	// }

	r := mux.NewRouter()
	r.HandleFunc("/acrpush", acrPushHandler).
		Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Listening on port: %d", 8080)

	log.Fatal(srv.ListenAndServe())
}
