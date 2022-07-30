package breezometer

import "fmt"

func addQueryMetadata(url *string, metadata *bool) {
	if metadata != nil && *metadata {
		*url = fmt.Sprintf("%s&metadata=%t", *url, *metadata)
	}
}

func addQueryUnits(url *string, units *string) {
	if units != nil {
		*url = fmt.Sprintf("%s&units=%s", *url, *units)
	}
}

// IconCodeToEmoji converts icon code to emoji
// Icon Code	Weather Text
// 1	Clear, cloudless sky
// 2	Clear, few cirrus
// 3	Clear with cirrus
// 4	Clear with few low clouds
// 5	Clear with few low clouds and few cirrus
// 6	Clear with few low clouds and cirrus
// 7	Partly cloudy
// 8	Partly cloudy and few cirrus
// 9	Partly cloudy and cirrus
// 10	Mixed with some thunderstorm clouds possible
// 11	Mixed with few cirrus with some thunderstorm clouds possible
// 12	Mixed with cirrus and some thunderstorm clouds possible
// 13	Clear but hazy
// 14	Clear but hazy with few cirrus
// 15	Clear but hazy with cirrus
// 16	Fog/low stratus clouds
// 17	Fog/low stratus clouds with few cirrus
// 18	Fog/low stratus clouds with cirrus
// 19	Mostly cloudy
// 20	Mostly cloudy and few cirrus
// 21	Mostly cloudy and cirrus
// 22	Overcast
// 23	Overcast with rain
// 24	Overcast with snow
// 25	Overcast with heavy rain
// 26	Overcast with heavy snow
// 27	Rain, thunderstorms likely
// 28	Light rain, thunderstorms likely
// 29	Storm with heavy snow
// 30	Heavy rain, thunderstorms likely
// 31	Mixed with showers
// 32	Mixed with snow showers
// 33	Overcast with light rain
// 34	Overcast with light snow
// 35	Overcast with mixture of snow and rain
//
// Source: https://docs.breezometer.com/api-documentation/weather-api/v1/#icon-codes-and-weather-texts
func IconCodeToEmoji(iconCode IconCode) string {
	switch int(iconCode) {
	case 1, 2, 3:
		return "☀️"
	case 4, 5, 6:
		return "🌤️"
	case 7, 8, 9:
		return "⛅"
	case 10, 11, 12:
		return "🌩️"
	case 13, 14, 15:
		return "🌁"
	case 16, 17, 18:
		return "🌫️"
	case 19, 20, 21, 22:
		return "☁️"
	case 23, 25, 31, 33, 35:
		return "🌧"
	case 24, 26, 32, 34:
		return "🌨"
	case 27, 28, 29, 30:
		return "⛈️"
	default:
		return "☀️"
	}
}
