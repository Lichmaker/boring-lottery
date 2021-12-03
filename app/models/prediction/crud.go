package prediction

import (
	"github.com/lichmaker/boring-lottery/pkg/model"
)

func (results *Prediction) Create() error {
	var err error
	r := model.DB.Create(&results)
	if err = r.Error; err != nil {
		return err
	} else {
		return nil
	}
}

func Get(count int) ([]Prediction, error) {
	var data []Prediction
	queryResult := model.DB.Order("id DESC").Limit(count).Find(&data)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	} else {
		return data, nil
	}
}
