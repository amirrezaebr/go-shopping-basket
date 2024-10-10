package handlers

import (
	db "main/database"
	"main/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func Create_basket(c echo.Context) error {
	new_basket := new(models.Basket)
	new_basket.Created_at = time.Now()
	new_basket.Updated_at = time.Now()
	new_basket.State = "Pending"

	if err := c.Bind(new_basket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := db.DB.Create(&new_basket).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	//new_basket.userId = userID
	return c.JSON(http.StatusCreated, new_basket)
}

func Get_baskets(c echo.Context) error {
	var baskets []models.Basket
	if err := db.DB.Find(&baskets).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, baskets)
}

func Update_basket(c echo.Context) error {

	basket_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Basket ID"})
	}

	var basket = new(models.Basket)
	if err := db.DB.First(&basket, basket_id).Error; err != nil { //searching the basket
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Basket not found"})
	}

	if basket.State == "Completed" { //checking the status
		return c.JSON(http.StatusNotFound, map[string]string{"error": "State is Completed"})
	}

	if err := c.Bind(&basket); err != nil { //replacing the object
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	if err := db.DB.Save(&basket).Error; err != nil { //saving the new obj in the DB
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, basket)
}

func Get_basket_by_id(c echo.Context) error {
	basket_id, err := strconv.Atoi(c.Param("id"))
	var basket models.Basket

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid id"})
	}

	if err := db.DB.Where("id = ?", basket_id).Find(&basket).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Basket not found"})
	}

	return c.JSON(http.StatusOK, basket)
}

func Delete_basket_by_id(c echo.Context) error {
	basket_id, err := strconv.Atoi(c.Param("id"))
	var basket models.Basket

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid id"})
	}

	if err := db.DB.First(&basket, basket_id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Item with id: " + strconv.Itoa(basket_id) + " Deleted from basket Successfully"})
}
