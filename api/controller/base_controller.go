package controller

import (
	"encoding/json"
	"net/http"

	"github.com/hayrullahcansu/cachy/core/request/context"
)

type BaseController struct {
	httpContext *context.HttpContext
}

func NewBaseController(w http.ResponseWriter, r *http.Request) *BaseController {
	return &BaseController{
		httpContext: &context.HttpContext{
			W: w,
			R: r,
		},
	}
}

func (c *BaseController) Ok() {
	c.setContentTypeJson()
	c.httpContext.W.WriteHeader(http.StatusOK)
}
func (c *BaseController) OkWithBody(data interface{}) {
	c.setContentTypeJson()
	c.httpContext.W.WriteHeader(http.StatusOK)
	json.NewEncoder(c.httpContext.W).Encode(data)
}

func (c *BaseController) InternalServerError() {
	c.setContentTypeJson()
	c.httpContext.W.WriteHeader(http.StatusInternalServerError)
}
func (c *BaseController) InternalServerErrorWithBody(data interface{}) {
	c.setContentTypeJson()
	c.httpContext.W.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(c.httpContext.W).Encode(data)
}

func (c *BaseController) NotFound() {
	c.setContentTypeJson()
	c.httpContext.W.WriteHeader(http.StatusNotFound)
}
func (c *BaseController) NotFoundWithBody(data interface{}) {
	c.setContentTypeJson()
	c.httpContext.W.WriteHeader(http.StatusNotFound)
	json.NewEncoder(c.httpContext.W).Encode(data)
}

func (c *BaseController) setContentTypeJson() {
	c.httpContext.W.Header().Set("Content-Type", "application/json")
}
