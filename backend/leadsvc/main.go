package main

import (
	"fmt"
	"leadsvc/pkg/routes"

	"github.com/msalemor/contoso-crm-common/pkg/utils"
)

func main() {
	var app = utils.NewCrmApplication()
	fmt.Println(app)
	routes.MapRoutes(app)
	utils.RunRouter(app)
}
