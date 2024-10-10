package main

import (
	"log"
	db "main/database"
	"main/handlers"
	"main/models"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	db.DB.AutoMigrate(&models.Basket{})

	db.DB.AutoMigrate(&models.User{})

	e := echo.New()

	// Routes
	basketRoute := e.Group("/basket")

	basketRoute.GET("/", handlers.Get_baskets)
	basketRoute.POST("/", handlers.Create_basket)
	basketRoute.GET("/:id", handlers.Get_basket_by_id)
	basketRoute.PATCH("/:id", handlers.Update_basket)
	basketRoute.DELETE("/:id", handlers.Delete_basket_by_id)

	userRoute := e.Group("/user")

	userRoute.GET("/", handlers.Get_all_users)
	userRoute.GET("/:id", handlers.Get_user_by_id)
	userRoute.POST("/", handlers.Create_user)
	userRoute.DELETE("/:id", handlers.Delete_user_by_id)
	userRoute.PATCH("/:id", handlers.ChangePassword)
	log.Fatal(e.Start("0.0.0.0:5432"))
}
