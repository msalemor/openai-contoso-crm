package routes

import (
	"github.com/msalemor/contoso-crm-common/pkg/models"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func MapRoutes(app *models.CrmApplication) {
	db = app.Db
	app.Router.GET("/contacts", getAllContacts)
	app.Router.GET("/contacts/:id", getContactByID)
	app.Router.POST("/contacts", createContact)
	app.Router.PUT("/contacts/:id", updateContact)
	app.Router.PATCH("/contacts/:id", patchContact)
	app.Router.DELETE("/contacts/:id", deleteContact)
}
