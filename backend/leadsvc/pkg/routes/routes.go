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
	app.Router.GET("/leads", getAll)
	app.Router.GET("/leads/:id", getByID)
	app.Router.POST("/leads", create)
	app.Router.PUT("/leads/:id", update)
	app.Router.PATCH("/leads/:id", patch)
	app.Router.DELETE("/leads/:id", delete)
}
