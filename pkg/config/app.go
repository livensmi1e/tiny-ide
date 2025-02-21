package config

type Config struct {
	Mode     string
	Addr     string
	Port     string
	Version  string
	Store    StoreConfig
	Queue    QueueConfig
	Executor Executor
}

// TODO: read from .env
func New() *Config {
	return &Config{
		Addr:    "localhost",
		Port:    "8000",
		Mode:    "development",
		Version: "v1",
		Store: StoreConfig{
			Driver: "postgres",
			DSN:    "postgres://postgres:password@localhost:5432/postgres?sslmode=disable",
		},
		Queue: QueueConfig{
			Addr:     "localhost:6379",
			Password: "",
		},
		Executor: Executor{
			Addr: "dns:localhost:50001",
		},
	}
}

func (cfg *Config) isDev() bool {
	return cfg.Mode == "development"
}
