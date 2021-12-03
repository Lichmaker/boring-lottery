package results

import (
	"github.com/lichmaker/boring-lottery/pkg/model"
)

func (results *Results) Create() error {
	var err error
	r := model.DB.Create(&results)
	if err = r.Error; err != nil {
		return err
	} else {
		return nil
	}
}

func GetLast10Days() ([]Results, error) {
	var data []Results
	// now := time.Now()
	// daysAgo := now.AddDate(0,0,-10)
	// dateString := daysAgo.Format("2006-01-02 15:04:05")
	queryResult := model.DB.Order("created_at DESC").Limit(10).Find(&data)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	} else {
		return data, nil
	}
}

func GetLatest() (*Results, error) {
	var data Results
	// now := time.Now()
	// daysAgo := now.AddDate(0,0,-10)
	// dateString := daysAgo.Format("2006-01-02 15:04:05")
	queryResult := model.DB.Order("created_at DESC").Limit(10).Find(&data)
	// fmt.Printf("第一个%p\n", &data)
	if queryResult.Error != nil {
		return &data, queryResult.Error
	} else {
		return &data, nil
	}
}
