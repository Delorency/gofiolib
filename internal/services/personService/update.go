package personservice

import "fiolib/internal/models"

func (ps *personService) Update(id uint, data *models.Person) (*models.Person, error) {
	return ps.repo.Update(id, data)
}
