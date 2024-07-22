package rkap_model

import "gorm.io/gorm"

type DdWorksheetValue struct {
	gorm.Model

	BudgetItemDetailId uint    `gorm:"not null" json:"-"`
	DdWorksheetJan     float32 `json:"-"`
	DdWorksheetFeb     float32 `json:"-"`
	DdWorksheetMar     float32 `json:"-"`
	DdWorksheetApr     float32 `json:"-"`
	DdWorksheetMay     float32 `json:"-"`
	DdWorksheetJun     float32 `json:"-"`
	DdWorksheetJul     float32 `json:"-"`
	DdWorksheetAug     float32 `json:"-"`
	DdWorksheetSep     float32 `json:"-"`
	DdWorksheetOct     float32 `json:"-"`
	DdWorksheetNov     float32 `json:"-"`
	DdWorksheetDec     float32 `json:"-"`
	Status             string  `json:"-"`

	BudgetItemDetail *BudgetItemDetail `json:"-"`
}
