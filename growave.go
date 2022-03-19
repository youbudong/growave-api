package growaveapi

import (
	"context"
	"net/url"

	"github.com/go-resty/resty/v2"
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

	baseURL *url.URL
	// Services used for communicating with the API
	User   UserService
	Reward RewardService
}

type Result struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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

	c := &Client{
		Client:  restyclient,
		baseURL: baseURL,
	}

	c.User = &UserServiceOp{client: c}
	c.Reward = &RewardServiceOp{client: c}

	return c
}
