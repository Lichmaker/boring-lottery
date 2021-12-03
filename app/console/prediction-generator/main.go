package main

import (
	"github.com/lichmaker/boring-lottery/app/models/prediction"
	"github.com/lichmaker/boring-lottery/bootstrap"
	"github.com/lichmaker/boring-lottery/config"
	"github.com/lichmaker/boring-lottery/pkg/god"
)

func main() {
	config.Initialize()
	bootstrap.SetupDB()

	v := god.Get()

	var model prediction.Prediction
	model.Blue = uint64(v.Blue)
	model.Red1 = uint64(v.Red1)
	model.Red2 = uint64(v.Red2)
	model.Red3 = uint64(v.Red3)
	model.Red4 = uint64(v.Red4)
	model.Red5 = uint64(v.Red5)
	model.Red6 = uint64(v.Red6)
	model.Create()
}
