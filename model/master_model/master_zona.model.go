package master_model

import "gorm.io/gorm"

type MasterZona struct {
	gorm.Model

	ApId uint   `gorm:"not null" json:"-"`
	Zone string `gorm:"not null" json:"-"`

	Ap *MasterAp `json:"-"`
}
