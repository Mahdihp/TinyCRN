package main

import (
	"TinyCRM/ent"
	"context"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	client, err := ent.Open("postgres", "host=localhost port=5432 user=root dbname=TinyCRMDB password=123456 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
