package prognosa_model

import (
	rkap_model "pertamina_abi/model/rkap"

	"gorm.io/gorm"
)

type ActualForecastNotes struct {
	gorm.Model

	BudgetItemDetailId       uint   `gorm:"not null" json:"-"`
	StatusOfForecast         string `json:"-"`
	StatusNotes              string `json:"-"`
	InvestmentActivityUpdate string `json:"-"`
	RealizationGapNotes      string `json:"-"`
	ActionPlan               string `json:"-"`

	BudgetItemDetail *rkap_model.BudgetItemDetail
}
