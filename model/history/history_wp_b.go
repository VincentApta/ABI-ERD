package history

import "gorm.io/gorm"

type HistoryWpB struct {
	gorm.Model

	HistoryMonthlySummaryId uint   `gorm:"not null" json:"-"`
	AfeNumber               string `json:"-"`
	AfeStatus               string `json:"-"`
	RkTitle                 string `json:"-"`
	AfeValue                int    `json:"-"`
	Tangible                int    `json:"-"`
	Intangible              int    `json:"-"`
	TangibleCurrentYear     int    `json:"-"`
	IntangibleCurrentYear   int    `json:"-"`
	AllocationCurrentYear   int    `json:"-"`
	BudgetSchedule          int    `json:"-"`
	Line                    int    `json:"-"`

	HistoryMonthlySummary *HistoryMonthlySummary `json:"-"`
}
