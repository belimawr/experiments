package config

// Config struct that holds all configuration parameters
type Config struct {
	Port int `env:"PROT" envDefault:"3000"`
}
