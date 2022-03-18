package growaveapi

import (
	"context"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	baseURL            = "https://growave.io/api"
	defaultHttpTimeout = 10
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
		TokenURL:     baseURL + "/access_token",
		Scopes:       app.Scopes,
	}

	client := cfg.Client(context.Background())
	restyclient := resty.NewWithClient(client)
	restyclient.SetBaseURL(baseURL)

	baseURL, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	logrus.Info(baseURL.Path)

	c := &Client{
		Client:  restyclient,
		baseURL: baseURL,
	}

	c.User = &UserServiceOp{client: c}
	return c
}
