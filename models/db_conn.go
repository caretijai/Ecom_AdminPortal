package models

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

//PostgresDB structure implements the database interface.
type PostgresDB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
	Handler  *gorm.DB
}

// GetDB gets the database reference
func GetDBConnection() (PostgresDB, error) {
	// Init DB connection
	pgdb := PostgresDB{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBname:   os.Getenv("POSTGRES_DB"),
		Port:     5432,
	}
	// pgdb.Port, _ = strconv.Atoi(os.Getenv("POSTGRES_DB")) //postgres port
	var err error
	pgdb.Handler, err = gorm.Open("postgres", pgdb.PsqlInfo())
	if err != nil {
		log.Printf("db_connection err: %v", err)
		panic(err)
	}
	// defer pgdb.Handler.Close()
	return pgdb, nil
}

//IsRecordNotFoundError returns an error if record is not found
func (pgdb *PostgresDB) IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
