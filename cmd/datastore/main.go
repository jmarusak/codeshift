package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"cloud.google.com/go/datastore"
	"os"
)

type Counter struct {
	Visits int
	LastVisit time.Time
}

func main() {
	ctx := context.Background()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	fmt.Println("Google Cloud Project ID:", projectID)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal("Datastore - Failed to create client %v", err)
	}
	defer client.Close()

	kind := "Counter"
	name := "counter"
	counterKey := datastore.NameKey(kind, name, nil)

	var counter Counter
	if err := client.Get(ctx, counterKey, &counter); err !=nil && err != datastore.ErrNoSuchEntity {	
		log.Fatalf("Datastore - Failed to get Counter: %v", err)
	}
	fmt.Println(counter)

    counter.Visits += 1
	counter.LastVisit = time.Now()

	if _, err := client.Put(ctx, counterKey, &counter); err != nil {
         log.Fatalf("Datastore - Failed to save Counter: %v", err)
	}

	fmt.Println("Done.")
}
