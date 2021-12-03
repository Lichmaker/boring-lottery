package routes

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"

	v1 "github.com/lichmaker/boring-lottery/app/http/controllers/v1"
	"github.com/lichmaker/boring-lottery/app/http/controllers/web"
)

func MyRender(r *gin.Engine, viewsFS embed.FS) {
	templ := template.Must(template.New("").ParseFS(viewsFS, "views/*"))
	r.SetHTMLTemplate(templ)
}

func MyWeb(r *gin.Engine) {
	var Web = web.WebGroupApp
	r.GET("/", Web.HomeController.Home)
	r.GET("/results", Web.HomeController.Results)
}

func MyApi(r *gin.Engine) {
	var Api = v1.ApiGroupApp.ApiApiGroup
	r.GET("/v1/api/get", Api.LotteryApi.Get)
}
