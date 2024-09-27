package controller

import "github.com/gin-gonic/gin"

var controllers map[string][]*AergiaController

type AergiaController struct {
	group    *string
	path     *string
	method   *string
	function func(*gin.Context)
}

func NewAergiaController(group, path, method string, function func(*gin.Context)) {
	if controllers == nil {
		controllers = make(map[string][]*AergiaController)
	}
	if _, ok := controllers[group]; !ok {
		controllers[group] = []*AergiaController{
			{
				function: function,
				group:    &group,
				method:   &method,
				path:     &path,
			},
		}
		return
	}
	controllers[group] = append(controllers[group], &AergiaController{
		group:    &group,
		method:   &method,
		function: function,
		path:     &path,
	})
}

func AddController(engine *gin.Engine) {
	for key, controller := range controllers {
		routerGroup := engine.Group(key)
		for _, controllerGroup := range controller {
			routerGroup.Handle(*controllerGroup.method, *controllerGroup.path, controllerGroup.function)
		}
	}
}

func StartControllers() {
	AuthControllers()
}
