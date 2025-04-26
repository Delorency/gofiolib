package persondb

import "fiolib/internal/models"

func (r *personDB) Create(data *models.Person) error {
	return r.db.Create(data).Error
}
