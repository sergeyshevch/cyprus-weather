package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"go.uber.org/zap"

	"github.com/sergeyshevch/cyprus-weather/config"
)

type Client struct {
	client *twitter.Client
	log    *zap.Logger
}

func InitClient(log *zap.Logger, accessToken, accessSecret string) *Client {
	oauthConfig := oauth1.NewConfig(config.ConsumerKey(), config.ConsumerSecret())
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	return &Client{
		client: client,
		log:    log,
	}
}

func (c *Client) VerifyParams() error {
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := c.client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return err
	}

	c.log.Info("User verified", zap.String("name", user.Name), zap.String("screen_name", user.ScreenName))

	return nil
}

func (c *Client) CreateTweet(content string) error {
	_, _, err := c.client.Statuses.Update(content, nil)
	if err != nil {
		return err
	}

	return nil
}
