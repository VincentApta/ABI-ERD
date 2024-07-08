package period_model

import (
	"pertamina_abi/model/master_model"
	"time"

	"gorm.io/gorm"
)

type ActualsForecastPeriod struct {
	gorm.Model

	BudgetYearId    uint      `gorm:"not null" json:"-"`
	MonthOfPeriodId uint      `gorm:"not null" json:"-"`
	StartDate       time.Time `gorm:"type:date; not null" json:"-"`
	EndDate         time.Time `gorm:"type:date; not null" json:"-"`
	IsActive        bool      `gorm:"default:true" json:"-"`

	BudgetYear    *BudgetYear                       `json:"-"`
	MonthOfPeriod *master_model.MasterMonthOfPeriod `json:"-"`
}
