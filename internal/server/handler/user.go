package handler

import (
	"blog_api/er"
	"blog_api/internal/server/mw/jwt"
	"blog_api/pkg/cache"
	"blog_api/pkg/user"
	"context"
	"fmt"
	"net/http"
	"time"

	model "blog_api/utils/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	log           *logrus.Logger
	jwtMiddleware *jwt.GinJWTMiddleware
	userService   *user.Service
	cacheService  *cache.Service
}

func newUserHandler(
	log *logrus.Logger,
	userService *user.Service,
	cacheService *cache.Service,
) *UserHandler {
	c := &gin.Context{}
	return &UserHandler{
		log,
		jwt.SetAuthMiddleware(userService.Repo.GetDBConnection(c)),
		userService,
		cacheService,
	}
}

func (h *UserHandler) UserRegistration(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &user.User{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	err = h.userService.UpsertUserRegistration(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Registration Sucessfully Done"
	res.Success = true
	res.Data = req
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) UserLogin(c *gin.Context) {
	var (
		err  error
		dCtx = context.Background()
		req  = user.User{}
		res  = model.GenericRes{}
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", req).Warn(err.Error())
		}
	}()

	//check if location pings are present for today
	//if not, then verify that rider should be at store before proceeding
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}

	user, ok, err := h.userService.FetchUserByMobileNumberOrEmail(dCtx, req)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	if ok {
		code, _, _ := h.jwtMiddleware.SetToken(c, user)
		if code == 0 {
			err = fmt.Errorf("jwt set token failed")
			return
		}
		err := h.cacheService.Repo.Set(fmt.Sprint(user.ID), nil, 10*time.Hour)
		if err != nil {
			err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
			return
		}
		res.Success = true
		c.JSON(http.StatusOK, res)
		return
	}

}
