package migrations

import (
	model "github.com/fiscaluno/hyoga/models"
)

var Institution = model.Institution{}

func init() {
	Create(&Institution)
}
