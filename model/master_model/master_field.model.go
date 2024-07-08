package master_model

import "gorm.io/gorm"

type MasterField struct {
	gorm.Model

	WkId uint   `gorm:"not null" json:"-"`
	Field string `gorm:"not null" json:"-"`

	Wk *MasterWk `json:"-"`
}
