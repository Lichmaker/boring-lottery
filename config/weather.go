package config

import "github.com/lichmaker/boring-lottery/pkg/config"

func init() {
	config.Viper.Set("weather", map[string]interface{}{
		"free":  config.Get("WEATHER_FREE_KEY", ""),
		"business": config.Get("WEATHER_BUSINESS_KEY", ""),
		"location": config.Get("WEATHER_LOCATION", ""),
	})
}
