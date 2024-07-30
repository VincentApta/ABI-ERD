package history

import "gorm.io/gorm"

type HistoryMonthlySummary struct {
	gorm.Model

	BudgetYear              int8    `gorm:"not null" json:"-"`
	MonthOfPeriod           string  `gorm:"not null" json:"-"`
	Code                    string  `gorm:"not null" json:"-"`
	Ap                      string  `gorm:"not null" json:"-"`
	Region                  string  `gorm:"not null" json:"-"`
	Wk                      string  `gorm:"not null" json:"-"`
	Zone                    string  `gorm:"not null" json:"-"`
	Field                   string  `gorm:"not null" json:"-"`
	WorksheetClassification string  `gorm:"not null" json:"-"`
	F00Number               string  `json:"-"`
	FidNumber               string  `json:"-"`
	FidStatus               string  `gorm:"not null" json:"-"`
	FidTarget               string  `json:"-"`
	Treshold                string  `gorm:"not null" json:"-"`
	WbsNumber               string  `json:"-"`
	BudgetItem              string  `gorm:"not null" json:"-"`
	WorkPlan                string  `gorm:"not null" json:"-"`
	ActivityType            string  `gorm:"not null" json:"-"`
	BusinessStream          string  `gorm:"not null" json:"-"`
	Category                string  `gorm:"not null" json:"-"`
	Priority                string  `gorm:"not null" json:"-"`
	Stage                   string  `gorm:"not null" json:"-"`
	StartYear               int     `gorm:"not null" json:"-"`
	EndYear                 int     `gorm:"not null" json:"-"`
	CauseOfDelay            string  `json:"-"`
	ActualPhysicalProgress  string  `gorm:"not null" json:"-"`
	EoyPhysicalProjection   string  `gorm:"not null" json:"-"`
	RealizedLastYear        float32 `json:"-"`
	TotalFidInvestment      uint64  `json:"-"`
	FollowUpStatus          string  `json:"-"`
	SbtConnection           string  `json:"-"`
	ThematicArea            string  `json:"-"`
}
