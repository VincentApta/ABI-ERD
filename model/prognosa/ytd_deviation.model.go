package prognosa_model

import (
	rkap_model "pertamina_abi/model/rkap"

	"gorm.io/gorm"
)

type YtdDeviation struct {
	gorm.Model

	BudgetItemDetailId uint `gorm:"not null" json:"-"`
	

	BudgetItemDetail *rkap_model.BudgetItemDetail `json:"-"`
}
