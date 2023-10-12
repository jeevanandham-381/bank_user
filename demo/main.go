package main

import (
	"mymodule/api/routes"
	"mymodule/config"
	"mymodule/models"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Initialize the database
    db, err := config.InitDb()
    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }
    defer db.Close()
	//db=db.Table("mydata")
	db.AutoMigrate(&models.BankUser{})


    // Setup routes
    routes.EmpRoutes(r)

    r.Run(":8080")
}
