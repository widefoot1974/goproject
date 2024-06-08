package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
		log.Printf("INFO: NATS_URL not set, using default [%v]\n", url)
	}
	log.Printf("INFO: Connecting to NATS server at [%v]\n", url)

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to connect to NATS: %v\n", err)
	}
	defer nc.Drain()
	log.Println("INFO: Successfully connected to NATS")

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("ERROR: Failed to initialize JetStream context: %v\n", err)
	}
	log.Println("INFO: JetStream context initialized")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kv, err := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{Bucket: "profiles"})
	if err != nil {
		log.Fatalf("ERROR: Failed to create KeyValue store: %v\n", err)
	}
	log.Println("INFO: KeyValue store created")

	w, err := kv.Watch(ctx, "sue.*")
	if err != nil {
		log.Fatalf("ERROR: Failed to start watcher: %v\n", err)
	}
	log.Println("INFO: Watcher started for pattern 'sue.*'")
	go watchUpdates(w)

	time.Sleep(time.Second * 3) // Delay after Watch

	// Perform color updates with delays
	updateColor(ctx, kv, "blue")
	time.Sleep(time.Second * 3) // Delay after blue
	updateColor(ctx, kv, "green")
	time.Sleep(time.Second * 3) // Delay after green
	updateColor(ctx, kv, "red")
	time.Sleep(time.Second * 3) // Final delay after red

	log.Println("Process completed.")
}

func updateColor(ctx context.Context, kv jetstream.KeyValue, color string) {
	log.Printf("INFO: Updating sue.color to %q\n", color)
	_, err := kv.Put(ctx, "sue.color", []byte(color))
	if err != nil {
		log.Fatalf("ERROR: Failed to update color: %v\n", err)
	}

	entry, err := kv.Get(ctx, "sue.color")
	if err != nil {
		log.Fatalf("ERROR: Failed to retrieve color: %v\n", err)
	}
	if entry != nil {
		log.Printf("INFO: %s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))
	} else {
		log.Println("ERROR: Entry is nil after update")
	}
}

func watchUpdates(w jetstream.KeyWatcher) {
	for kve := range w.Updates() {
		if kve == nil {
			log.Println("ERROR: [Watcher] Received nil kve from the watcher")
			continue
		}
		log.Printf("INFO: [Watcher] %s @ %d -> %q (op: %s)\n", kve.Key(), kve.Revision(), string(kve.Value()), kve.Operation())
	}
	log.Println("INFO: [Watcher] Watcher channel closed")
}
