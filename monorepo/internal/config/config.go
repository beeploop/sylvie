package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	JSON_DB_PATH string
}

type FFMPEG struct {
	FfmpegPath  string
	FfprobePath string
}

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
	DB      DB
	FFMPEG  FFMPEG
	Server  Server
	Queue   RabbitMQ
	Storage Storage
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load configuration file: %s\n", err)
	}

	return &Config{
		DB: DB{
			JSON_DB_PATH: mustLoadEnv("JSON_DB_FILE_PATH", "tmp/db.json"),
		},
		FFMPEG: FFMPEG{
			FfmpegPath:  mustLoadEnv("FFMPEG_PATH", "/usr/bin/ffmpeg"),
			FfprobePath: mustLoadEnv("FFPROBE_PATH", "/usr/bin/ffprobe"),
		},
		Server: Server{
			PORT: mustLoadEnv("PORT", ":3000"),
		},
		Queue: RabbitMQ{
			ConnectionString: mustLoadEnv("RABBIT_CONNECTION_STR", "amqp://guest:guest@localhost:5672"),
			Name:             mustLoadEnv("RABBIT_QUEUE_NAME", "sylvie:transcode"),
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
