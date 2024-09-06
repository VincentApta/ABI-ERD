package history

import "gorm.io/gorm"

type HistoryActualDeviationNarration struct {
	gorm.Model
	Code          string  `gorm:"not null"`
	BudgetItem    string  `gorm:"not null"`
	Gap           float32 `gorm:"not null"`
	Narration     string  `gorm:"not null"`
	Category      string  `gorm:"not null"`
	Zone          string  `gorm:"not null"`
	BudgetYear    int     `gorm:"not null"`
	MonthOfPeriod string  `gorm:"not null"`
	RoleGroup     string  `gorm:"not null"`
}
