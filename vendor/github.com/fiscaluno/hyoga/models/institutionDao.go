package models

import(
    "github.com/fiscaluno/hyoga/database"
    _ "github.com/jinzhu/gorm"
)

var db = database.GetInstance()
