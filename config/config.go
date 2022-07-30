package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	errMissingConsumerKey    = errors.New("consumer key variable is missing")
	errMissingConsumerSecret = errors.New("consumer secret variable is missing")
)

func InitConfig(log *zap.Logger) error {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".cyprus-weather")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file", zap.String("file", viper.ConfigFileUsed()))
	}

	if err := checkInitialConfiguration(); err != nil {
		return err
	}

	return nil
}

func ConsumerKey() string {
	return viper.GetString("consumer_key")
}

func ConsumerSecret() string {
	return viper.GetString("consumer_secret")
}

func AccessKey() string {
	return viper.GetString("access_key")
}

func AccessSecret() string {
	return viper.GetString("access_secret")
}

func BreezometerToken() string {
	return viper.GetString("breezometer_token")
}

func InteractiveLogin() bool {
	return viper.GetBool("interactive_login")
}

func checkInitialConfiguration() error {
	if ConsumerKey() == "" {
		return errMissingConsumerKey
	}

	if ConsumerSecret() == "" {
		return errMissingConsumerSecret
	}

	return nil
}
