package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email    	string `json:"email" gorm:"unique;not null"`
	Password 	string `json:"password" gorm:"not null"`
	Phone   *string `json:"phone,omitempty" binding:"omitempty,numeric"`
	Address *string `json:"address,omitempty"`
	Avatar  *string `json:"avatar,omitempty"`
	
	RoleID	 	uint	 `json:"role_id"`
	Role		 	Role	 `json:"role" gorm:"foreignKey:RoleID"`
}

