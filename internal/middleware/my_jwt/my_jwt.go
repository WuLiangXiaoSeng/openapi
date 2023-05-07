package my_jwt

import (
	"math"
	"strconv"
	"strings"
	"time"
	"wuliangxiaoseng/errcode"

	"wuliangxiaoseng/openapi/internal/common"
	"wuliangxiaoseng/openapi/internal/consts"
	"wuliangxiaoseng/openapi/internal/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Profile struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

var Anti_replay_interval int32 = 0

var Secret = []byte("secret") //密码自行设定
var ExpireAt int64 = 60 * 5
var TokenField string = "authorization"
var Timestamp string = "timestamp"

// 解析Token
func parseToken(tokenString string) (*MyClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		// logger.Errorf("jwt", errcode.RESULT_ERROR_COMMON, err.Error())
		return nil, false
	}
	if err != nil {
		// logger.Errorf("jwt", errcode.RESULT_ERROR_COMMON, err.Error())
		return nil, false
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}

func RecordLoginToken(userToken string) bool {
	//需要记录token的时候用
	// if _, ok := parseToken(userToken); ok {
	// 	return true
	// } else {
	// 	return false
	// }
	return true
}

func createToken(claims MyClaims) (string, bool) {
	// 生成jwt格式的header、claims 部分
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	token, _ := tokenPartA.SignedString(Secret)
	return token, true
}

func GenerateToken(username string, expireAt int64) (string, bool) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	c := MyClaims{
		username, // 自定义字段
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),            // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireAt, // 失效截止时间
			Issuer:    "eqt",                        // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token, _ := createToken(c)
	return token, true
}

func refreshToken(oldToken, clientIp string) (newToken string, res bool) {
	//预留验证clientIp的位置
	if newToken, ok := resetTokenTime(oldToken, ExpireAt); ok {
		if _, ok := parseToken(newToken); ok {
			return newToken, true
		}
	}
	return "", false
}

// 更新token
func resetTokenTime(tokenString string, extraAddSeconds int64) (string, bool) {
	if CustomClaims, ok := parseToken(tokenString); ok {
		if time.Now().Unix() > CustomClaims.ExpiresAt {
			// logger.Errorf("jwt", errcode.RESULT_ERROR_TIME_OUT, "resetTokenTime fail, please relogin\n")
			return "", false
		}
		CustomClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		return createToken(*CustomClaims)
	} else {
		return "", false
	}
}

// 刷新用户token
func RefreshToken(context *gin.Context) {
	oldToken := context.Request.Header.Get(TokenField)
	if oldToken == "" {
		response.Fail(context, errcode.RESULT_ERROR_INVALID_CERT, consts.JwtTokenMustValid, "")
		context.Abort()
		return
	}

	// 按空格分割
	parts := strings.SplitN(oldToken, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.ErrorTokenAuthFail(context)
		return
	}

	if newToken, ok := refreshToken(parts[1], context.ClientIP()); ok {
		response.ReturnTokenJson(context, newToken, errcode.RESULT_SUCCESS, "", "")
	} else {
		response.Fail(context, errcode.RESULT_ERROR_COMMON, consts.CurdRefreshTokenFailMsg, "")
	}
}

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(TokenField)
		if authHeader == "" {
			response.Fail(c, errcode.RESULT_ERROR_INVALID_CERT, consts.JwtTokenMustValid, "")
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.ErrorTokenAuthFail(c)
			return
		}

		// 按.分割
		part := strings.Split(parts[1], ".")
		if len(part) != 3 {
			response.ErrorTokenAuthFail(c)
			return
		}
		_, ok := parseToken(parts[1])
		if !ok {
			response.ErrorTokenAuthFail(c)
			return
		}
		//stamp := config.ServerConfig.TimeStamp
		stamp := Anti_replay_interval
		common.Println("stamp:", stamp)
		if stamp != 0 {
			timeStamp := c.Query("_t")
			common.Println(timeStamp)
			localTime := time.Now().Unix() * 1000
			remoteTime, _ := strconv.Atoi(timeStamp)

			if math.Abs(float64(localTime-int64(remoteTime))) > float64(stamp*1000) {
				common.Println("ERR: ", "localTime:", localTime, "remoteTime:", remoteTime)
				response.Fail(c, errcode.RESULT_ERROR_INVALID_CERT, consts.CurdOperateFailMsg, localTime)
				return
			} else {
				common.Println("OK: ", "localTime:", localTime, "remoteTime:", remoteTime)
			}
		}
		c.Next()
	}
}

func TimeMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		timeStamp := c.Query("_t")
		common.Println(timeStamp)
		localTime := time.Now().Unix() * 1000
		remoteTime, _ := strconv.Atoi(timeStamp)

		stamp := Anti_replay_interval

		common.Println("stamp:", stamp)
		if stamp != 0 {
			if math.Abs(float64(localTime-int64(remoteTime))) > float64(stamp*1000) {
				common.Println("ERR: ", "localTime:", localTime, "remoteTime:", remoteTime)
				response.Fail(c, errcode.RESULT_ERROR_INVALID_CERT, consts.CurdOperateFailMsg, localTime)
				return
			} else {
				common.Println("OK: ", "localTime:", localTime, "remoteTime:", remoteTime)
			}
		}
		c.Next()
	}
}

//此接口只用于用户退出时，获取用户名使用
func GetUserMsg(c *gin.Context) string {
	authHeader := c.Request.Header.Get(TokenField)
	parts := strings.SplitN(authHeader, " ", 2)
	msg, _ := parseToken(parts[1])
	return msg.Username
}
