package main

import (
	"fmt"

	"github.com/msalemor/contoso-crm-common/pkg/models"
	"github.com/msalemor/contoso-crm-common/pkg/utils"
)

func main() {
	contact := &models.Contact{}
	fmt.Println(contact)
	lead := &models.Lead{}
	fmt.Println(lead)
	opportunity := &models.Opportunity{}
	fmt.Println(opportunity)
	company := &models.Company{}
	fmt.Println(company)
	app := utils.NewCrmApplication()
	fmt.Println(app)
	utils.RunRouter(app)
}
