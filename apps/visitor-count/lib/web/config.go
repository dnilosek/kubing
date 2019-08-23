package web

// Defaults
var (
	defaultEnv     = "dev"
	defaultPort    = 8080
	defaultAPIPath = "/"
	defaultWebDir  = "./web"
	defaultDBURL   = "redis://localhost:6379"
)

// Config type
type Config struct {
	Env     string
	Port    int
	APIPath string
	WebDir  string
	DBURL   string
}

// Create config
func NewConfig(env string, port int, apiPath, webDir, dbURL string) *Config {
	return &Config{
		Env:     env,
		Port:    port,
		APIPath: apiPath,
		WebDir:  webDir,
		DBURL:   dbURL,
	}
}

// Create default
func DefaultConfig() *Config {
	return NewConfig(
		defaultEnv,
		defaultPort,
		defaultAPIPath,
		defaultWebDir,
		defaultDBURL)
}
