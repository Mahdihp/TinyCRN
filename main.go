package main

import (
	"TinyCRM/config"
	"TinyCRM/ent"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	appConfig := config.Initialize()

	ConnetionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		appConfig.DbConfig.Host, appConfig.DbConfig.DbPort, appConfig.DbConfig.Username, appConfig.DbConfig.DbName, appConfig.DbConfig.Password)

	client, err := ent.Open("postgres", ConnetionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
