package handler

import (
	"blog_api/internal/server/mw/jwt"
	"blog_api/pkg/cache"
	"blog_api/pkg/user"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestUserHandler_UserLogin(t *testing.T) {
	type fields struct {
		log           *logrus.Logger
		jwtMiddleware *jwt.GinJWTMiddleware
		userService   *user.Service
		cacheService  *cache.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				log:           tt.fields.log,
				jwtMiddleware: tt.fields.jwtMiddleware,
				userService:   tt.fields.userService,
				cacheService:  tt.fields.cacheService,
			}
			h.UserLogin(tt.args.c)
		})
	}
}
