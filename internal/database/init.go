package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DBServ struct {
	*gorm.DB
	SQL      *sql.DB
	Host     string
	Port     string
	User     string
	Password string
	Timezone string
	Name     string
}

func (db *DBServ) sqlDB() (*sql.DB, error) {

	return db.DB.DB()
}

func (db *DBServ) Close() error {

	return db.SQL.Close()
}

var (
	DB      *DBServ
	Schemas map[string]interface{}
)

func init() {

	DB = &DBServ{
		DB:       nil,
		SQL:      nil,
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Timezone: "Europe/Moscow",
		Name:     os.Getenv("POSTGRES_DB"),
	}

	//DB = &DBServ{
	//	DB:       nil,
	//	SQL:      nil,
	//	Host:     "localhost",
	//	Port:     "5432",
	//	User:     "postgres",
	//	Password: "password",
	//	Timezone: "Europe/Moscow",
	//	Name:     "test",
	//}

	Schemas = map[string]interface{}{
		"task":       Task{},
		"image_data": ImageData{},
		"face_data":  FaceData{},
		"statistic":  Statistic{},
	}

	createDatabase()
	createDatatypes()
	createTables()
}

func createDatabase() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		DB.Host, DB.Port, DB.User, DB.Password, DB.Name, DB.Timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s",
			DB.Host, DB.Port, DB.User, DB.Password, DB.Timezone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
		if err != nil {
			log.Println(err)
		}

		db.Exec(
			fmt.Sprintf("CREATE DATABASE %s", DB.Name),
		)
		sqlDB, err := DB.sqlDB()
		if err != nil {
			log.Fatalln(err)
		}

		if err = sqlDB.Close(); err != nil {
			log.Fatalln(err)
		}

		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			DB.Host, DB.Port, DB.User, DB.Password, DB.Name, DB.Timezone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
		if err != nil {
			log.Fatalln(err)
		}
	}

	DB.DB = db
	sqlDB, err := DB.sqlDB()
	if err != nil {
		log.Fatalln(err)
	}
	DB.SQL = sqlDB
}

func createDatatypes() {

	if _, err := DB.SQL.Exec("CREATE TYPE status AS ENUM ('forming', 'processing', 'completed', 'error');"); err != nil {
		log.Println(err)
	}
	log.Println("type -----status----- successfully created")

	if _, err := DB.SQL.Exec("CREATE TYPE image_status AS ENUM ('untouched', 'processed', 'error');"); err != nil {
		log.Println(err)
	}
	log.Println("type -----image_status----- successfully created")

	if _, err := DB.SQL.Exec("CREATE TYPE sex AS ENUM ('male', 'female');"); err != nil {
		log.Println(err)
	}
	log.Println("type -----sex----- successfully created")

}

func createTables() {

	for name, schema := range Schemas {
		if !DB.Migrator().HasTable(&schema) {
			if err := DB.Migrator().CreateTable(&schema); err != nil {
				log.Fatalln(err)
			}
			log.Printf("table -----%s----- successfully created", name)
		}
	}

}
