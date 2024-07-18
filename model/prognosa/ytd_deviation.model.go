package prognosa_model

import (
	rkap_model "pertamina_abi/model/rkap"

	"gorm.io/gorm"
)

type YtdDeviation struct {
	gorm.Model

	BudgetItemDetailId   uint    `gorm:"not null" json:"-"`
	RealizationWorksheet float32 `json:"-"`
	EeRealization        float32 `json:"-"`
	WorksheetForecast    float32 `json:"-"`
	ForecastRealization  float32 `json:"-"`

	BudgetItemDetail *rkap_model.BudgetItemDetail `json:"-"`
}
