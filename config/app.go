package config

import "github.com/lichmaker/boring-lottery/pkg/config"

func init() {
	config.Viper.Set("app", map[string]interface{}{
		"port":  config.Get("PORT", "8002"),
		"debug": config.Get("DEBUG", "FALSE"),
	})
}
