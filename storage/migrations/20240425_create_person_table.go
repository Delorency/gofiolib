package migrations

import (
	"fiolib/internal/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePersonTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240425_create_person_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Person{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(models.Person{}.TableName())
		},
	}
}
