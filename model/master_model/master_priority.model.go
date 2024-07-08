package master_model

import "gorm.io/gorm"

type MasterPriority struct {
	gorm.Model

	Priority string `gorm:"not null" json:"-"`
}
