package prognosa_model

import (
	rkap_model "pertamina_abi/model/rkap"

	"gorm.io/gorm"
)

type YtdRealizationValue struct {
	gorm.Model

	BudgetItemDetailId uint    `gorm:"not null" json:"-"`
	YtdRealizationJan  float32 `json:"-"`
	YtdRealizationFeb  float32 `json:"-"`
	YtdRealizationMar  float32 `json:"-"`
	YtdRealizationApr  float32 `json:"-"`
	YtdRealizationMay  float32 `json:"-"`
	YtdRealizationJun  float32 `json:"-"`
	YtdRealizationJul  float32 `json:"-"`
	YtdRealizationAug  float32 `json:"-"`
	YtdRealizationSep  float32 `json:"-"`
	YtdRealizationOct  float32 `json:"-"`
	YtdRealizationNov  float32 `json:"-"`
	YtdRealizationDec  float32 `json:"-"`

	BudgetItemDetail *rkap_model.BudgetItemDetail `json:"-"`
}
