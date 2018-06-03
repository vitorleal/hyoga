package migrations

import (
	"fmt"
	"github.com/fiscaluno/hyoga/database"
	log "github.com/fiscaluno/hyoga/fiscalog"
	"reflect"
	"regexp"
)

// Database Instance
var db = database.GetInstance()

func init() {
	log.Info("Migrating database... Wait a moment please")
}

func Migrate() {
	log.Info("Database migrated")
}

// Create new Schema to database
func Create(Schema interface{}) {
	pattern := regexp.MustCompile(`\.(\w+)$`)
	schemaString := reflect.TypeOf(Schema).String()
	tableName := pattern.FindStringSubmatch(schemaString)[1]

	if !db.HasTable(Schema) {
		log.Info(fmt.Sprintf("Creating %s table...", tableName))
		defer log.Info(fmt.Sprintf("%s table created", tableName))
		db.AutoMigrate(Schema)
	} else {
		log.Info(fmt.Sprintf("%s table exists...", tableName))
	}
}
