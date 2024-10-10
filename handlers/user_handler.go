package handlers

import (
	db "main/database"
	"main/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Get_all_users (c echo.Context) error {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)

}

func Create_user(c echo.Context) error {
	new_user := new(models.User)

	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := db.DB.Create(&new_user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, new_user)
}

func Delete_user_by_id(c echo.Context) error {
	user_id, err := strconv.Atoi(c.Param("id"))
	var user models.User

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid id"})
	}

	if err := db.DB.Delete("User_id = ?", user_id).Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})

}

func Get_user_by_id(c echo.Context) error {
	user_id, err := strconv.Atoi(c.Param("id"))
	var user models.User
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "id not found"})
	}

	if err := db.DB.Where("User_id = ?", user_id).Find(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}
