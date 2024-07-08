package master_model

import "gorm.io/gorm"

type MasterWk struct {
	gorm.Model

	Wk string `gorm:"not null" json:"-"`
}
