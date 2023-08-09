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
	app.Router.GET("/contacts", getAll)
	app.Router.GET("/contacts/:id", getByID)
	app.Router.POST("/contacts", create)
	app.Router.PUT("/contacts/:id", update)
	app.Router.PATCH("/contacts/:id", patch)
	app.Router.DELETE("/contacts/:id", delete)
}
