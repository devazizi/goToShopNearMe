package main

import (
	"fmt"
	"goToShopNearMe/adapter"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello goToShopNearMe")
	var dsn string = "root:@tcp(127.0.0.1:3306)/go_to_shop_near_me?charset=utf8mb4&parseTime=True&loc=Local"
	db := adapter.InitDB(dsn)
	router(db)
}

func router(db adapter.SQL_DB) {
	router := gin.Default()

	router.GET("/home", func(c *gin.Context) {
		c.JSON(200, map[string]any{
			"status":  true,
			"message": "your request was successful",
		})
	})

	router.Run("0.0.0.0:8000")
}
