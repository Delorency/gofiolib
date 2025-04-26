package persondb

import (
	"fiolib/internal/models"

	"gorm.io/gorm"
)

func (pd *personDB) Update(id uint, data *models.Person) (*models.Person, error) {
	obj := models.Person{Model: gorm.Model{ID: id}}

	if err := pd.db.Model(&obj).Updates(data).Error; err != nil {
		return nil, err
	}

	if err := pd.db.First(&obj).Error; err != nil {
		return nil, err
	}

	return &obj, nil
}
