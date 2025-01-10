package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/buraksezer/olric"
	"github.com/buraksezer/olric/config"
)

func main() {
	// Sample for Olric v0.5.x

	// Deployment scenario: embedded-member
	// This creates a single-node Olric cluster. It's good enough for experimenting.

	// config.New returns a new config.Config with sane defaults. Available values for env:
	// local, lan, wan
	c := config.New("local")

	// Callback function. It's called when this node is ready to accept connections.
	ctx, cancel := context.WithCancel(context.Background())
	c.Started = func() {
		defer cancel()
		log.Println("[INFO] Olric is ready to accept connections")
	}

	// Create a new Olric instance.
	db, err := olric.New(c)
	if err != nil {
		log.Fatalf("Failed to create Olric instance: %v", err)
	}

	// Start the instance. It will form a single-node cluster.
	go func() {
		// Call Start at background. It's a blocker call.
		err = db.Start()
		if err != nil {
			log.Fatalf("olric.Start returned an error: %v", err)
		}
	}()

	<-ctx.Done()

	// In embedded-member scenario, you can use the EmbeddedClient. It implements
	// the Client interface.
	e := db.NewEmbeddedClient()

	dm, err := e.NewDMap("bucket-of-arbitrary-items")
	if err != nil {
		log.Fatalf("olric.NewDMap returned an error: %v", err)
	}

	ctx, cancel = context.WithCancel(context.Background())

	// Magic starts here!
	fmt.Println("##")
	fmt.Println("Simple Put/Get on a DMap instance:")
	err = dm.Put(ctx, "my-key", "Olric Rocks!")
	if err != nil {
		log.Fatalf("Failed to call Put: %v", err)
	}

	gr, err := dm.Get(ctx, "my-key")
	if err != nil {
		log.Fatalf("Failed to call Get: %v", err)
	}

	// Olric uses the Redis serialization format.
	value, err := gr.String()
	if err != nil {
		log.Fatalf("Failed to read Get response: %v", err)
	}

	fmt.Println("Response for my-key:", value)
	fmt.Println("##")

	// Don't forget the call Shutdown when you want to leave the cluster.
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.Shutdown(ctx)
	if err != nil {
		log.Printf("Failed to shutdown Olric: %v", err)
	}
}
