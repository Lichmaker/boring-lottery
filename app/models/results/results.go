package results

import "time"

type Results struct {
	Period    string    `gorm:"column:period;type:varchar(10);primaryKey"`
	Blue      uint64    `gorm:"column:blue;type:uint;size:2;"`
	Red1      uint64    `gorm:"column:red_1;type:uint;size:2;"`
	Red2      uint64    `gorm:"column:red_2;type:uint;size:2;"`
	Red3      uint64    `gorm:"column:red_3;type:uint;size:2;"`
	Red4      uint64    `gorm:"column:red_4;type:uint;size:2;"`
	Red5      uint64    `gorm:"column:red_5;type:uint;size:2;"`
	Red6      uint64    `gorm:"column:red_6;type:uint;size:2;"`
	CreatedAt time.Time `gorm:"column:created_at;"`
}

type Tabler interface {
	TableName() string
}

func (Results) TableName() string {
	return "results"
}
