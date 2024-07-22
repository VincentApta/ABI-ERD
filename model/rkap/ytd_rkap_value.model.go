package rkap_model

import "gorm.io/gorm"

type YtdWorksheetValue struct {
	gorm.Model

	BudgetItemDetailId      uint    `gorm:"not null" json:"-"`
	AbiWorksheetCurrentYear float32 `json:"-"`
	YtdWorksheetJan         float32 `json:"-"`
	YtdWorksheetFeb         float32 `json:"-"`
	YtdWorksheetMar         float32 `json:"-"`
	YtdWorksheetApr         float32 `json:"-"`
	YtdWorksheetMay         float32 `json:"-"`
	YtdWorksheetJun         float32 `json:"-"`
	YtdWorksheetJul         float32 `json:"-"`
	YtdWorksheetAug         float32 `json:"-"`
	YtdWorksheetSep         float32 `json:"-"`
	YtdWorksheetOct         float32 `json:"-"`
	YtdWorksheetNov         float32 `json:"-"`
	YtdWorksheetDec         float32 `json:"-"`
	Status                  string  `json:"-"`

	BudgetItemDetail *BudgetItemDetail `json:"-"`
}
