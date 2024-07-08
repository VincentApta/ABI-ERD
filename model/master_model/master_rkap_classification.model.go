package master_model

import "gorm.io/gorm"

type MasterRkapClassification struct {
	gorm.Model

	Code        int8   `gorm:"not null" json:"-"`
	Description string `json:"description"`
}
