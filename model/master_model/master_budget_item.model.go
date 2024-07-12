package master_model

import "gorm.io/gorm"

type MasterBudgetItem struct {
	gorm.Model

	BudgetItem string `gorm:"not null" json:"-"`
}
