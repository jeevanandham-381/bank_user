package main

import (
	//"fmt"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	//ID   uint
	Name string
	City string
}

var db *gorm.DB

func main() {
	// Initialize Gin
	r := gin.Default()
	// Initialize GORM and open a connection to the PostgreSQL database
	var err error
	db, err = gorm.Open("postgres", "user=postgres password=root dbname=demo sslmode=disable")
	if err != nil {
		panic("Failed to connect to the database")
	}
	defer db.Close()

	// AutoMigrate will create the table if it doesn't exist
	db.AutoMigrate(&Employee{})

	r.GET("/getall", GetAllEmployees)
	r.POST("/newemp", Postemp)
	r.GET("/getemp/:id", Getempbyid)
	r.DELETE("/del/:id", DeleteById)
	r.PATCH("/update/:id", UpdateById)
	r.GET("/getemp", GetempByCity)
	r.GET("/searchanykey", SearchByAny)

	r.Run(":8080")
}
//--------------------------------------------
func GetAllEmployees(c *gin.Context) {
	var employees []Employee
	db.Find(&employees)

	// Return the employees as JSON
	c.JSON(http.StatusOK, employees)
}
//------------------------------------------
func Postemp(c *gin.Context) {
	var newEmployee Employee
	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := db.Create(&newEmployee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return

	}
	c.JSON(http.StatusCreated, newEmployee)

}
//--------------------------------------
func Getempbyid(c *gin.Context) {

	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var emp Employee
	if err := db.First(&emp, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
	}
	c.JSON(http.StatusOK, emp)

}
//---------------------------------------------
func DeleteById(c *gin.Context) {
	empId := c.Param("id")
	id, err := strconv.ParseUint(empId, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	//var emp Employee
	if err := db.Where("id=?", id).Delete(&Employee{}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})

}
//--------------------------------------------------------
func UpdateById(c *gin.Context) {
	id := c.Param("id")
	eid, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var Updemp Employee

	if err := c.ShouldBindJSON(&Updemp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//exisiting data
	var ExistData Employee
	if err := db.First(&ExistData, eid).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusInternalServerError, gin.H{"messge": "employee not found"})
			return
		}
	}
	ExistData.Name = Updemp.Name
	ExistData.City = Updemp.City

	if err := db.Save(&ExistData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
		return

	}
	c.JSON(http.StatusOK, ExistData)

}
//---------------------------------------------------------

func GetempByCity(c *gin.Context) {
	city := c.DefaultQuery("city", "")
	var employees []Employee
	if city != "" {
		db.Where("city=?", city).Find(&employees)
	}
	c.JSON(http.StatusOK, employees)
}

//searching in one api

func SearchByAny(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	name := c.DefaultQuery("name", "")
	city := c.DefaultQuery("city", "")

	var FilterData []Employee
	if id != "" {
		var emp Employee
		if err := db.First(&emp, id).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found "})
				return

			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "query error"})
			return
		}
		FilterData = append(FilterData, emp)
	} else if name != "" && city != "" {

		db.Where("name = ?", name).Where("city=?", city).Find(&FilterData)

	} else if city != "" {

		db.Where("city = ?", city).Find(&FilterData)
	} else if name != "" {

		db.Where("name = ?", name).Find(&FilterData)
	}
	//db.Find(&FilterData)

	c.JSON(http.StatusOK, FilterData)

}
