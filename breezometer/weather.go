package breezometer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type WeatherService struct {
	client *Client
}

type DailyForecast struct {
	StartDate string `json:"start_date"`
	Sun       struct {
		SunriseTime string `json:"sunrise_time"`
		SunsetTime  string `json:"sunset_time"`
	} `json:"sun"`
	Moon struct {
		MoonRiseTime string  `json:"moonrise_time"`
		MoonSetTime  string  `json:"moonset_time"`
		MoonPhase    string  `json:"moon_phase"`
		MoonAge      float64 `json:"moon_age"`
	} `json:"moon"`
	MaxUVIndex int `json:"max_uv_index"`
}

type WeatherForecast struct {
	Datetime             string         `json:"datetime"`
	IsDayTime            bool           `json:"is_day_time"`
	IconCode             IconCode       `json:"icon_code"`
	WeatherText          string         `json:"weather_text"`
	Temperature          *Measure       `json:"temperature"`
	FeelsLikeTemperature *Measure       `json:"feels_like_temperature"`
	RelativeHumidity     int32          `json:"relative_humidity"`
	Precipitation        *Precipitation `json:"precipitation"`
	Wind                 *Wind          `json:"wind"`
	WindGust             *Measure       `json:"wind_gust"`
	Pressure             *Measure       `json:"pressure"`
	Visibility           *Measure       `json:"visibility"`
	DewPoint             *Measure       `json:"dew_point"`
	CloudCover           int            `json:"cloud_cover"`
}

type Precipitation struct {
	PrecipitationProbability int      `json:"precipitation_probability"`
	TotalPrecipitation       *Measure `json:"total_precipitation"`
}

type Wind struct {
	Speed     *Measure `json:"speed"`
	Direction int      `json:"direction"`
}

type currentWeatherConditionsResponse struct {
	Metadata *Metadata        `json:"metadata"`
	Data     *WeatherForecast `json:"data"`
}

type dailyWeatherForecastResponse struct {
	Metadata *Metadata        `json:"metadata"`
	Data     []*DailyForecast `json:"data"`
}

func (s *WeatherService) CurrentConditions(
	ctx context.Context,
	lat string,
	lon string,
	metadata *bool,
	units *string,
) (*WeatherForecast, *http.Response, error) {
	apiEndpoint := fmt.Sprintf(
		"/weather/v1/current-conditions?lat=%s&lon=%s",
		url.QueryEscape(lat),
		url.QueryEscape(lon),
	)

	addQueryMetadata(&apiEndpoint, metadata)
	addQueryUnits(&apiEndpoint, units)

	req, err := s.client.NewRequest(ctx, apiEndpoint)
	if err != nil {
		return nil, nil, err
	}

	currentConditions := new(currentWeatherConditionsResponse)
	resp, err := s.client.Do(req, currentConditions)
	if err != nil {
		return nil, resp, err
	}

	return currentConditions.Data, resp, nil
}

func (s *WeatherService) DailyForecast(
	ctx context.Context,
	lat string,
	lon string,
	days *int,
	metadata *bool,
) ([]*DailyForecast, *http.Response, error) {
	daysCount := 1
	if days != nil {
		daysCount = *days
	}

	apiEndpoint := fmt.Sprintf(
		"/weather/v1/forecast/daily?lat=%s&lon=%s&days=%s",
		url.QueryEscape(lat),
		url.QueryEscape(lon),
		url.QueryEscape(strconv.Itoa(daysCount)),
	)

	addQueryMetadata(&apiEndpoint, metadata)

	req, err := s.client.NewRequest(ctx, apiEndpoint)
	if err != nil {
		return nil, nil, err
	}

	currentConditions := new(dailyWeatherForecastResponse)
	resp, err := s.client.Do(req, currentConditions)
	if err != nil {
		return nil, resp, err
	}

	return currentConditions.Data, resp, nil
}
