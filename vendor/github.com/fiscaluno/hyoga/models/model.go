package models

import (
	"github.com/fiscaluno/hyoga/database"
)

var DB = database.GetInstance()

// Model contract
type Model interface {
	All()
	Find()
	Save()
	Delete()
	Update()
}
