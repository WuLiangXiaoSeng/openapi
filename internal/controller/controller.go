package controller

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var controller Controller

type Controller struct {
	engine  *gin.Engine
	modules sync.Map
}

func Init(engine *gin.Engine) bool {
	controller.engine = engine
	controller.modules.Range(func(key, value interface{}) bool {
		return value.(func() bool)()
	})
	return true
}

func Register(module string, initFunc func() bool) {
	controller.modules.Store(module, initFunc)
}
