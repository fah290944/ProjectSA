package entity

import (

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`

	Schedule []Schedule  `gorm:"foreignKey:DoctorID"`

}