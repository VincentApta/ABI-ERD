package master_model

import "gorm.io/gorm"

type MasterActivityType struct {
	gorm.Model

	Activity string `gorm:"not null" json:"-"`
}
