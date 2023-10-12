package handlers

import (
	"fmt"

	"mymodule/config"
	"mymodule/models"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
)

func GetAllEmployees(c *gin.Context) {
	db, err := config.InitDb()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}
	defer db.Close()
	var employees []models.BankUser
	db.Find(&employees)

	// Return the employees as JSON
	c.JSON(http.StatusOK, employees)
}

// --------------------------------------------------------------------------
func NewEmp(c *gin.Context) {
	db, err := config.InitDb()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}
	defer db.Close()
	var newEmployee models.BankUser
	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := db.Create(&newEmployee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return

	}
}

//-------------------------------------------------------------------------------
//search any attribute

func SearchByAny(c *gin.Context) {
	db, err := config.InitDb()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}
	defer db.Close()
	phone := c.DefaultQuery("phone", "")
	accno := c.DefaultQuery("acc_no", "")
	adharno := c.DefaultQuery("adhar_no", "")
	panno := c.DefaultQuery("pan_no", "")

	var FilterData []models.BankUser
	//FilterData=append(FilterData, )
	if accno != "" {

		db.Where("acc_no = ?", accno).Find(&FilterData)

	} else if adharno != "" {

		db.Where("adhar_no = ?", adharno).Find(&FilterData)

	} else if panno != "" {

		db.Where("pan_no = ?", panno).Find(&FilterData)
	} else if phone != "" {

		db.Where("phone= ?", phone).Find(&FilterData)
	}
	//db.Find(&FilterData)

	c.JSON(http.StatusOK, FilterData)

}

//------------------------------------------------------------------

func UpdateUser(c *gin.Context) {
	db, err := config.InitDb()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
		return
	}
	defer db.Close()

	accNo := c.Param("acc_no")

	// Query the user by acc_no
	var existUser models.BankUser
	if err := db.Where("acc_no = ?", accNo).First(&existUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	
	var updatedUser models.BankUser
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	
	existUser.Name = updatedUser.Name
	existUser.City = updatedUser.City
	existUser.Phone = updatedUser.Phone
	existUser.AdharNo = updatedUser.AdharNo
	existUser.PanNo = updatedUser.PanNo

	

	if err := db.Save(&existUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
		return
	}

	c.JSON(http.StatusOK, existUser)
}

// ---------------------------------------------------------------------------
func DeleteUser(c *gin.Context) {
	db, err := config.InitDb()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to the database"})
		return
	}
	defer db.Close()

	accno := c.Param("acc_no")
    //accno := c.DefaultQuery("acc_no", "")

	if accno != "" {
		var userToDelete models.BankUser

		// Query the user by acc_no
		if err := db.Where("acc_no = ?", accno).First(&userToDelete).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		// Delete the user
		if err := db.Delete(&userToDelete).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "acc_no parameter is required"})
	}
}


//----------------------------------------------------------------

func GetAllEmployeesAcc(c *gin.Context) {
	db, err := config.InitDb()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}
	defer db.Close()
	var AllAcc []int64
	if err:=db.Model(&models.BankUser{}).Pluck("acc_no",&AllAcc).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"no acc"})
		return
	}
	c.JSON(http.StatusOK,AllAcc)
}


