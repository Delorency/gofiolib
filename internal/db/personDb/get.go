package persondb

import (
	"fiolib/internal/models"
	schemes "fiolib/internal/schemes"

	"gorm.io/gorm"
)

func (pd *personDB) List(p *schemes.Pagination) (*[]models.Person, error) {
	var persons []models.Person

	err := pd.db.Where(models.Person{}).Limit(p.Limit).Offset((p.Page - 1) * p.Limit).Order("created_at desc").Find(&persons).Error

	if err != nil {
		return nil, err
	}

	return &persons, err
}
func (pd *personDB) Retrieve(id uint) (*models.Person, error) {
	obj := models.Person{Model: gorm.Model{ID: id}}

	err := pd.db.First(&obj).Error

	if err != nil {
		return nil, err
	}

	return &obj, err
}
