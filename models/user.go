package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"unique;not null"`
	LastName  string `json:"lastName" gorm:"unique;not null"`
	Email    	string `json:"email" gorm:"unique;not null"`
	Password 	string `json:"password" gorm:"not null"`
	
	RoleID	 	uint	 `json:"role_id"`
	Role		 	Role	 `json:"role" gorm:"foreignKey:RoleID"`
}

