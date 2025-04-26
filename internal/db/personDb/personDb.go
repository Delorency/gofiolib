package persondb

import (
	"fiolib/internal/models"

	"gorm.io/gorm"
)

type PersonDBI interface {
	// List(*schemes.Pagination) (*[]models.Group, error)
	Create(*models.Person) error
	// Update(uint, *models.Group) (*models.Group, error)
	// Retrieve(uint) (*models.Group, error)
}

type personDB struct {
	db *gorm.DB
}

func NewPersonDB(db *gorm.DB) PersonDBI {
	return &personDB{db}
}
