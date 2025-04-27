package personservice

import (
	persondb "fiolib/internal/db/personDb"
	"fiolib/internal/models"
	schemes "fiolib/internal/schemes"
)

type PersonServiceI interface {
	List(*schemes.Pagination, *schemes.Filter) (*[]models.Person, error)
	Create(*models.Person) error
	Update(uint, *models.Person) (*models.Person, error)
	Retrieve(uint) (*models.Person, error)
	Delete(uint) error
}

type personService struct {
	repo persondb.PersonDBI
}

func NewPersonService(repo persondb.PersonDBI) PersonServiceI {
	return &personService{repo}
}
