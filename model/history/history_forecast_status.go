package history

import "gorm.io/gorm"

type HistoryForecastStatus struct {
	gorm.Model

	HistoryMonthlySummaryId uint    `gorm:"not null" json:"-"`
	ForecastStatusJan       float32 `json:"-"`
	ForecastStatusFeb       float32 `json:"-"`
	ForecastStatusMar       float32 `json:"-"`
	ForecastStatusApr       float32 `json:"-"`
	ForecastStatusMay       float32 `json:"-"`
	ForecastStatusJun       float32 `json:"-"`
	ForecastStatusJul       float32 `json:"-"`
	ForecastStatusAug       float32 `json:"-"`
	ForecastStatusSep       float32 `json:"-"`
	ForecastStatusOct       float32 `json:"-"`
	ForecastStatusNov       float32 `json:"-"`
	ForecastStatusDec       float32 `json:"-"`

	HistoryMonthlySummary *HistoryMonthlySummary `json:"-"`
}
