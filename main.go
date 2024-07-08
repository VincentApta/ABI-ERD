package main

import (
	"flag"
	"fmt"
	"os"
	"pertamina_abi/database"
	"pertamina_abi/model/master_model"
	narration_model "pertamina_abi/model/narration"
	"pertamina_abi/model/period_model"
	prognosa_model "pertamina_abi/model/prognosa"
	rkap_model "pertamina_abi/model/rkap"
	"pertamina_abi/server"

	"github.com/joho/godotenv"
)

func main() {
	flag.Parse()
	godotenv.Load(".env")

	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go up | down")
		os.Exit(1)
	}

	fmt.Println("Initiating Migration")

	switch os.Args[1] {
	case "up":
		up()
	case "down":
		down()
	default:
		fmt.Println("Usage: main.go up | down")
		os.Exit(1)
	}

}

func up() {
	serverName := "DB Migration" // Change server name accordingly
	dbConfig := database.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	// database.CreateDB(dbConfig)

	server := server.New(serverName, dbConfig)
	db := server.DB.GetGormDB()

	db.AutoMigrate(

		//Master
		&master_model.MasterAp{},
		&master_model.MasterRegion{},
		&master_model.MasterZone{},
		&master_model.MasterWk{},
		&master_model.MasterField{},
		&master_model.MasterRkapClassification{},
		&master_model.MasterActivityType{},
		&master_model.MasterBusinessStream{},
		&master_model.MasterCategory{},
		&master_model.MasterFollowUpStatus{},
		&master_model.MasterSbtConnection{},
		&master_model.MasterThematicArea{},
		&master_model.MasterMonthOfPeriod{},

		//Period
		&period_model.BudgetYear{},
		&period_model.ActualsForecastPeriod{},
		&period_model.NarrationPeriod{},

		//RKAP
		&rkap_model.BudgetItemDetail{},

		//Actuals Forecast
		&prognosa_model.ActualsForecastData{},

		//Narration
		&narration_model.ActualsDeviationNarration{},
		&narration_model.ForecastNarration{},
		&narration_model.FidStatusNarration{},
	)

	db.Exec("ALTER TABLE budget_item_details ADD CONSTRAINT unique_no_wbs UNIQUE NULLS DISTINCT(no_wbs)")
	print("Database Migrated")
}

func down() {
	serverName := "DB Migration" // Change server name accordingly
	dbConfig := database.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
	server := server.New(serverName, dbConfig)
	db := server.DB.GetGormDB()

	db.Migrator().DropTable(
		//Narration
		&narration_model.FidStatusNarration{},
		&narration_model.ForecastNarration{},
		&narration_model.ActualsDeviationNarration{},

		//Actuals Forecast
		&prognosa_model.ActualsForecastData{},

		//RKAP
		&rkap_model.BudgetItemDetail{},

		//Period
		&period_model.NarrationPeriod{},
		&period_model.ActualsForecastPeriod{},
		&period_model.BudgetYear{},

		//Master
		&master_model.MasterMonthOfPeriod{},
		&master_model.MasterThematicArea{},
		&master_model.MasterSbtConnection{},
		&master_model.MasterFollowUpStatus{},
		&master_model.MasterCategory{},
		&master_model.MasterBusinessStream{},
		&master_model.MasterActivityType{},
		&master_model.MasterRkapClassification{},
		&master_model.MasterField{},
		&master_model.MasterWk{},
		&master_model.MasterZone{},
		&master_model.MasterRegion{},
		&master_model.MasterAp{},
	)
	print("Database Dropped")
}
