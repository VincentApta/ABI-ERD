package master_model

import "gorm.io/gorm"

type ActualPhysicalProgress struct {
	gorm.Model

	Status string `gorm:"not null" json:"-"`
}
