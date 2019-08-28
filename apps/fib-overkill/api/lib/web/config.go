package web

// Defaults
var (
	defaultPort        = 8080
	defaultAPIPath     = "/"
	defaultRedisURL    = "redis://localhost:6379"
	defaultPostgresURL = "postgresql://localhost:5432/test?sslmode=disable"
)

// Config type
type Config struct {
	Port        int
	APIPath     string
	RedisURL    string
	PostgresURL string
}

// Create config
func NewConfig(port int, apiPath, redisURL, postgresURL string) *Config {
	return &Config{
		Port:        port,
		APIPath:     apiPath,
		RedisURL:    redisURL,
		PostgresURL: postgresURL,
	}
}

// Create default
func DefaultConfig() *Config {
	return NewConfig(
		defaultPort,
		defaultAPIPath,
		defaultRedisURL,
		defaultPostgresURL)
}
