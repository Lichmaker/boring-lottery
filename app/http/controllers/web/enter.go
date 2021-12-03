package web

type WebGroup struct {
	HomeController HomeController
}

var WebGroupApp = new(WebGroup)