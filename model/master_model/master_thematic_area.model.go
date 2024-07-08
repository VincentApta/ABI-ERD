package master_model

import "gorm.io/gorm"

type MasterThematicArea struct {
	gorm.Model

	Area string `gorm:"not null" json:"-"`
}
