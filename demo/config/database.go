package config


import (
	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

var db *gorm.DB
var err error

func InitDb() (*gorm.DB, error) {
    db, err = gorm.Open("postgres", "user=postgres password=root dbname=jerry sslmode=disable")
    if err != nil {
        return nil, err
    }
    return db, nil

    
}


	




