package bootstrap

import (
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/lichmaker/boring-lottery/routes"
)

func SetupRoute(viewsFS embed.FS) *gin.Engine {
	router := gin.Default()
	routes.MyRender(router, viewsFS)
	routes.MyApi(router)
	routes.MyWeb(router)
	return router
}
