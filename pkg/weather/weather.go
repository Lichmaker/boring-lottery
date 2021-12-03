package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lichmaker/boring-lottery/pkg/config"
)

// https://dev.qweather.com/docs/api/air/air-now/
type airQueryRespond struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Now        struct {
		PubTime  string `json:"pubTime"`
		Aqi      string `json:"aqi"`
		Level    string `json:"level"`
		Category string `json:"category"`
		Primary  string `json:"primary"`
		Pm10     string `json:"pm10"`
		Pm2P5    string `json:"pm2p5"`
		No2      string `json:"no2"`
		So2      string `json:"so2"`
		Co       string `json:"co"`
		O3       string `json:"o3"`
	} `json:"now"`
	Station []struct {
		PubTime  string `json:"pubTime"`
		Name     string `json:"name"`
		ID       string `json:"id"`
		Aqi      string `json:"aqi"`
		Level    string `json:"level"`
		Category string `json:"category"`
		Primary  string `json:"primary"`
		Pm10     string `json:"pm10"`
		Pm2P5    string `json:"pm2p5"`
		No2      string `json:"no2"`
		So2      string `json:"so2"`
		Co       string `json:"co"`
		O3       string `json:"o3"`
	} `json:"station"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}

type weatherQueryRespond struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Now        struct {
		ObsTime   string `json:"obsTime"`
		Temp      string `json:"temp"`
		FeelsLike string `json:"feelsLike"`
		Icon      string `json:"icon"`
		Text      string `json:"text"`
		Wind360   string `json:"wind360"`
		WindDir   string `json:"windDir"`
		WindScale string `json:"windScale"`
		WindSpeed string `json:"windSpeed"`
		Humidity  string `json:"humidity"`
		Precip    string `json:"precip"`
		Pressure  string `json:"pressure"`
		Vis       string `json:"vis"`
		Cloud     string `json:"cloud"`
		Dew       string `json:"dew"`
	} `json:"now"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}

func QueryAir() airQueryRespond {
	var res airQueryRespond
	url := fmt.Sprintf("https://api.qweather.com/v7/air/now?location=%s&key=%s", config.Get("weather.location"), config.Get("weather.business"))
	httpClient := new(http.Client)
	airRequest, _ := http.NewRequest("GET", url, nil)
	respond, err := httpClient.Do(airRequest)
	if err != nil {
		fmt.Println("爬虫请求执行失败:", err)
		return res
	}
	defer respond.Body.Close()
	dataByte, _ := ioutil.ReadAll(respond.Body)
	json.Unmarshal(dataByte, &res)
	return res
}

func QueryWeather() weatherQueryRespond {
	var res weatherQueryRespond
	url := fmt.Sprintf("https://api.qweather.com/v7/weather/now?location=%s&key=%s", config.Get("weather.location"), config.Get("weather.business"))
	httpClient := new(http.Client)
	airRequest, _ := http.NewRequest("GET", url, nil)
	respond, err := httpClient.Do(airRequest)
	if err != nil {
		fmt.Println("爬虫请求执行失败:", err)
		return res
	}
	defer respond.Body.Close()
	dataByte, _ := ioutil.ReadAll(respond.Body)
	json.Unmarshal(dataByte, &res)
	return res
}
