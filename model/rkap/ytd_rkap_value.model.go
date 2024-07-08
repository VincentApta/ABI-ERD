package rkap_model

import "gorm.io/gorm"

type YtdRkapValue struct {
	gorm.Model

	BudgetItemDetailId uint    `gorm:"not null" json:"-"`
	AbiRkapCurrentYear float32 `json:"-"`
	YtdRkapJan         float32 `json:"-"`
	YtdRkapFeb         float32 `json:"-"`
	YtdRkapMar         float32 `json:"-"`
	YtdRkapApr         float32 `json:"-"`
	YtdRkapMay         float32 `json:"-"`
	YtdRkapJun         float32 `json:"-"`
	YtdRkapJul         float32 `json:"-"`
	YtdRkapAug         float32 `json:"-"`
	YtdRkapSep         float32 `json:"-"`
	YtdRkapOct         float32 `json:"-"`
	YtdRkapNov         float32 `json:"-"`
	YtdRkapDec         float32 `json:"-"`

	BudgetItemDetail *BudgetItemDetail `json:"-"`
}
