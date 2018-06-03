package database

import (
	"fmt"
	"github.com/fiscaluno/hyoga/config"
	log "github.com/fiscaluno/hyoga/fiscalog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	if db == nil {
		var err interface{}
		fmt.Println(getConnectionString())
		db, err = gorm.Open("postgres", getConnectionString())

		if err != nil {
			log.Error(err)
		}
	}

	return db
}

func getConnectionString() (connect string) {
	host := config.DB_HOST
	dbname := config.DB_NAME
	user := config.DB_USER
	pass := config.DB_PASS
	port := config.DB_PORT
	connect = fmt.Sprintf(
			"host=%s user=%s password=%s port=%s dbname=%s",
			host,
			user,
			pass,
			port,
			dbname,
		)
	return
}
