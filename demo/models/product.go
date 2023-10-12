package models

import "github.com/jinzhu/gorm"

type BankUser struct {
	gorm.Model
	Name      string
	City      string
	Phone     int64
	AccountNo int64 `gorm:"column:acc_no"`

	AdharNo int64  `gorm:"column:adhar_no"`
	PanNo   string `gorm:"column:pan_no"`
}

// var newEmployee models.BankUser
// if err := c.ShouldBindJSON(&newEmployee); err != nil {
// 	c.JSON(http.StatusBadRequest, err.Error())
// 	return
// }
// if err := db.Create(&newEmployee).Error; err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
// 	return

// }
