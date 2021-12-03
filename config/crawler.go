package config

import "github.com/lichmaker/boring-lottery/pkg/config"

func init() {
	config.Viper.Set("crawler", map[string]interface{}{
		"ua":          config.Get("CRAWLER_UA", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36"),
		"cookiename":  config.Get("CRAWLER_COOKIE_NAME", ""),
		"cookievalue": config.Get("CRAWLER_COOKIE_VALUE", ""),
	})
}
