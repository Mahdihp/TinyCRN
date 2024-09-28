package postgres

import (
	"TinyCRM/config"
	"TinyCRM/ent"
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"log"
)

type PostgresDB struct {
	cfg    config.DbConfig
	client *ent.Client
}

func (m *PostgresDB) Connection() *ent.Client {
	return m.client
}

func New(cfg config.DbConfig) *PostgresDB {
	//ConnetionString := "postgres://postgres:postgres@localhost:5432/postgres"
	ConnetionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.DbPort, cfg.Username, cfg.DbName, cfg.Password)

	client, err := ent.Open(dialect.Postgres, ConnetionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if cfg.SeedData {
		DataSeed(client)
	}

	return &PostgresDB{cfg: cfg, client: client}
}

func DataSeed(client *ent.Client) {
	ctx := context.Background()

	count, err := client.Expert.Query().Count(ctx)
	if err != nil {
		log.Fatalf("failed count query: %v", err)
	}

	if count == 0 {
		client.Expert.CreateBulk(
			client.Expert.Create().SetUsername("mahdihp"),
		).Save(ctx)
	}
}
