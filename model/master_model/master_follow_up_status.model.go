package master_model

import "gorm.io/gorm"

type MasterFollowUpStatus struct {
	gorm.Model

	Status string `gorm:"not null" json:"-"`
}
