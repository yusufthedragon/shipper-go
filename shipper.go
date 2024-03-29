package shipper

// Configuration is struct for API config.
type Configuration struct {
	APIKey  string
	BaseURL string
}

// Conf is instance of Configuration struct.
var Conf = Configuration{
	BaseURL: "https://api.sandbox.shipper.id/public/v1",
}

// SetAPIKey sets the API Key for API Call.
func (c *Configuration) SetAPIKey(token string) *Configuration {
	c.APIKey = token

	return c
}

// SetProductionMode sets the API base URL to production.
func (c *Configuration) SetProductionMode(productionMode bool) {
	if productionMode {
		c.BaseURL = "https://api.shipper.id/prod/public/v1"
	}
}
