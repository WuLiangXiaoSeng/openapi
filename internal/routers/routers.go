package routers

import (
	"net/http"
	"sync"

	"wuliangxiaoseng/openapi/internal/middleware/my_jwt"

	"github.com/gin-gonic/gin"
)

var modules sync.Map
var staticSrcPath string = "/pm/web"

func Init(e *gin.Engine) {
	e.NoRoute(gin.WrapH(http.FileServer(gin.Dir(staticSrcPath, false))))
	apiGroup := e.Group("api")
	apiGroup.Use(my_jwt.JWTAuthMiddleware())
	e.MaxMultipartMemory = 8 << 20 // 8 MiB
	modules.Range(func(key, value interface{}) bool {
		value.(func(*gin.Engine, *gin.RouterGroup))(e, apiGroup)
		return true
	})
}

func Register(module string, regFunc func(engine *gin.Engine, apiGroup *gin.RouterGroup)) {
	modules.Store(module, regFunc)
}
