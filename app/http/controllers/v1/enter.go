package v1

import "github.com/lichmaker/boring-lottery/app/http/controllers/v1/api"

type ApiGroup struct {
	ApiApiGroup api.ApiApiGroup
}

var ApiGroupApp = new(ApiGroup)
