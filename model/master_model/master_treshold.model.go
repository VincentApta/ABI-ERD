package master_model

import "gorm.io/gorm"

type MasterTreshold struct {
	gorm.Model

	Treshold string `gorm:"not null" json:"-"`
}
