package period_model

import (
	"time"

	"gorm.io/gorm"
)

type BudgetYear struct {
	gorm.Model

	Year            int8      `gorm:"unique;not null" json:"-"`
	StartAdjustment time.Time `json:"-"`
	EndAdjustment   time.Time `json:"-"`
}
