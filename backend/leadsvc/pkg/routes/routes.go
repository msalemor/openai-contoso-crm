package routes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/msalemor/contoso-crm-common/pkg/models"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	sbClient *azservicebus.Client
)

func MapRoutes(app *models.CrmApplication) {
	db = app.Db
	sbClient = app.SbClient
	app.Router.GET("/leads", getAll)
	app.Router.GET("/leads/:id", getByID)
	app.Router.POST("/leads", create)
	app.Router.PUT("/leads/:id", update)
	app.Router.PATCH("/leads/:id", patch)
	app.Router.DELETE("/leads/:id", delete)
}
