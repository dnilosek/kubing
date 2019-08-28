package web

// Defaults
var (
	defaultDBURL         = "redis://localhost:6379"
	defaultInputChannel  = "message"
	defaultOutputChannel = "values"
)

// Config type
type Config struct {
	DBURL         string
	InputChannel  string
	OutputChannel string
}

// Create config
func NewConfig(dbURL, inChan, outChan string) *Config {
	return &Config{
		DBURL:         dbURL,
		InputChannel:  inChan,
		OutputChannel: outChan,
	}
}

// Default config
func DefaultConfig() *Config {
	return NewConfig(defaultDBURL, defaultInputChannel, defaultOutputChannel)
}
