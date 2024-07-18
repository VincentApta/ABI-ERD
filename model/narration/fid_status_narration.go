package narration_model

import (
	"pertamina_abi/model/period_model"

	"gorm.io/gorm"
)

type FidStatusNarration struct {
	gorm.Model

	ActualForecastPeriodId uint   `gorm:"not null" json:"-"`
	Narration              string `gorm:"not null" json:"-"`
	
	ActualForecastPeriod *period_model.ActualsForecastPeriod `json:"-"`
}
