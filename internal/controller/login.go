package controller

// import (
// 	"hstcmler/errcode"

// 	"wuliangxiaoseng/openapi/internal/common"
// 	"wuliangxiaoseng/openapi/internal/consts"
// 	"wuliangxiaoseng/openapi/internal/logger"
// 	"wuliangxiaoseng/openapi/internal/middleware/my_jwt"
// 	"wuliangxiaoseng/openapi/internal/response"
// 	"wuliangxiaoseng/openapi/internal/routers"
// 	system_service "wuliangxiaoseng/openapi/internal/service/system"
// 	user_manage "wuliangxiaoseng/openapi/internal/service/user"

// 	"github.com/gin-gonic/gin"
// )

// type Profile struct {
// 	Username string `db:"username"`
// 	Password string `db:"password"`
// }

// var IP, LoginUsername string

// func login(context *gin.Context) {
// 	var user Profile
// 	err := context.ShouldBindJSON(&user)
// 	if err != nil {
// 		response.Fail(context, errcode.RESULT_ERROR_INVALID_PARAM, consts.ValidatorParamsCheckFailMsg, "")
// 		return
// 	}
// 	if v, e := user_manage.RetrieveAntiReplayInterval(); e == nil {
// 		my_jwt.Anti_replay_interval = v
// 	}

// 	if userMsg, rtncode := user_manage.UserLogin(user.Username, user.Password); rtncode == errcode.RESULT_SUCCESS {
// 		if userToken, ok := my_jwt.GenerateToken(user.Username, my_jwt.ExpireAt); ok {
// 			if my_jwt.RecordLoginToken(userToken) {
// 				response.ReturnTokenJson(context, userToken, errcode.RESULT_SUCCESS, "", userMsg)
// 				logger.Oplog("The web user %s login successful.", user.Username)
// 				return
// 			}
// 		} else {
// 			logger.Oplog("The web user %s generate token fail.", user.Username)
// 			response.Fail(context, errcode.RESULT_ERROR_COMMON, consts.ServerOccurredErrorMsg, "")
// 		}
// 	} else if rtncode == errcode.RESULT_ERROR_INVALID_USERNAME {
// 		logger.Oplog("The web user %s is invalid.", user.Username)
// 		response.Fail(context, rtncode, consts.CurdLoginUserErr, "")
// 	} else if rtncode == errcode.RESULT_ERROR_INVALID_PASSWORD {
// 		logger.Oplog("The web user %s password is invalid.", user.Username)
// 		response.Fail(context, rtncode, consts.CurdLoginPasswdErr, "")
// 	} else if rtncode == errcode.RESULT_ERROR_USERNAME_LOCK {
// 		logger.Oplog("The web user %s is locked.", user.Username)
// 		response.Fail(context, rtncode, consts.CurdLoginUserLock, "")
// 	} else {

// 	}
// }
// func logout(context *gin.Context) {
// 	response.Success(context, consts.LogoutOk, "")

// 	logger.Oplog("The web user %s logout successful.", my_jwt.GetUserMsg(context))
// }

// func loginV2(context *gin.Context) {
// 	//告警需要的数据需得当场获取
// 	if v, e := user_manage.RetrieveAntiReplayInterval(); e == nil {
// 		my_jwt.Anti_replay_interval = v
// 	}
// 	IP = context.ClientIP()
// 	if IP == "" {

// 		IP = context.Request.Header.Get("X-Forward-For")
// 		common.Printf("X-Forward-For IP:%s   ------------", IP)

// 	}
// 	var position = IP + ":" + "web"
// 	//LoginUsername = my_jwt.GetUserMsg(context)  就只有登陆的时候用不了这个获取，因为这个阶段没有带token
// 	//告警需要的数据需得当场获取
// 	var user Profile
// 	err := context.ShouldBindJSON(&user)
// 	if err != nil {
// 		response.Fail(context, errcode.RESULT_ERROR_INVALID_PARAM, consts.ValidatorParamsCheckFailMsg, "")
// 		return
// 	}
// 	if userMsg, rtncode := user_manage.UserLogin(user.Username, user.Password); rtncode == errcode.RESULT_SUCCESS {
// 		if userToken, ok := my_jwt.GenerateToken(user.Username, my_jwt.ExpireAt); ok {
// 			if my_jwt.RecordLoginToken(userToken) {
// 				response.ReturnTokenJson(context, userToken, errcode.RESULT_SUCCESS, "", userMsg)
// 				logger.Oplog("The web user %s login successful.", user.Username)
// 				logger.AlarmfV2(position, user.Username, logger.WEB_ALARM_LOGIN_SUCCESS, logger.ALARM_LEVEL_NOTICE, "The web user %s login successful!", user.Username)
// 				return
// 			}
// 		} else {
// 			logger.Oplog("The web user %s generate token fail.", user.Username)
// 			response.Fail(context, errcode.RESULT_ERROR_COMMON, consts.ServerOccurredErrorMsg, "")
// 		}
// 	} else if rtncode == errcode.RESULT_ERROR_INVALID_USERNAME {
// 		logger.Oplog("The web user %s is invalid.", user.Username)
// 		logger.AlarmfV2(position, user.Username, logger.WEB_ALARM_LOGIN_FAIL, logger.ALARM_LEVEL_NOTICE, "The web user %s is invalid.", user.Username)
// 		response.Fail(context, rtncode, consts.CurdLoginUserOrPasswdErr, "")
// 	} else if rtncode == errcode.RESULT_ERROR_INVALID_PASSWORD {
// 		logger.Oplog("The web user %s password is invalid.", user.Username)
// 		logger.AlarmfV2(position, user.Username, logger.WEB_ALARM_LOGIN_FAIL, logger.ALARM_LEVEL_NOTICE, "The web user %s password is invalid.", user.Username)
// 		response.Fail(context, rtncode, consts.CurdLoginUserOrPasswdErr, "")
// 	} else if rtncode == errcode.RESULT_ERROR_USERNAME_LOCK {
// 		logger.Oplog("The web user %s is locked.", user.Username)
// 		logger.AlarmfV2(position, user.Username, logger.WEB_ALARM_LOGIN_FAIL, logger.ALARM_LEVEL_NOTICE, "The web user %s is locked.", user.Username)
// 		response.Fail(context, rtncode, consts.CurdLoginUserLock, "")
// 	} else {

// 	}
// }
// func logoutV2(context *gin.Context) {
// 	//告警需要的数据需得当场获取
// 	IP = context.ClientIP()
// 	if IP == "" {

// 		IP = context.Request.Header.Get("X-Forward-For")
// 		common.Printf("X-Forward-For IP:%s   ------------", IP)

// 	}
// 	var position = IP + ":" + "web"
// 	LoginUsername = my_jwt.GetUserMsg(context)
// 	//告警需要的数据需得当场获取
// 	response.Success(context, consts.LogoutOk, "")
// 	logger.Oplog("The web user %s logout successful.", my_jwt.GetUserMsg(context))
// 	logger.AlarmfV2(position, LoginUsername, logger.WEB_ALARM_LOGOUT_SUCCESS, logger.ALARM_LEVEL_NOTICE, "The web user %s logout successful!", my_jwt.GetUserMsg(context))
// }

// // 用户登录
// // func userLogin(userName string, pass string) (*rtnMsg, bool) {
// // 	//需要用户验证的时候用
// // 	if userName == "anos" && pass == "anos" {
// // 		userMsg := &rtnMsg{
// // 			AccessType: 1,
// // 			Username:   userName,
// // 			AuthType:   4,
// // 			RoleID:     1,
// // 		}
// // 		return userMsg, true
// // 	}
// // 	return nil, false
// // }
// func getProjectType(c *gin.Context) {
// 	//response.NotImpl(c)

// 	prjtype, err := system_service.PrjTypeGet()
// 	if err != nil {
// 		response.Fail(c, errcode.RESULT_ERROR_COMMON, err.Error(), nil)
// 		return
// 	}
// 	if userToken, ok := my_jwt.GenerateToken("administrator", my_jwt.ExpireAt); ok {
// 		if my_jwt.RecordLoginToken(userToken) {
// 			response.ReturnTokenJson(c, userToken, errcode.RESULT_SUCCESS, "", prjtype)
// 			//logger.Oplog("The web user %s login successful.", user.Username)
// 			return
// 		}
// 	} else {
// 		logger.Oplog("The generate token fail.")
// 		response.Fail(c, errcode.RESULT_ERROR_COMMON, consts.ServerOccurredErrorMsg, "")
// 	}
// 	//response.Success(c, consts.QueryCfgOk, prjtype)
// }
// func loginRoute(e *gin.Engine, apiGroup *gin.RouterGroup) {
// 	publicGroup := e.Group("public")
// 	publicGroup.Use(my_jwt.TimeMiddleware())

// 	publicGroup.PUT("/login", login)
// 	publicGroup.PUT("/logout", logout)

// 	apiGroup.GET("/retoken", my_jwt.RefreshToken)

// 	publicGroup.Group("v2").
// 		PUT("/login", loginV2).
// 		PUT("/logout", logoutV2).
// 		GET("/prjtype", getProjectType) //获取项目类型，选择对应系统的图标

// 	//apiGroup.GET("/retoken", my_jwt.RefreshToken)
// }

// func init() {
// 	routers.Register("login", loginRoute)
// }
