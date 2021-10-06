package directusgosdk

type Directus struct {
	config *Config
	client *Client
}

func NewDirectus(c *Config) *Directus {
	return &Directus{
		config: c,
		client: NewClient(c),
	}
}

func (d *Directus) GetClient() *Client {
	return d.client
}
