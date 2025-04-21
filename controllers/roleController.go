package controllers

import (
	"net/http"
	"r1estate-service/config"
	"r1estate-service/models"
	"r1estate-service/schemas"
	"r1estate-service/utils"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var body schemas.RoleRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	role := models.Role{RoleName: body.RoleName}
	if err := config.DB.Create(&role).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusCreated, role)
}

func GetRoles(c *gin.Context) {
	var roles []models.Role
	
	if err := config.DB.Find(&roles).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusOK, roles)
}

func GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Role not found")
		return
	}
	utils.RespondJSON(c, http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Role not found")
		return
	}

	var input schemas.RoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	role.RoleName = input.RoleName
	if err := config.DB.Save(&role).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Role not found")
		return
	}
	config.DB.Delete(&role)
	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Role deleted successfully"})
}