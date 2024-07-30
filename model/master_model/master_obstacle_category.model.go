package master_model

import "gorm.io/gorm"

type MasterObstacleCategory struct {
	gorm.Model

	Category string `gorm:"not null" json:"-"`
}
