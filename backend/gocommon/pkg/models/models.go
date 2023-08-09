package models

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CrmApplication struct {
	Db                 *gorm.DB
	Router             *gin.Engine
	SbConnectionString string
	SbClient           *azservicebus.Client
	GinMode            string
	App_Host           string
	App_Port           string
	DB_Connection      string
	ContactSvcURI      string
	LeadSvcURI         string
	CompanySvcURI      string
	OpportunitySvcURI  string
	QuerySvcURI        string
	ProcessorSvcURI    string
}

type EventMessage struct {
	Action  string `json:"action"`
	Model   string `json:"model"`
	Payload any    `json:"payload"`
}

type Lead struct {
	gorm.Model
	FirstName    string `gorm:"size:50;not null"`
	LastName     string `gorm:"size:50;not null"`
	Email        string `gorm:"size:100;not null"`
	Phone        string `gorm:"size:25;not null"`
	JobTitle     string `gorm:"size:100;not null"`
	CompanyID    uint
	Address      string `gorm:"size:100;not null"`
	City         string `gorm:"size:50;not null"`
	State        string `gorm:"size:2;not null"`
	Zip          string `gorm:"size:10;not null"`
	OtherDetails string
	Test         string
}

type Contact struct {
	gorm.Model
	FirstName    string `gorm:"size:50;not null"`
	LastName     string `gorm:"size:50;not null"`
	Email        string `gorm:"size:100;not null"`
	Phone        string `gorm:"size:25;not null"`
	JobTitle     string `gorm:"size:100;not null"`
	CompanyID    uint
	Address      string `gorm:"size:100;not null"`
	City         string `gorm:"size:50;not null"`
	State        string `gorm:"size:2;not null"`
	Zip          string `gorm:"size:10;not null"`
	OtherDetails string
}

type Company struct {
	gorm.Model
	Name         string `gorm:"size:100;not null"`
	Email        string `gorm:"size:100;not null"`
	Phone        string `gorm:"size:25;not null"`
	WebSite      string
	Address      string `gorm:"size:100;not null"`
	City         string `gorm:"size:50;not null"`
	State        string `gorm:"size:2;not null"`
	Zip          string `gorm:"size:10;not null"`
	OtherDetails string
}

type Opportunity struct {
	gorm.Model
	CustomerID      uint
	ProductService  string  `gorm:"size:100;not null"`
	ExpectedRevenue float64 `gorm:"not null,default:0"`
	DealStage       string  `gorm:"not null,default:'Pending'"`
	Probability     float64 `gorm:"not null,default:0"`
	OtherDetails    string
}
