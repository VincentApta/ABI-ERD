package master_model

import "gorm.io/gorm"

type EoyPhysicalProjection struct {
	gorm.Model

	Projection string `gorm:"not null" json:"-"`
}
