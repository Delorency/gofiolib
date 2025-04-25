package container

import (
	persondb "fiolib/internal/db/personDb"
	personservice "fiolib/internal/services/personService"

	"gorm.io/gorm"
)

type Container struct {
	// Репозитории
	PersonRepo persondb.PersonDBI

	// Сервисы
	PersonService personservice.PersonServiceI
}

func NewContainer(db *gorm.DB) *Container {
	personrepo := persondb.NewPersonDB(db)

	personserv := personservice.NewPersonService(personrepo)

	return &Container{
		PersonRepo:    personrepo,
		PersonService: personserv,
	}
}
