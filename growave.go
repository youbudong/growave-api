package growaveapi

import (
	"context"
	"net/url"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	baseURL              = "https://growave.io"
	defaultHttpTimeout   = 10
	defaultApiPathPrefix = "/api"
)

type App struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
}

type Client struct {
	Client *resty.Client

	baseURL    *url.URL
	pathPrefix string

	// Services used for communicating with the API
	User UserService
}

func New(app App) *Client {

	cfg := clientcredentials.Config{
		ClientID:     app.ClientID,
		ClientSecret: app.ClientSecret,
		TokenURL:     baseURL + defaultApiPathPrefix + "/access_token",
		Scopes:       app.Scopes,
	}

	client := cfg.Client(context.Background())
	restyclient := resty.NewWithClient(client)

	baseURL, _ := url.Parse(baseURL)

	c := &Client{
		Client:     restyclient,
		baseURL:    baseURL,
		pathPrefix: defaultApiPathPrefix,
	}
	return c
}
