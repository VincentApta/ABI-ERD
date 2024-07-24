package history

import "gorm.io/gorm"

type HistoryDdRealization struct {
	gorm.Model

	HistoryMonthlySummaryId uint    `gorm:"not null" json:"-"`
	DdRealizationJan        float32 `json:"-"`
	DdRealizationFeb        float32 `json:"-"`
	DdRealizationMar        float32 `json:"-"`
	DdRealizationApr        float32 `json:"-"`
	DdRealizationMay        float32 `json:"-"`
	DdRealizationJun        float32 `json:"-"`
	DdRealizationJul        float32 `json:"-"`
	DdRealizationAug        float32 `json:"-"`
	DdRealizationSep        float32 `json:"-"`
	DdRealizationOct        float32 `json:"-"`
	DdRealizationNov        float32 `json:"-"`
	DdRealizationDec        float32 `json:"-"`

	HistoryMonthlySummary *HistoryMonthlySummary `json:"-"`
}
