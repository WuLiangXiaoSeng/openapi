package middleware

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"

// 	"wuliangxiaoseng/openapi/internal/common"
// 	"wuliangxiaoseng/openapi/internal/logger"
// 	"wuliangxiaoseng/openapi/internal/middleware/my_jwt"

// 	"github.com/gin-gonic/gin"
// 	//"fmt"
// )

// // Request logger
// func LoggerMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var body []byte
// 		var ok error
// 		buffer := new(bytes.Buffer)
// 		method := c.Request.Method
// 		url := c.Request.URL.String()
// 		if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
// 			if body, ok = ioutil.ReadAll(c.Request.Body); ok == nil {
// 				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
// 				_ = json.Compact(buffer, body)
// 			}
// 		}

// 		// TODO: 增加操作用户信息
// 		if method != "GET" {
// 			common.Printf("%s\n", method+" "+url+": "+buffer.String())
// 			index := strings.Split(url, "?")
// 			if syslog_filter_url(index[0]) {
// 			} else {
// 				position := my_jwt.GetUserMsg(c) + "@" + c.ClientIP() + "@" + "web"
// 				common.Printf("position = %s\n", position)
// 				if buffer.String() == "" {
// 					//logger.Oplog("%s", method+" "+url)
// 					logger.OplogV2(position, "%s", method+" "+url)
// 				} else {
// 					//logger.Oplog("%s", method+" "+url+" "+buffer.String())
// 					logger.OplogV2(position, "%s", method+" "+url+" "+buffer.String())
// 				}

// 			}
// 		}
// 		c.Next()
// 	}
// }

// var filterUrl = []string{
// 	"/api/v2/system/user",
// 	"/api/v2/system/user/check",
// 	"/api/system/user",
// 	"/public/v2/login",
// 	"/public/login",
// 	"/api/v2/internet/snmp/user",
// }

// func syslog_filter_url(url string) bool {
// 	for _, v := range filterUrl {
// 		if url == v {
// 			return true
// 		}
// 	}
// 	return false
// }
