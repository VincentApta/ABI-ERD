package master_model

import "gorm.io/gorm"

type MasterAp struct {
	gorm.Model

	Name string `json:"-"`
	Code string `gorm:"not null" json:"-"`
}
