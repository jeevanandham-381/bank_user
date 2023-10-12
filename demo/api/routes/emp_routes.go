package routes

import (
	"mymodule/api/handlers"
	//"mymodule/handlers"

	"github.com/gin-gonic/gin"
)

func EmpRoutes(r *gin.Engine) {
	// EmpRou:=r.Group("/employees")
	// {
	// 	EmpRou.GET("/",handlers.GetAllEmployees)
	// 	EmpRou.POST("/newemp",handlers.NewEmp)
	// }
	r.GET("/getalluser", handlers.GetAllEmployees)
	r.POST("/newuser", handlers.NewEmp)
	r.GET("/search", handlers.SearchByAny)
	r.PATCH("/update/:acc_no", handlers.UpdateUser)
	r.DELETE("/delete/:acc_no", handlers.DeleteUser)
	r.GET("/getalluseracc", handlers.GetAllEmployeesAcc)

}
