// Package mw is user Middleware package
package mw

import (
	"errors"
	"fmt"
	"net/http"

	"blog_api/er"
	"blog_api/pkg/cache"
	"blog_api/pkg/rbac"

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

func RoleCheckMiddleware(rdb *cache.Service, roleID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("id")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Please login uisng your secured credentials!"})
			c.Abort()
		}
		userRoles := []rbac.UserRole{}
		err := rdb.Get(fmt.Sprint(userID), &userRoles)
		if err != nil {
			err = er.New(err, er.Unauthorized).SetStatus(http.StatusUnauthorized)
			c.Abort()
		}

		for _, userRole := range userRoles {
			for _, role := range userRole.Roles {
				for _, permission := range role.RolePermission {
					if roleID == permission.ResourceID {
						c.Next()
						return
					}
				}
			}
		}
		er.New(errors.New("invalid access"), er.Unauthorized).SetStatus(http.StatusUnauthorized)
		c.Abort()
		return
	}
}
