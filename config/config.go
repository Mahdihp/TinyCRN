package config

import (
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
	"time"
)

var k = koanf.New(".config")

type AppConfiguration struct {
	Server          Server
	DbConfig        DbConfig
	MigrationConfig Migration
	AuthConfig      AuthConfig
}

type Server struct {
	Host string
	Port string
}
type DbConfig struct {
	Host     string
	DbPort   int
	DbName   string
	Username string
	Password string
	SeedData bool
}
type Migration struct {
	MigrationDir   string
	MigrationTable string
}

type AuthConfig struct {
	SignKey                 string
	AccessExpirationTime    time.Duration
	RefreshExpirationTime   time.Duration
	AccessSubject           string
	RefreshSubject          string
	GracefulShutdownTimeout time.Duration
}

func Initialize() AppConfiguration {

	if err := k.Load(file.Provider("config/config.json"), json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	var out AppConfiguration

	// Quick unmarshal.
	k.Unmarshal("Server", &out.Server)
	k.Unmarshal("DbConfig", &out.DbConfig)
	k.Unmarshal("AuthConfig", &out.AuthConfig)

	return out
}
