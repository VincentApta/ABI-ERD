package history

import "gorm.io/gorm"

type HistoryActualForecastNotes struct {
	gorm.Model

	HistoryMonthlySummaryId  uint   `gorm:"not null" json:"-"`
	StatusOfForecast         string `json:"-"`
	StatusNotes              string `json:"-"`
	InvestmentActivityUpdate string `json:"-"`
	RealizationGapNotes      string `json:"-"`
	ObstacleCategory         string `json:"-"`
	ActionPlan               string `json:"-"`

	HistoryMonthlySummary *HistoryMonthlySummary `json:"-"`
}
