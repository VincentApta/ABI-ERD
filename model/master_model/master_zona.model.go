package master_model

import "gorm.io/gorm"

type MasterZone struct {
	gorm.Model

	RegionId uint   `gorm:"not null" json:"-"`
	Zone     string `gorm:"not null" json:"-"`

	Region *MasterRegion `json:"-"`
}
