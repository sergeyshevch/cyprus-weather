package main

import (
  "context"
  "fmt"
  "net/http/httputil"

  "github.com/samber/lo"
  "go.uber.org/zap"

  "github.com/sergeyshevch/cyprus-weather/breezometer"
  "github.com/sergeyshevch/cyprus-weather/config"
  "github.com/sergeyshevch/cyprus-weather/twitter"
  "github.com/sergeyshevch/cyprus-weather/utils"
)

const limassolLat, limassolLon = "34.67889288021057", "33.04125490452733"

func main() {
  ctx := context.Background()

  log := utils.InitLogger()
  defer func(log *zap.Logger) {
    _ = log.Sync()
  }(log)

  err := config.InitConfig(log)
  if err != nil {
    log.Panic("Config initialization failed", zap.Error(err))
  }
  // Your code here...

  var accessKey, accessSecret string
  if config.AccessKey() == "" || config.AccessSecret() == "" {
    if config.InteractiveLogin() {
      accessKey, accessSecret, err = utils.TwitterLogin(log)
    } else {
      log.Panic("Access key and secret are not set and InteractiveLogin is not enabled")
    }
  } else {
    accessKey = config.AccessKey()
    accessSecret = config.AccessSecret()
  }

  client := twitter.InitClient(log, accessKey, accessSecret)

  if err := client.VerifyParams(); err != nil {
    log.Panic("Verification failed", zap.Error(err))
  }

  breezometerClient, err := breezometer.NewClient(nil, "293cca03eafb40f1b461c339fa24bec8")
  if err != nil {
    log.Panic("Breezometer initialization failed", zap.Error(err))
  }

  currentConditions, resp, err := breezometerClient.WeatherService.CurrentConditions(ctx, limassolLat, limassolLon, nil, nil)
  if err != nil {
    dr, _ := httputil.DumpResponse(resp, true)

    log.Panic("Breezometer request failed", zap.Error(err), zap.String("response", string(dr)))
  }

  dailyForecast, resp, err := breezometerClient.WeatherService.DailyForecast(ctx, limassolLat, limassolLon, lo.ToPtr(1), nil)
  if err != nil {
    dr, _ := httputil.DumpResponse(resp, true)

    log.Panic("Breezometer request failed", zap.Error(err), zap.String("response", string(dr)))
  }

  if len(dailyForecast) < 1 {
    log.Panic("Today's daily forecast is not available")
  }
  todayForecast := dailyForecast[0]

  tweet := fmt.Sprintf(
    // "Good morning! Today will be %s (%s). \nTemperature: %.1f (Feels like: %.1f)\nHuminidity: %d%%\nWind Speed: %.1f %s\nMax UV: %d\n Info by @breezometer",
    "Good morning! Today will be %s (%s). \nTemperature: %.1f (Feels like: %.1f)\nHuminidity: %d%%\nWind Speed: %.1f %s\nMax UV: %d",
    breezometer.IconCodeToEmoji(currentConditions.IconCode),
    currentConditions.WeatherText,
    currentConditions.Temperature.Value,
    currentConditions.FeelsLikeTemperature.Value,
    currentConditions.RelativeHumidity,
    currentConditions.Wind.Speed.Value,
    currentConditions.Wind.Speed.Units,
    todayForecast.MaxUVIndex,
  )

  if err := client.CreateTweet(tweet); err != nil {
    log.Panic("Cannot create tweet", zap.Error(err))
  }

  log.Info("Tweet created")
}
