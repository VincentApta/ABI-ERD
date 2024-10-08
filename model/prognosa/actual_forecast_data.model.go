package prognosa_model

import (
	"pertamina_abi/model/period_model"
	rkap_model "pertamina_abi/model/rkap"

	"gorm.io/gorm"
)

type ActualsForecastData struct {
	gorm.Model

	ActualsForecastPeriodId uint    `json:"-"`
	BudgetItemDetailId      uint    `json:"-"`
	EeYtd                   float32 `json:"-"`
	DdRealization           float32 `json:"-"`
	ForecastStatus           float32 `json:"-"`

	ActualsForecastPeriod *period_model.ActualsForecastPeriod `json:"-"`
	BudgetItemDetail      *rkap_model.BudgetItemDetail        `json:"-"`
}
