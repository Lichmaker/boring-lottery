package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lichmaker/boring-lottery/app/models/prediction"
	"github.com/lichmaker/boring-lottery/app/models/results"
	"github.com/lichmaker/boring-lottery/bootstrap"
	"github.com/lichmaker/boring-lottery/config"
	"github.com/lichmaker/boring-lottery/pkg/god"
)

func main() {
	config.Initialize()
	bootstrap.SetupDB()

	// 周二/四/日开奖日的晚上 21:00 ~ 次日 01:30 之间， 不进行预测生成
	now := time.Now()
	location, _ := time.LoadLocation("Asia/Shanghai")
	nowTime := now.Unix()
	weekDay := now.Weekday().String()
	today := now.Format("2006-01-02")
	switch weekDay {
	case time.Monday.String(), time.Wednesday.String(), time.Friday.String():
		if f, _ := time.ParseInLocation("2006-01-02 15:04:05", today+" 01:30:00", location); f.Unix() >= nowTime {
			fmt.Println("暂停生成！", f.Format("2006-01-02 15:04:05"))
			return
		}
	case time.Tuesday.String(), time.Thursday.String(), time.Sunday.String():
		if f, _ := time.ParseInLocation("2006-01-02 15:04:05", today+" 21:00:00", location); f.Unix() <= nowTime {
			fmt.Println("暂停生成！", f.Format("2006-01-02 15:04:05"))
			return
		}
	}

	v := god.Get()

	// 读取最新一个期号，然后+1
	query, _ := results.GetLatest()
	nextPeriod, _ := strconv.ParseInt(query.Period, 10, 64)
	nextPeriod++

	var model prediction.Prediction
	model.Blue = uint64(v.Blue)
	model.Red1 = uint64(v.Red1)
	model.Red2 = uint64(v.Red2)
	model.Red3 = uint64(v.Red3)
	model.Red4 = uint64(v.Red4)
	model.Red5 = uint64(v.Red5)
	model.Red6 = uint64(v.Red6)
	model.Period = fmt.Sprintf("%d", nextPeriod)
	model.Create()
}
