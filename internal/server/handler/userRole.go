package handler

import (
	"blog_api/er"
	"blog_api/internal/server/mw/jwt"
	"blog_api/pkg/rbac"
	"blog_api/pkg/user"
	model "blog_api/utils/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserRoleHandler struct {
	log             *logrus.Logger
	jwtMiddleware   *jwt.GinJWTMiddleware
	userRoleService *rbac.Service
	userService     *user.Service
}

func newUserRoleHandler(
	log *logrus.Logger,
	userRoleService *rbac.Service,
	userService *user.Service,
) *UserRoleHandler {
	c := &gin.Context{}
	return &UserRoleHandler{
		log,
		jwt.SetAuthMiddleware(userService.Repo.GetDBConnection(c)),
		userRoleService,
		userService,
	}
}

func (h *UserRoleHandler) CreateUserRole(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &rbac.Role{}
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
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}
	go func() {
		for {
			select {
			case err := <-h.userRoleService.ErrorChannel:
				// Handle the error (e.g., log it, return it, etc.)
				h.log.Error("Error:", err)
			}
		}
	}()

	err = h.userRoleService.CreateUserRole(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusServiceUnavailable)
		return
	}
	res.Data = req
	res.Success = true
	res.Message = "Sucessfully Created"
	c.JSON(http.StatusCreated, res)
}

func (h *UserRoleHandler) RoleAssignment(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &rbac.UserRole{}
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
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}

	err = h.userRoleService.AssignRole(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}

	res.Data = req
	res.Success = true
	res.Message = "Role Assigned Sucessfully"
	c.JSON(http.StatusCreated, res)
}

func (h *UserRoleHandler) UserRoleAssignedDetails(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &rbac.UserRole{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	userID, ok := c.Get("id") //take this value from user session
	if !ok {
		err = er.New(err, er.Unauthorized).SetStatus(http.StatusUnauthorized)
		return
	}
	data, err := h.userRoleService.FetchUserRole(dCtx, userID)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	res.Data = data
	res.Message = "Success"
	res.Success = true
	c.JSON(http.StatusOK, res)
}
func (h *UserRoleHandler) RoleDetails(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &rbac.UserRole{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		h.log.WithField("span", roleID).Info("error while converting string to int: " + err.Error())
		return
	}
	data, err := h.userRoleService.FetchRole(dCtx, roleID)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	res.Data = data
	res.Message = "Success"
	res.Success = true
	c.JSON(http.StatusOK, res)
}
func (h *UserRoleHandler) RoleList(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &rbac.UserRole{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	data, err := h.userRoleService.FetchAllRoles(dCtx)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	res.Data = data
	res.Message = "Success"
	res.Success = true
	c.JSON(http.StatusOK, res)
}
