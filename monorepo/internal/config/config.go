package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	PORT string
}

type RabbitMQ struct {
	ConnectionString string
	Name             string
}

type Storage struct {
	BaseDir string
}

type Config struct {
	Server  Server
	Queue   RabbitMQ
	Storage Storage
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load configuration file: %s\n", err)
	}

	return &Config{
		Server: Server{
			PORT: mustLoadEnv("PORT", ":3000"),
		},
		Queue: RabbitMQ{
			ConnectionString: mustLoadEnv("RABBIT_CONNECTION_STR", "amqp://guest:guest@localhost:5672"),
		},
		Storage: Storage{
			BaseDir: mustLoadEnv("STORAGE_DIR", "tmp"),
		},
	}
}

func mustLoadEnv(key string, fallback string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return fallback
	}
	return value
}
