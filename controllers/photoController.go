package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// get photo by token user
func GetPhotos(c *gin.Context) {
	var db = database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}
	User := models.User{}

	userID := uint(userData["id"].(float64))

	Photo.UserId = userID

	err := db.Where("user_id = ?", userID).Find(&Photo).Error
	errUser := db.Where("id = ?", userID).Find(&User).Error

	if err != nil || errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if Photo.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":         Photo.Id,
			"title":      Photo.Title,
			"caption":    Photo.Caption,
			"photo_url":  Photo.Photo_url,
			"user_id":    Photo.UserId,
			"created_at": Photo.CreatedAt,
			"updated_at": Photo.UpdatedAt,
			"User": gin.H{
				"email":    User.Email,
				"username": User.UserName,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "You don't have any photo",
		})
	}
}

func GetAllPhoto(c *gin.Context) {
	var db = database.GetDB()

	var Photo []models.Photo

	err := db.Model(&models.Photo{}).Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Photos successfully!",
		"data":    Photo,
	})
}

func GetOnePhoto(c *gin.Context) {
	var db = database.GetDB()

	var Photo models.Photo

	err := db.First(&Photo, "id = ?", c.Param("photoId")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": Photo})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userId
	Photo.Id = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, Photo_url: Photo.Photo_url}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.Id,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.UserId,
		"updated_at": Photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	Photo.Id = uint(photoId)
	Photo.UserId = userId

	err := db.Model(&Photo).Where("id = ?", photoId).Delete(models.Photo{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
