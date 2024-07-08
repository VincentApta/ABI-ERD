package master_model

import "gorm.io/gorm"

type MasterBusinessStream struct {
	gorm.Model

	Stream string `gorm:"not null" json:"-"`
}
