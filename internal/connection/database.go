package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/muhamadnirwansyah/authentication-service/internal/config"
)

func GetDatabase(configParam config.Database) *sql.DB {
	log.Printf("host=%s , user=%s password=%s dbname=%s port=%s", configParam.Host, configParam.Name, configParam.Pass, configParam.Name, configParam.Port)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		configParam.Host,
		configParam.User,
		configParam.Pass,
		configParam.Name,
		configParam.Port,
		configParam.Tz,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed open connection to database and process migrate : ", err.Error())
	}

	return db
}
