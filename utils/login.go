package utils

import (
	"fmt"

	"github.com/dghubble/oauth1"
	twauth "github.com/dghubble/oauth1/twitter"
	"go.uber.org/zap"
)

const outOfBand = "oob"

func TwitterLogin(log *zap.Logger) (string, string, error) {
	config := oauth1.Config{
		ConsumerKey:    "2EIYYJmAnR1IuNlC5vzZI0OUg",
		ConsumerSecret: "5SaxOZ8fXNsef5FclWRiIJmthISnoX4l2twXuLuTzEwDujDTKI",
		CallbackURL:    outOfBand,
		Endpoint:       twauth.AuthorizeEndpoint,
	}

	requestToken, err := login(log, config)
	if err != nil {
		log.Error("Request Token Phase", zap.Error(err))

		return "", "", err
	}

	accessToken, err := receivePIN(log, config, requestToken)
	if err != nil {
		log.Error("Access Token Phase ", zap.Error(err))

		return "", "", err
	}

	log.Info("Login successful")
	log.Info("Access Token: ", zap.String("token", accessToken.Token))
	log.Info("Access Secret: ", zap.String("secret", accessToken.TokenSecret))

	return accessToken.Token, accessToken.TokenSecret, nil
}

func login(log *zap.Logger, config oauth1.Config) (requestToken string, err error) {
	requestToken, _, err = config.RequestToken()
	if err != nil {
		return "", err
	}

	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", err
	}

	log.Info("Please visit this URL to authorize the app: ", zap.String("url", authorizationURL.String()))

	return requestToken, err
}

func receivePIN(log *zap.Logger, config oauth1.Config, requestToken string) (*oauth1.Token, error) {
	log.Info("Please enter the PIN code: ")
	var verifier string
	_, err := fmt.Scanf("%s", &verifier)
	if err != nil {
		return nil, err
	}

	// Twitter ignores the oauth_signature on the access token request. The user
	// to which the request (temporary) token corresponds is already known on the
	// server. The request for a request token earlier was validated signed by
	// the consumer. Consumer applications can avoid keeping request token state
	// between authorization granting and callback handling.
	accessToken, accessSecret, err := config.AccessToken(requestToken, "secret does not matter", verifier)
	if err != nil {
		return nil, err
	}
	return oauth1.NewToken(accessToken, accessSecret), err
}
