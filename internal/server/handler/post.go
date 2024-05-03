package handler

import (
	"blog_api/er"
	model "blog_api/utils/models"
	"context"
	"errors"
	"strconv"

	"blog_api/pkg/posts"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PostHandler struct {
	log         *logrus.Logger
	postService *posts.Service
}

func newPostHandler(
	log *logrus.Logger,
	postService *posts.Service,
) *PostHandler {
	return &PostHandler{
		log:         log,
		postService: postService,
	}
}

func (h *PostHandler) UpsertPost(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &posts.Post{}
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
	err = h.postService.UpsertPost(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Success = true
	res.Data = req
	c.JSON(http.StatusOK, res)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = model.PostFilter{}
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
	if req.ID == 0 {
		id, ok := c.Params.Get("id")
		if !ok {
			err := errors.New("id not provided")
			err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
			return
		}
		ID, err := strconv.Atoi(id)
		if err != nil {
			err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
			return
		}
		req.ID = ID
	}
	data, pagination, err := h.postService.GetPosts(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Successfully retrieved posts"
	res.Success = true
	res.Data = data
	res.Meta = pagination
	c.JSON(http.StatusOK, res)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &posts.Post{}
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
	if req.ID == 0 {
		id, ok := c.Params.Get("id")
		if !ok {
			err := errors.New("id not provided")
			err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
			return
		}
		ID, err := strconv.Atoi(id)
		if err != nil {
			err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
			return
		}
		req.ID = ID
	}
	err = h.postService.UpdatePost(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Successfully retrieved posts"
	res.Success = true
	res.Data = req
	c.JSON(http.StatusOK, res)
}
func (h *PostHandler) DeletePost(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()

	id, ok := c.Params.Get("id")
	if !ok {
		err = errors.New("id empty in param")
		h.log.WithField("span", "delete_post").Info(err.Error())
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}
	ID, err := strconv.Atoi(id)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}

	err = h.postService.DeletePosts(dCtx, ID)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}
	res.Message = "Successfully Deleted post"
	res.Success = true
	res.Data = nil
	c.JSON(http.StatusOK, res)
}
