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
	app.Router.GET("/companies", getAll)
	app.Router.GET("/companies/:id", getByID)
	app.Router.POST("/companies", createContact)
	app.Router.PUT("/companies/:id", update)
	app.Router.PATCH("/companies/:id", patch)
	app.Router.DELETE("/companies/:id", delete)
}
