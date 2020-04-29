package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/breathbath/go_utils/utils/env"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var IdentityKeyBack = "erply_id"

type Session struct {
	SessionID string `json:"session_id" binding:"required"`
	ErplyID   string `json:"erply_id" binding:"required"`
}

type ErplyRemoteAuth struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	ErplyID  string `json:"clientCode" binding:"required"`
}

/**
@apiDefine AuthBackHeader
@apiHeader {String} Authorization JWT token value (use /docs/#api-Auth-Login_backend to get JWT)
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
@api {post} /back-login Login backend
@apiDescription Login for user against backend API
@apiName Login backend
@apiGroup Auth
@apiUse JsonHeader
@apiParamExample {json} Body:
{
	"session_id": "Dfsajfjflkdjfldsjflsdfja",
	"erply_id": "506722"
}

@apiParamExample {json} Body2:
{
	"username": "no@mail.me",
	"password": "Dfsajfjflkdjfldsjflsdfja",
	"clientCode": "506722"
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
@api {post} /back-refresh Refresh token back
@apiDescription Refreshes the auth token for the backend API
@apiName Refresh token back
@apiGroup Auth

@apiUse JsonHeader
@apiUse AuthBackHeader

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
func BuildBackMiddleWare() (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       env.ReadEnv("AUTH_REALM", "testing"),
		Key:         []byte(env.ReadEnv("AUTH_SECRET_BACK", "12345698")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKeyBack,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(Session); ok {
				return jwt.MapClaims{
					IdentityKeyBack: v.ErplyID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return Session{
				ErplyID: claims[IdentityKeyBack].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var authData ErplyRemoteAuth

			cl := &http.Client{Timeout: 10 * time.Second}

			if err := c.ShouldBind(&authData); err == nil {
				erplySess, err := api.VerifyUser(authData.UserName, authData.Password, authData.ErplyID, cl)
				if err != nil {
					log.Errorf("Failed verify session with the remote API: %v", err)
					return nil, jwt.ErrFailedAuthentication
				}

				return Session{SessionID: erplySess, ErplyID: authData.ErplyID}, nil
			}

			var sess Session
			if err := c.ShouldBind(&sess); err != nil {
				log.Errorf("Invalid session data in input: %v", err)
				return "", jwt.ErrMissingLoginValues
			}

			sessUser, err := api.GetSessionKeyUser(sess.SessionID, sess.ErplyID, cl)
			if err != nil {
				log.Errorf("Failed verify session with the remote API: %v", err)
				return nil, jwt.ErrFailedAuthentication
			}

			if sessUser.SessionKey != sess.SessionID {
				log.Errorf("Unexpected session key, expected key is %v", sessUser.SessionKey)
				return nil, jwt.ErrFailedAuthentication
			}

			return sess, nil

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(Session); ok {
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
