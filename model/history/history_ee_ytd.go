package history

import "gorm.io/gorm"

type HistoryEeYtd struct {
	gorm.Model

	HistoryMonthlySummaryId uint    `gorm:"not null" json:"-"`
	EeYtdJan                float32 `json:"-"`
	EeYtdFeb                float32 `json:"-"`
	EeYtdMar                float32 `json:"-"`
	EeYtdApr                float32 `json:"-"`
	EeYtdMay                float32 `json:"-"`
	EeYtdJun                float32 `json:"-"`
	EeYtdJul                float32 `json:"-"`
	EeYtdAug                float32 `json:"-"`
	EeYtdSep                float32 `json:"-"`
	EeYtdOct                float32 `json:"-"`
	EeYtdNov                float32 `json:"-"`
	EeYtdDec                float32 `json:"-"`

	HistoryMonthlySummary *HistoryMonthlySummary `json:"-"`
}
