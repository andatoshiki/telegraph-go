package telegraph

// http://telegra.ph/api

// constants
const (
	APIBaseURL = "https://api.telegra.ph"
)

// Verbose flag for logging
var Verbose bool // default: false

// Client struct
type Client struct {
	AccessToken string
	Socks5Proxy string
}

// Create creates a new Telegraph client.
func Create(shortName, authorName, authorURL string, socks5Proxy string) (*Client, error) {
	client := Client{Socks5Proxy: socks5Proxy}

	created, err := client.CreateAccount(shortName, authorName, authorURL)

	if err == nil {
		return &Client{AccessToken: created.AccessToken, Socks5Proxy: socks5Proxy}, nil
	}

	return nil, err
}

// Load a Telegraph client with access token.
func Load(accessToken string, socks5Proxy string) (*Client, error) {
	client := Client{AccessToken: accessToken, Socks5Proxy: socks5Proxy}

	_, err := client.GetAccountInfo(nil)

	if err == nil {
		return &client, nil
	}

	return nil, err
}
