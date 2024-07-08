package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type Service interface {
	Health() map[string]string
	GetGormDB() *gorm.DB
}

type service struct {
	db     *sql.DB
	GormDB *gorm.DB
}

func New(config Config) Service {

	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.Database)
	// db, err := sql.Open("pgx", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Database,
	)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &service{
		// db:     db,
		GormDB: gormDB,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetGormDB() *gorm.DB {
	return s.GormDB
}

func CreateDB(config Config) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting db connection: %s", err.Error())
	}
	defer sqlDB.Close()

	query := fmt.Sprintf("CREATE DATABASE \"%s\"", config.Database)

	return db.Exec(query).Error
}
