package controllers

import (
	"net/http"

	"r1estate-service/config"
	"r1estate-service/models"
	"r1estate-service/schemas"
	"r1estate-service/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
    var body schemas.UserRequest

    if err := c.ShouldBindJSON(&body); err != nil {
        utils.ResponseError(c, http.StatusBadRequest, err.Error())
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
    user := models.User{
        FirstName: body.FirstName,
        LastName:  body.LastName,
        Email:     body.Email,
        Password:  string(hashedPassword),
        RoleID:    body.RoleID,
    }
    if err := config.DB.Create(&user).Error; err != nil {
        utils.ResponseError(c, http.StatusNotFound,"Failed to create user")
        return
    }
    var createdUser models.User
    if err := config.DB.Preload("Role").First(&createdUser, user.ID).Error; err != nil {
        utils.ResponseError(c, http.StatusNotFound, "Failed to load role: "+err.Error())
        return
    }
    utils.RespondJSON(c, http.StatusCreated, createdUser)
}


func GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Preload("Role").Find(&users).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.First(&user, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(c, http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var body schemas.UserRequest
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "User not found")
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Email = body.Email
	user.Password = body.Password
	user.RoleID = body.RoleID

	if err := config.DB.Save(&user).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, user)
}

func SignIn(c *gin.Context) {
	var body schemas.SignInRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	var user models.User
	if err := config.DB.Preload("Role").Where("email = ?", body.Email).First(&user).Error; err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}
	token, err := utils.GenerateJWT(user.ID, user.Role.RoleName)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusOK, gin.H{"token": token, "role": user.Role.RoleName})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.First(&user, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "User not found")
		return
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}