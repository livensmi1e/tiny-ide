package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Mode     string
	Addr     string
	Port     string
	Version  string
	Store    StoreConfig
	Queue    QueueConfig
	Executor Executor
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	return &Config{
		Addr:    getEnv("SERVER_HOST", "localhost"),
		Port:    getEnv("SERVER_PORT", "8000"),
		Mode:    getEnv("MODE", "development"),
		Version: getEnv("VERSION", "v1"),
		Store: StoreConfig{
			Driver: "postgres",
			DSN:    buildDSN(),
		},
		Queue: QueueConfig{
			Addr:     buildQueueAddr(),
			Password: "",
		},
		Executor: Executor{
			Addr: buildRPCAddr(),
		},
	}
}

func (cfg *Config) isDev() bool {
	return cfg.Mode == "development"
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func buildDSN() string {
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	dbname := getEnv("DB_NAME", "postgres")

	return "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
}

func buildQueueAddr() string {
	host := getEnv("QUEUE_HOST", "localhost")
	port := getEnv("QUEUE_PORT", "6379")
	return host + ":" + port
}

func buildRPCAddr() string {
	host := getEnv("RPC_SERVER_HOST", "localhost")
	port := getEnv("RPC_SERVER_PORT", "50001")
	return "dsn:" + host + ":" + port
}
