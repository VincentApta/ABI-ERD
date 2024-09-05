package narration_model

import (
	"pertamina_abi/model/period_model"
	rkap_model "pertamina_abi/model/rkap"

	"gorm.io/gorm"
)

type ForecastNarration struct {
	gorm.Model

	ActualsForecastPeriodId uint   `gorm:"not null" json:"-"`
	Category                string `gorm:"not null" json:"-"`
	BudgetItemDetailId      uint   `gorm:"not null" json:"-"`
	Narration               string `json:"-"`
	RoleGroup               string `gorm:"not null"`

	ActualsForecastPeriod *period_model.ActualsForecastPeriod `json:"-"`
	BudgetItemDetail      *rkap_model.BudgetItemDetail        `json:"-"`
}
