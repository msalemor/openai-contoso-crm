package utils

import (
	"log"

	"github.com/msalemor/contoso-crm-common/pkg/models"
	"gorm.io/gorm"
)

func InitDB(db *gorm.DB) {
	db.AutoMigrate(&models.Contact{})
	//db.Create(&models.Contact{FirstName: "John", LastName: "Doe", Email: "jdoe@company.com"})
	log.Println("Database migrated")
}
