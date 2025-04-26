package persondb

import "fiolib/internal/models"

func (pd *personDB) Create(data *models.Person) error {
	return pd.db.Create(data).Error
}
