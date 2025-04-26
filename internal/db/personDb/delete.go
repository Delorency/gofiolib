package persondb

import (
	"fiolib/internal/models"

	"gorm.io/gorm"
)

func (pd *personDB) Delete(id uint) error {
	obj := models.Person{Model: gorm.Model{ID: id}}

	if err := pd.db.First(&obj).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return pd.db.Delete(&obj).Error
}
