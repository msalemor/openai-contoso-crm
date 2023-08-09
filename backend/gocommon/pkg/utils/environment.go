package utils

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/msalemor/contoso-crm-common/pkg/models"
)

func loadEnvironment(app *models.CrmApplication) {
	godotenv.Load()
	shost := os.Getenv("APP_HOST")
	if shost != "" {
		app.App_Host = shost
	} else {
		app.App_Host = "localhost"
	}
	sport := os.Getenv("APP_PORT")
	if sport != "" {
		app.App_Port = sport
	} else {
		app.App_Port = "8080"
	}
	scondb := os.Getenv("DB_CONNECTION")
	if scondb != "" {
		app.DB_Connection = scondb
	} else {
		app.DB_Connection = "./contosocrm.db"
	}
	sGinMode := os.Getenv("GIN_MODE")
	if sGinMode != "" {
		app.GinMode = sGinMode
	} else {
		app.GinMode = "Debug"
	}
	if app.GinMode == "Release" || app.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	app.ContactSvcURI = os.Getenv("CONTACT_SVC_URI")
	app.LeadSvcURI = os.Getenv("LEAD_SVC_URI")
	app.CompanySvcURI = os.Getenv("COMPANY_SVC_URI")
	app.OpportunitySvcURI = os.Getenv("OPPORTUNITY_SVC_URI")
	app.QuerySvcURI = os.Getenv("QUERY_SVC_URI")
	app.ProcessorSvcURI = os.Getenv("PROCESSOR_SVC_URI")

	// TODO: Remove this when the service bus client is implemented
	app.SbConnectionString = os.Getenv("SB_CONNECTION_STRING")
}
