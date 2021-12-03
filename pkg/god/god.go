package god

import (
	"reflect"
	"strconv"

	"github.com/lichmaker/boring-lottery/pkg/weather"
)

type godsPrediction struct {
	Blue int
	Red1 int
	Red2 int
	Red3 int
	Red4 int
	Red5 int
	Red6 int
}

func Get() *godsPrediction {
	var res godsPrediction

	// 从各个空气监测站读取空气信息
	airQuery := weather.QueryAir()
	// fmt.Println(airQuery)
	for _, v := range airQuery.Station {
		red1Uint, _ := strconv.ParseUint(v.Pm2P5, 10, 64)
		res.Red1 = res.Red1 + int(red1Uint)

		red2Uint, _ := strconv.ParseUint(v.Pm10, 10, 64)
		res.Red2 = res.Red2 + int(red2Uint)

		no2Uint, _ := strconv.ParseUint(v.No2, 10, 64)
		so2Uint, _ := strconv.ParseUint(v.So2, 10, 64)
		res.Red3 = res.Red3 + int(no2Uint) + int(so2Uint)

		o3Uint, _ := strconv.ParseUint(v.O3, 10, 64)
		res.Red4 = res.Red4 + int(o3Uint)
	}

	// 实况天气温度
	weatherQuery := weather.QueryWeather()
	tempUint, _ := strconv.ParseUint(weatherQuery.Now.Temp, 10, 64)
	feelsLikeUint, _ := strconv.ParseUint(weatherQuery.Now.FeelsLike, 10, 64)
	res.Blue = int(tempUint) + int(feelsLikeUint)

	// 风向
	wind360Uint, _ := strconv.ParseUint(weatherQuery.Now.Wind360, 10, 64)
	res.Red5 = int(wind360Uint)

	// 风力和风速
	scalUint, _ := strconv.ParseUint(weatherQuery.Now.WindScale, 10, 64)
	speedUint, _ := strconv.ParseUint(weatherQuery.Now.WindSpeed, 10, 64)
	res.Red6 = int(scalUint) * int(speedUint)

	// 取模，去重，返回
	return dataMod(&res)
}

func dataMod(data *godsPrediction) *godsPrediction {
	existsRedNumber := make(map[int64]interface{})

	reflectKey := reflect.TypeOf(data).Elem()
	reflectValue := reflect.ValueOf(data).Elem()
	for i := 0; i < reflectValue.NumField(); i++ {
		valueInt := reflectValue.Field(i).Int()
		keyString := reflectKey.Field(i).Name

		// 取模
		if keyString == "Blue" {
			valueInt = valueInt % 16
		} else {
			valueInt = valueInt % 33
		}

		if valueInt == 0 {
			if keyString == "Blue" {
				valueInt = 16
			} else {
				valueInt = 33
			}
		}

		// 去重
		if keyString != "Blue" {
			if _, ok := existsRedNumber[valueInt]; ok {
				valueInt++
				if valueInt > 33 {
					valueInt = 1
				}
			} else {
				existsRedNumber[valueInt] = 1
			}
		}
		reflectValue.Field(i).SetInt(valueInt)
	}

	return data
}
