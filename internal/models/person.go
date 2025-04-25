package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name       string `gorm:"size:128"`
	Surname    string `gorm:"size:128"`
	Patronymic string `gorm:"size:128;"`
	Age        uint8
	Gender     string `gorm:"size:48"` // male bool
	Nat        string `gorm:"size:16"`
}

func (Person) TableName() string {
	return "persons"
}
