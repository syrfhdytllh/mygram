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

func CreateSocialmedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserId = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

// get social media by token user
func GetSocialMedias(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}
	User := models.User{}

	userID := uint(userData["id"].(float64))

	SocialMedia.UserId = userID

	err := db.Where("user_id = ?", userID).Find(&SocialMedia).Error
	errUser := db.Where("id = ?", userID).Find(&User).Error

	if err != nil || errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if SocialMedia.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"social_media": gin.H{
				"id":               SocialMedia.Id,
				"name":             SocialMedia.Name,
				"social_media_url": SocialMedia.Social_media_url,
				"user_id":          SocialMedia.UserId,
				"created_at":       SocialMedia.CreatedAt,
				"updated_at":       SocialMedia.UpdatedAt,
				"User": gin.H{
					"id":       User.Id,
					"username": User.UserName,
				},
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "You don't have any data",
		})
	}
}

func GetAllSocialMedia(c *gin.Context) {
	var db = database.GetDB()

	var SocialMedia []models.SocialMedia

	err := db.Model(&models.SocialMedia{}).Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Social Media successfully!",
		"data":    SocialMedia,
	})
}

func GetOneSocialMedia(c *gin.Context) {
	var db = database.GetDB()

	var SocialMedia models.SocialMedia

	err := db.First(&SocialMedia, "id = ?", c.Param("socialMediaId")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": SocialMedia})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserId = userId
	SocialMedia.Id = uint(socialMediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{Name: SocialMedia.Name, Social_media_url: SocialMedia.Social_media_url}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.Id,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.Social_media_url,
		"user_id":          SocialMedia.UserId,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userId := uint(userData["id"].(float64))

	SocialMedia.Id = uint(socialMediaId)
	SocialMedia.UserId = userId

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Delete(models.SocialMedia{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
