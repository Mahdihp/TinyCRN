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

	cfg := config.Initialize()
	fmt.Println(cfg)
	InitDb(cfg)

}

func InitDb(cfg config.AppConfiguration) {

	ConnetionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DbConfig.Host, cfg.DbConfig.DbPort, cfg.DbConfig.Username, cfg.DbConfig.DbName, cfg.DbConfig.Password)

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
