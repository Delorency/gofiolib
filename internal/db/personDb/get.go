package persondb

import (
	"fiolib/internal/models"
	schemes "fiolib/internal/schemes"
	"time"

	"gorm.io/gorm"
)

func (pd *personDB) List(p *schemes.Pagination, f *schemes.Filter) (*[]models.Person, error) {
	var persons []models.Person
	query := pd.db.Model(&models.Person{})

	if f.ID > 0 {
		query = query.Where("id = ?", f.ID)
	}
	if f.Name != "" {
		query = query.Where("name LIKE ?", "%"+f.Name+"%")
	}
	if f.Surname != "" {
		query = query.Where("surname LIKE ?", "%"+f.Surname+"%")
	}
	if f.Patronymic != "" {
		query = query.Where("patronymic LIKE ?", "%"+f.Patronymic+"%")
	}
	if f.Age > 0 {
		query = query.Where("age >= ?", f.Age)
	}
	if f.Gender != "" {
		query = query.Where("gender = ?", f.Gender)
	}
	if f.Nat != "" {
		query = query.Where("nat = ?", f.Nat)
	}
	if f.CreatedAt != "" {
		createdAt, err := time.Parse("2006-01-02", f.CreatedAt)
		if err == nil {
			query = query.Where("created_at >= ?", createdAt)
		}
	}
	if f.UpdatedAt != "" {
		updatedAt, err := time.Parse("2006-01-02", f.UpdatedAt)
		if err == nil {
			query = query.Where("updated_at >= ?", updatedAt)
		}
	}
	if f.DeletedAt != "" {
		deletedAt, err := time.Parse("2006-01-02", f.DeletedAt)
		if err == nil {
			query = query.Where("deleted_at >= ?", deletedAt)
		}
	}

	err := query.Where(models.Person{}).Limit(p.Limit).Offset((p.Page - 1) * p.Limit).Order("created_at desc").Find(&persons).Error

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
