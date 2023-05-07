package api

import (
	"fmt"
	"strconv"

	"wuliangxiaoseng/openapi/internal/controller"
	"wuliangxiaoseng/openapi/internal/routers"

	"crypto/tls"
	"net/http"

	"wuliangxiaoseng/openapi/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

type ServerApp struct {
	GinInst *gin.Engine
}

func NewServerApp() *ServerApp {
	app := &ServerApp{
		GinInst: gin.Default(),
	}
	// app.GinInst.Use(middleware.LoggerMiddleware())

	// logger.Init()
	controller.Init(app.GinInst)
	routers.Init(app.GinInst)

	return app
}

func (s *ServerApp) Start(isHttps bool) {
	if isHttps {
		var httpsPort int = 443
		s.GinInst.Use(TlsHandler(httpsPort))

		// 指定tls加密算法，版本
		tlsconf := &tls.Config{
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
			MaxVersion:               tls.VersionTLS13,
		}
		// 下面中列出的算法，剔除了DES，漏洞解决
		tlsconf.CipherSuites = []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			//tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			//tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			//tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			//tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		}
		server := &http.Server{
			Addr:      ":" + strconv.Itoa(httpsPort),
			Handler:   s.GinInst,
			TLSConfig: tlsconf,
		}

		err := server.ListenAndServeTLS("/pm/web/eqt.pem", "/pm/web/eqt.key")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if err := s.GinInst.Run(":" + strconv.Itoa(config.ServerConfig.ListenPort)); err != nil {
		fmt.Println("start error!! ", err)
	}
}

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
