package main

import (
	"log"

	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}

	c.ExtractPipeline()


	// Create a repo called "projects."
	if err := c.CreateRepo("projects"); err != nil {
		log.Fatal(err)
	}
}
