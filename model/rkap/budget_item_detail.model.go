package rkap_model

import (
	"pertamina_abi/model/master_model"
	"pertamina_abi/model/period_model"
	"time"

	"gorm.io/gorm"
)

type BudgetItemDetail struct {
	gorm.Model

	BudgetYearId              uint      `gorm:"not null" json:"-"`
	Code                      string    `gorm:"not null" json:"-"`
	ApId                      uint      `gorm:"not null" json:"-"`
	RegionId                  uint      `gorm:"not null" json:"-"`
	WkId                      uint      `gorm:"not null" json:"-"`
	ZoneId                    uint      `gorm:"not null" json:"-"`
	FieldId                   uint      `gorm:"not null" json:"-"`
	WorksheetClassificationId uint      `gorm:"not null" json:"-"`
	F00Number                 string    `json:"-"`
	FidNumber                 string    `json:"-"`
	FidStatusId               uint      `gorm:"not null" json:"-"`
	FidTarget                 time.Time `gorm:"type:date" json:"-"`
	TresholdId                uint      `gorm:"not null" json:"-"`
	WbsNumber                 string    `json:"-"`
	BudgetItemId              uint      `gorm:"not null" json:"-"`
	WorkPlan                  string    `gorm:"not null" json:"-"`
	ActivityTypeId            uint      `gorm:"not null" json:"-"`
	BusinessStreamId          uint      `gorm:"not null" json:"-"`
	CategoryId                uint      `gorm:"not null" json:"-"`
	PriorityId                uint      `gorm:"not null" json:"-"`
	StageId                   uint      `gorm:"not null" json:"-"`
	StartYear                 int       `gorm:"not null" json:"-"`
	EndYear                   int       `gorm:"not null" json:"-"`
	CauseOfDelay              string    `json:"-"`
	ActualPhysicalProgressId  uint      `gorm:"not null" json:"-"`
	EoyPhysicalProjectionId   uint      `gorm:"not null" json:"-"`
	RealizedLastYear          float32   `json:"-"`
	TotalFidInvestment        uint64    `json:"-"`
	StatusOfForecast          string    `json:"-"`
	StatusNotes               string    `json:"-"`
	InvestmentActivityUpdate  string    `json:"-"`
	RealizationGapNotes       string    `json:"-"`
	ObstacleCategoryId        uint      `json:"-"`
	ActionPlan                string    `json:"-"`
	FollowUpStatusId          uint      `json:"-"`
	SbtConnectionId           uint      `json:"-"`
	ThematicAreaId            uint      `json:"-"`

	BudgetYear              *period_model.BudgetYear                    `json:"-"`
	MasterMonthOfPeriod     *master_model.MasterMonthOfPeriod           `json:"-"`
	Ap                      *master_model.MasterAp                      `json:"-"`
	Region                  *master_model.MasterRegion                  `json:"-"`
	Wk                      *master_model.MasterWk                      `json:"-"`
	Zone                    *master_model.MasterZona                    `json:"-"`
	Field                   *master_model.MasterField                   `json:"-"`
	WorksheetClassification *master_model.MasterWorksheetClassification `json:"-"`
	FidStatus               *master_model.MasterFidStatus               `json:"-"`
	Treshold                *master_model.MasterTreshold                `json:"-"`
	BudgetItem              *master_model.MasterBudgetItem              `json:"-"`
	ActivityType            *master_model.MasterActivityType            `json:"-"`
	BusinessStream          *master_model.MasterBusinessStream          `json:"-"`
	Category                *master_model.MasterCategory                `json:"-"`
	Priority                *master_model.MasterPriority                `json:"-"`
	Stage                   *master_model.MasterStage                   `json:"-"`
	ActualPhysicalProgress  *master_model.ActualPhysicalProgress        `json:"-"`
	EoyPhysicalProjection   *master_model.EoyPhysicalProjection         `json:"-"`
	ObstacleCategory        *master_model.MasterObstacleCategory        `json:"-"`
	FollowUpStatus          *master_model.MasterFollowUpStatus          `json:"-"`
	SbtConnection           *master_model.MasterSbtConnection           `json:"-"`
	ThematicArea            *master_model.MasterThematicArea            `json:"-"`
}
