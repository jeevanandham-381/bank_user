package main

import (
	//"net/http"
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

type emplyee struct {
	ID     uint
	Name   string
	Salary float64
	City   string
}

var emp = []emplyee{
	{ID: 1, Name: "kalai", Salary: 10000, City: "chennai"},
	{ID: 2, Name: "ajith", Salary: 15000, City: "vellore"},
	{ID: 3, Name: "vasim", Salary: 20000, City: "kovao"},
}

func Getemp(c *gin.Context) {
	c.IndentedJSON(200, emp)
}

func Postemp(c *gin.Context) {
	var newemp emplyee
	if err := c.BindJSON(&newemp); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	emp = append(emp, newemp)
	c.IndentedJSON(200, emp)
}

func GetelmID(c *gin.Context){
	idStr := c.Param("id")
    id,err := strconv.ParseUint(idStr, 10,0)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
	}
	for _,a:=range emp{
		if a.ID==uint(id){
			c.IndentedJSON(http.StatusOK,a)
			return
		}
		

	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})

}

func main() {
	fmt.Printf("%T", emp)
	r := gin.Default()
	r.GET("/", Getemp)
	r.GET("/getemp/:id", GetelmID)
	r.POST("/new", Postemp)
	r.Run("localhost:8080")

}
