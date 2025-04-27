package personservice

import (
	"fiolib/internal/models"
	schemes "fiolib/internal/schemes"
)

func (ps *personService) List(p *schemes.Pagination, f *schemes.Filter) (*[]models.Person, error) {
	return ps.repo.List(p, f)
}
func (ps *personService) Retrieve(id uint) (*models.Person, error) {
	return ps.repo.Retrieve(id)
}
