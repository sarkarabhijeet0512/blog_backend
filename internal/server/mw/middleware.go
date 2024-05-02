// Package mw is user Middleware package
package mw

import (
	"errors"
	"fmt"
	"net/http"

	"blog_api/er"
	"blog_api/pkg/cache"
	model "blog_api/utils/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandlerX(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := c.Errors.Last()
			if err == nil {
				// no errors, abort with success
				return
			}

			e := er.From(err.Err)

			// if !e.NOP {
			// 	sentry.CaptureException(e)
			// }

			httpStatus := http.StatusInternalServerError
			if e.Status > 0 {
				httpStatus = e.Status
			}

			c.JSON(httpStatus, e)
		}()

		c.Next()
	}
}

func RoleCheckMiddleware(rdb *cache.Service, role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("id")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Please login uisng your secured credentials!"})
			c.Abort()
		}
		userRoles := model.UserRoles{}
		err := rdb.Get(fmt.Sprint(userID), &userRoles)
		if err != nil {
			err = er.New(err, er.Unauthorized).SetStatus(http.StatusUnauthorized)
			c.Abort()
		}
		roleFound := false
		for _, userRole := range userRoles.Resource {
			if role == userRole {
				roleFound = true
				break
			}
		}
		if !roleFound {
			er.New(errors.New("invalid access"), er.Unauthorized).SetStatus(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
