package config

type Config struct {
	Mode    string
	Addr    string
	Port    string
	Version string
	Store   StoreConfig
	Queue   QueueConfig
}

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
	}
}

func (cfg *Config) isDev() bool {
	return cfg.Mode == "development"
}
