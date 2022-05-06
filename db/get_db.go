package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
)

var Client *datastore.Client

func init() {
	ctx := context.Background()

	var err error

	// set DATASTORE_PROJECT_ID env var
	Client, err = datastore.NewClient(ctx, "")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	if Client == nil {
		log.Fatalf("no client")
	} else {
		fmt.Println("client made")
	}
}
