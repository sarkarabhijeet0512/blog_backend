package jwt

import (
	"blog_api/pkg/cache"
	"blog_api/pkg/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
)

var (
	identityKey = "id"
)

// type (
// 	login struct {
// 		Mobile   int    `form:"mobile" json:"mobile" binding:"required"`
// 		Password string `form:"password" json:"password" binding:"required"`
// 	}
// 	signup struct {
// 		Mobile   int    `form:"mobile" json:"mobile" binding:"required"`
// 		Password string `form:"password" json:"password" binding:"required"`
// 	}
// )

type JWTHandler struct {
	redisService *cache.Service
}

func newJWTHandler(
	redisService *cache.Service,
) *JWTHandler {
	return &JWTHandler{
		redisService: redisService,
	}
}

// the jwt middleware
func SetAuthMiddleware(db *pg.DB) *GinJWTMiddleware {
	// t := time.Now().Add(time.Hour * 24 * 30)
	// // t = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, t.Nanosecond(), t.Location())
	// logoutTimeout := time.Since(t)
	// if logoutTimeout.Seconds() < 0 {
	// 	logoutTimeout = logoutTimeout * -1
	// }
	authMiddleware, err := New(&GinJWTMiddleware{
		Realm:       "ecommerce backend",
		Key:         []byte("ecommerceusersecretnew1"),
		Timeout:     time.Hour * 24 * 30,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		SendCookie:  true,
		PayloadFunc: func(data interface{}) MapClaims {
			if v, ok := data.(*user.User); ok {
				return MapClaims{
					identityKey: v.ID,
				}
			}
			return MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := ExtractClaims(c)
			return claims["id"]
		},
		// Authenticator: func(c *gin.Context) (interface{}, error) {
		// 	var loginVals user.User
		// 	if err := c.ShouldBind(&loginVals); err != nil {
		// 		return "", ErrMissingLoginValues
		// 	}
		// 	userID := loginVals.ID
		// 	password := loginVals.Password

		// 	// Example validation, replace with your actual authentication logic
		// 	if userID == 1 && password == "admin" {
		// 		return &user.User{
		// 			ID: userID,
		// 		}, nil
		// 	}
		// 	return nil, ErrFailedAuthentication
		// },
		// add logic if admin is creating user than has option for creating admin/write access user account else normal user account to be created
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*user.User); ok {
				return true
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}
	return authMiddleware
}

func RoleAccess(db *pg.DB, role, user int) bool {

	return true
}
