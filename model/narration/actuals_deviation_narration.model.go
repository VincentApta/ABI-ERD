package narration_model

import (
	"pertamina_abi/model/master_model"
	"pertamina_abi/model/period_model"
	rkap_model "pertamina_abi/model/rkap"

	"gorm.io/gorm"
)

type ActualsDeviationNarration struct {
	gorm.Model

	ActualsForecastPeriodId uint   `gorm:"not null" json:"-"`
	CategoryId              uint   `gorm:"not null" json:"-"`
	BudgetItemDetailId      uint   `gorm:"not null" json:"-"`
	Narration               string `json:"-"`

	ActualsForecastPeriod *period_model.ActualsForecastPeriod `json:"-"`
	Category              *master_model.MasterCategory        `json:"-"`
	BudgetItemDetail      *rkap_model.BudgetItemDetail        `json:"-"`
}
