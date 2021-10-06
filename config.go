package directusgosdk

type Config struct {
	endpoint string
}

func NewConfigWithEndpoint(endpoint string) *Config {
	return &Config{
		endpoint,
	}
}

func (c *Config) GetEndpoint() string {
	return c.endpoint
}
