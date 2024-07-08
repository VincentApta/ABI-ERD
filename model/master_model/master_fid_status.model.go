package master_model

import "gorm.io/gorm"

type MasterFidStatus struct {
	gorm.Model

	Status string `gorm:"not null" json:"-"`
}
