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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.Id,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"created_at": Comment.CreatedAt,
	})
}

// get comment by user token
func GetComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	User := models.User{}
	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	Photo.UserId = userID
	User.Id = userID
	Comment.UserId = userID

	err := db.Where("user_id = ?", userID).Find(&Comment).Error
	errUser := db.Where("id = ?", userID).Find(&User).Error
	errPhoto := db.Where("user_id = ?", userID).Find(&Photo).Error

	if err != nil || errUser != nil || errPhoto != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.Id,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
		"created_at": Comment.CreatedAt,
		"User": gin.H{
			"id":       User.Id,
			"email":    User.Email,
			"username": User.UserName,
		},
		"Photo": gin.H{
			"id":        Photo.Id,
			"title":     Photo.Title,
			"caption":   Photo.Caption,
			"photo_url": Photo.Photo_url,
			"user_id":   Photo.UserId,
		},
	})
}

func GetAllComments(c *gin.Context) {
	var db = database.GetDB()

	var Comment []models.Comment

	err := db.Model(&models.Comment{}).Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Comment successfully!",
		"data":    Comment,
	})
}

func GetOneComment(c *gin.Context) {
	var db = database.GetDB()

	var Comment models.Comment

	err := db.First(&Comment, "id = ?", c.Param("commentId")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": Comment})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.Id = uint(commentId)
	Comment.UserId = uint(userId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.Id,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))

	Comment.Id = uint(commentId)
	Comment.UserId = userId

	err := db.Model(&Comment).Where("id = ?", commentId).Delete(models.Comment{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
