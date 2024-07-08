package master_model

import "gorm.io/gorm"

type MasterCategory struct {
	gorm.Model

	Category string `gorm:"not null" json:"-"`
}
