package main
import "github.com/gin-gonic/gin"

func main(){
	q:=gin.Default()
	q.GET("/", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"name":"kalai",
			"city":"chennai",

		})
	})
		q.Run()
		
	}

