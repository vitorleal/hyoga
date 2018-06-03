package institution

import (
	models "github.com/fiscaluno/hyoga/models"
	"github.com/mitchellh/mapstructure"
)

type Institution models.Institution

// database instance
var db = models.DB

// Creates new model of institution
// Return institution
func Create(attributes map[string]interface{}) *Institution {
	var institution Institution
	mapstructure.Decode(attributes, &institution)
	return &institution
}

// Save new institution at database
func (institution *Institution) Save() {
	db.Create(&institution)
}

// Find institution by id
func Find(id int) (institution *Institution) {
	institution = &Institution{}
	db.First(institution, id)
	return
}

// All institutions
func All() (institutions []Institution) {
	db.Find(&institutions)
	return
}
