package web

// Defaults
var (
	defaultDBURL = "redis://localhost:6379"
)

// Config type
type Config struct {
	DBURL string
}

// Create config
func NewConfig(dbURL string) *Config {
	return &Config{
		DBURL: dbURL,
	}
}

// Default config
func DefaultConfig() *Config {
	return NewConfig(defaultDBURL)
}
