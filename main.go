package main

import (
	"embed"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/lichmaker/boring-lottery/bootstrap"
	"github.com/lichmaker/boring-lottery/config"
	configPkg "github.com/lichmaker/boring-lottery/pkg/config"
)

//go:embed views
var views embed.FS

func main() {
	config.Initialize()
	bootstrap.SetupDB()
	r := bootstrap.SetupRoute(views)

	port := configPkg.Get("app.port")
	endless.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
