package web

// Defaults
var (
	defaultEnv     = "dev"
	defaultPort    = 8080
	defaultAPIPath = "/"
	defaultWebDir  = "./web"
)

// Config type
type Config struct {
	Env     string
	Port    int
	APIPath string
	WebDir  string
}

// Create config
func NewConfig(env string, port int, apiPath, webDir string) *Config {
	return &Config{
		Env:     env,
		Port:    port,
		APIPath: apiPath,
		WebDir:  webDir,
	}
}

// Create default
func DefaultConfig() *Config {
	return NewConfig(
		defaultEnv,
		defaultPort,
		defaultAPIPath,
		defaultWebDir)
}
