package rkap_model

import "gorm.io/gorm"

type DdRkapValue struct {
	gorm.Model

	BudgetItemDetailId uint    `gorm:"not null" json:"-"`
	DdRkapJan          float32 `json:"-"`
	DdRkapFeb          float32 `json:"-"`
	DdRkapMar          float32 `json:"-"`
	DdRkapApr          float32 `json:"-"`
	DdRkapMay          float32 `json:"-"`
	DdRkapJun          float32 `json:"-"`
	DdRkapJul          float32 `json:"-"`
	DdRkapAug          float32 `json:"-"`
	DdRkapSep          float32 `json:"-"`
	DdRkapOct          float32 `json:"-"`
	DdRkapNov          float32 `json:"-"`
	DdRkapDec          float32 `json:"-"`

	BudgetItemDetail *BudgetItemDetail `json:"-"`
}
