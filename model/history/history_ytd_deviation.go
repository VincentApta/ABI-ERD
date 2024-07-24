package history

import "gorm.io/gorm"

type HistoryYtdDeviation struct {
	gorm.Model

	HistoryMonthlySummaryId uint    `gorm:"not null" json:"-"`
	RealizationWorksheet    float32 `json:"-"`
	EeRealization           float32 `json:"-"`
	WorksheetForecast       float32 `json:"-"`
	ForecastRealization     float32 `json:"-"`

	HistoryMonthlySummary *HistoryMonthlySummary `json:"-"`
}
