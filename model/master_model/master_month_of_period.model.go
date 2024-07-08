package master_model

import "gorm.io/gorm"

type MasterMonthOfPeriod struct {
	gorm.Model

	Name string `gorm:"not null" json:"-"`
	Code string `gorm:"not null" json:"-"`
}
