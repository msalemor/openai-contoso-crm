package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/msalemor/contoso-crm-common/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func defaultGinConfig(app *models.CrmApplication) {
	app.Router = gin.Default()
	r := app.Router
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}))
}

func RunRouter(app *models.CrmApplication) {
	if app == nil {
		log.Fatal("The gin router is not initialized. Cannot run the server.")
	}
	err := app.Router.Run(app.App_Host + ":" + app.App_Port)
	if err != nil {
		log.Fatal("Error running router: ")
	}
}

func defaultDbConfig(app *models.CrmApplication) {
	db, err := gorm.Open(sqlite.Open(app.DB_Connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database", db)
	app.Db = db
}

func initTables(app *models.CrmApplication) {
	db := app.Db
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Lead{})
	db.AutoMigrate(&models.Opportunity{})
	db.AutoMigrate(&models.Company{})
}

func NewCrmApplication() *models.CrmApplication {
	app := &models.CrmApplication{}
	loadEnvironment(app)
	defaultGinConfig(app)
	defaultDbConfig(app)
	initTables(app)
	return app
}
