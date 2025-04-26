package persondb

import (
	"fiolib/internal/models"

	schemes "fiolib/internal/schemes"

	"gorm.io/gorm"
)

type PersonDBI interface {
	List(*schemes.Pagination) (*[]models.Person, error)
	Create(*models.Person) error
	Update(uint, *models.Person) (*models.Person, error)
	Retrieve(uint) (*models.Person, error)
	Delete(uint) error
}

type personDB struct {
	db *gorm.DB
}

func NewPersonDB(db *gorm.DB) PersonDBI {
	return &personDB{db}
}
