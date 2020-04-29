package auth

import (
	"crypto/subtle"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/breathbath/go_utils/utils/env"
	"github.com/gin-gonic/gin"
	"time"
)

var identityKeyFront = "id"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
}

/**
@apiDefine JsonHeader
@apiHeader {String} Content-Type="application/json" Json content type
@apiHeaderExample {String} Content-Type
   Content-Type:"application/json"
*/

/**
@apiDefine AuthFrontHeader
@apiHeader {String} Authorization JWT token value (use /docs/#api-Auth-Login_front to get JWT)
@apiHeaderExample {String} Authorization Header
   Authorization: "Bearer eyJhbGciOi.JSUzUxMiIsIn.R5cCI6IkpXVCJ9"
@apiErrorExample Unauthorized(401)
HTTP/1.1 401 Unauthorized
{
    "code": 401,
    "message": "cookie token is empty"
}
*/

/**
@api {post} /front-login Login front
@apiDescription Login for user against front API
@apiName Login front
@apiGroup Auth
@apiUse JsonHeader
@apiParamExample {json} Body:
{
	"username": "admin",
	"password": "admin"
}

@apiSuccessExample Success-Response
HTTP/1.1 200 OK
{
    "code": 200,
    "expire": "2020-04-20T13:30:49+03:00",
    "token": "YOUR_TOKEN_RESULT"
}

@apiErrorExample Bad request(401)
HTTP/1.1 401 Unauthorized
{
    "code": 401,
    "message": "incorrect Username or Password"
}
*/

/**
@api {post} /front-refresh Refresh token front
@apiDescription Refreshes the auth token for the front API
@apiName Refresh token front
@apiGroup Auth

@apiUse JsonHeader
@apiUse AuthFrontHeader

@apiSuccessExample Success-Response
HTTP/1.1 200 OK
{
    "code": 200,
    "expire": "2020-04-20T13:47:59+03:00",
    "token": "YOUR_TOKEN_RESULT"
}

@apiErrorExample Bad request(401)
HTTP/1.1 401 Unauthorized
{
    "code": 401,
    "message": "cookie token is empty"
}
*/
//BuildFrontMiddleWare builds middleware for auth
func BuildFrontMiddleWare() (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       env.ReadEnv("AUTH_REALM", "testing"),
		Key:         []byte(env.ReadEnv("AUTH_SECRET", "123456")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKeyFront,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKeyFront: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKeyFront].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			expectedUser := env.ReadEnv("AUTH_USER", "root")
			if subtle.ConstantTimeCompare([]byte(userID), []byte(expectedUser)) == 0 {
				return nil, jwt.ErrFailedAuthentication
			}

			exprectedPass := env.ReadEnv("AUTH_PASS", "root")
			if subtle.ConstantTimeCompare([]byte(password), []byte(exprectedPass)) == 0 {
				return nil, jwt.ErrFailedAuthentication
			}
			return &User{
				UserName:  userID,
			}, nil

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
