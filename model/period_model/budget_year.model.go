package period_model

import "gorm.io/gorm"

type BudgetYear struct {
	gorm.Model

	Year     int8 `gorm:"unique;not null" json:"-"`
	IsActive bool `gorm:"default:false" json:"-"`
}
