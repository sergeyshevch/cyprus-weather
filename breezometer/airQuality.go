package breezometer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type AirQualityService struct {
	client *Client
}

type AirQualityIndex struct {
	DisplayName       string `json:"display_name"`
	AQI               int32  `json:"aqi"`
	AQIDisplay        string `json:"aqi_display"`
	Color             string `json:"color"`
	Category          string `json:"category"`
	DominantPollutant string `json:"dominant_pollutant"`
}

type currentConditionsResponse struct {
	Metadata *Metadata          `json:"metadata"`
	Data     *CurrentConditions `json:"data"`
}

// TODO: Incomplete, Documentation missing
type CurrentConditions struct {
	Datetime      string                     `json:"datetime"`
	DataAvailable bool                       `json:"data_available"`
	Indexes       map[string]AirQualityIndex `json:"indexes"`
}

func (s *AirQualityService) CurrentConditions(
	ctx context.Context,
	lat string,
	lon string,
) (*CurrentConditions, *http.Response, error) {

	apiEndpoint := fmt.Sprintf(
		"/air-quality/v2/current-conditions?lat=%s&lon=%s",
		url.QueryEscape(lat),
		url.QueryEscape(lon),
	)

	req, err := s.client.NewRequest(ctx, apiEndpoint)
	if err != nil {
		return nil, nil, err
	}

	currentConditions := new(currentConditionsResponse)
	resp, err := s.client.Do(req, currentConditions)
	if err != nil {
		return nil, resp, err
	}

	return currentConditions.Data, resp, nil
}
