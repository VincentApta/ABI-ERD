package master_model

import "gorm.io/gorm"

type MasterRegion struct {
	gorm.Model

	Region string `gorm:"not null" json:"-"`
}
