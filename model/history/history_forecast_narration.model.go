package history

import "gorm.io/gorm"

type HistoryForecastNarration struct {
	gorm.Model
	WorkPlan      string  `gorm:"not null"`
	Category      string  `gorm:"not null"`
	Gap           float32 `gorm:"not null"`
	Narration     string  `gorm:"not null"`
	Zone          string  `gorm:"not null"`
	BudgetYear    int     `gorm:"not null"`
	MonthOfPeriod string  `gorm:"not null"`
	RoleGroup     string  `gorm:"not null"`
}
