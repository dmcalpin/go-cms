package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
)

var client *datastore.Client

func init() {
	ctx := context.Background()

	var err error

	// set DATASTORE_PROJECT_ID env var
	client, err = datastore.NewClient(ctx, "")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	if client == nil {
		log.Fatalf("no client")
	} else {
		fmt.Println("client made")
	}
}

func GetClient() *datastore.Client {
	return client
}
