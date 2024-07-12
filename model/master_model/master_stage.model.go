package master_model

import "gorm.io/gorm"

type MasterStage struct {
	gorm.Model

	Stage string `gorm:"not null" json:"-"`
}
