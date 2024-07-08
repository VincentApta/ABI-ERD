package master_model

import "gorm.io/gorm"

type MasterSbtConnection struct {
	gorm.Model

	Connection string `gorm:"not null" json:"-"`
}
