package main

import (
	"companysvc/pkg/routes"
	"fmt"

	"github.com/msalemor/contoso-crm-common/pkg/utils"
)

func main() {
	var app = utils.NewCrmApplication()
	fmt.Println(app)
	routes.MapRoutes(app)
	utils.RunRouter(app)
}
