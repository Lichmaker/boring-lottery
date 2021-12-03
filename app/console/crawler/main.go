package main

import (
	"fmt"

	"github.com/lichmaker/boring-lottery/app/models/results"
	"github.com/lichmaker/boring-lottery/bootstrap"
	"github.com/lichmaker/boring-lottery/config"
	"github.com/lichmaker/boring-lottery/pkg/crawler"
)

func main() {
	config.Initialize()
	bootstrap.SetupDB()

	data, err := crawler.Get()
	if err != nil {
		fmt.Println("失败！ ", err)
		return
	}

	model := results.Results{
		Period: data.Period,
		Blue:   data.Blue,
		Red1:   data.Red1,
		Red2:   data.Red2,
		Red3:   data.Red3,
		Red4:   data.Red4,
		Red5:   data.Red5,
		Red6:   data.Red6,
	}
	model.Create()
}
