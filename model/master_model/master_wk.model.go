package master_model

import "gorm.io/gorm"

type MasterWk struct {
	gorm.Model

	ZonaId uint   `gorm:"not null" json:"-"`
	Wk     string `gorm:"not null" json:"-"`

	Zona *MasterZona `json:"-"`
}
