package personservice

import persondb "fiolib/internal/db/personDb"

type PersonServiceI interface {
	// List(*schemes.Pagination) (*[]models.Group, error)
	// Create(*models.Group) error
	// Update(uint, *models.Group) (*models.Group, error)
	// Retrieve(uint) (*models.Group, error)
}

type personService struct {
	repo persondb.PersonDBI
}

func NewPersonService(repo persondb.PersonDBI) PersonServiceI {
	return &personService{repo}
}
