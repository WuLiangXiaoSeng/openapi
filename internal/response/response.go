package response

import (
	"net/http"
	"wuliangxiaoseng/errcode"

	"wuliangxiaoseng/openapi/internal/consts"
	"wuliangxiaoseng/openapi/internal/inter_errors"

	"github.com/gin-gonic/gin"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode errcode.RESULT, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code":    dataCode,
		"message": msg,
		"value":   data,
	})
}

func ReturnTokenJson(Context *gin.Context, token string, dataCode errcode.RESULT, msg string, data interface{}) {

	Context.Header("authorization", token) //可以根据实际情况在头部添加额外的其他信息
	Context.JSON(http.StatusOK, gin.H{
		"code":    dataCode,
		"message": msg,
		"value":   data,
	})
}

//ReturnJsonFromString 将json字符串以标准json格式返回（例如，从redis读取json格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

//Success 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, errcode.RESULT_SUCCESS, msg, data)
}

//Fail 失败的业务逻辑
func Fail(c *gin.Context, dataCode errcode.RESULT, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

// ErrorTokenBaseInfo token 基本的格式错误
func ErrorTokenBaseInfo(c *gin.Context) {
	ReturnJson(c, http.StatusBadRequest, http.StatusBadRequest, inter_errors.ErrorsTokenBaseInfo, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

//ErrorTokenAuthFail token 权限校验失败
func ErrorTokenAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, inter_errors.ErrorsNoAuthorization, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

//ErrorTokenRefreshFail token不符合刷新条件
func ErrorTokenRefreshFail(c *gin.Context) {
	ReturnJson(c, http.StatusBadRequest, http.StatusBadRequest, inter_errors.ErrorsRefreshTokenFail, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

// ErrorCasbinAuthFail 鉴权失败，返回 405 方法不允许访问
func ErrorCasbinAuthFail(c *gin.Context, msg interface{}) {
	ReturnJson(c, http.StatusMethodNotAllowed, http.StatusMethodNotAllowed, inter_errors.ErrorsCasbinNoAuthorization, msg)
	c.Abort()
}

//ErrorParam 参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, errcode.RESULT_ERROR_INVALID_PARAM, errcode.RESULT_name[int32(errcode.RESULT_ERROR_INVALID_PARAM)], wrongParam)
	c.Abort()
}

// ErrorSystem 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, errcode.RESULT_ERROR_SYSMGR_NOT_SUPPORT, consts.ServerOccurredErrorMsg+msg, data)
	c.Abort()
}

//NotImpl 尚未实现
func NotImpl(c *gin.Context) {
	ReturnJson(c, http.StatusNotImplemented, errcode.RESULT_ERROR_NOT_IMPLEMENTED, errcode.RESULT_name[int32(errcode.RESULT_ERROR_NOT_IMPLEMENTED)], "")
	c.Abort()
}
