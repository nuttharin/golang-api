package httpserver

import (
	"context"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Context interface {
	Bind(interface{}) error
	JSON(code int, obj interface{})
	Redirect(code int, location string)
	GetQuery(string) string
	GetQueryInt(string) (int, error)
	GetPath() string
	GetParam(string) string
	GetParamInt(string) (int, error)
	AttachError(err error)
	Set(key string, value any)
	Get(key string) (value any, exists bool)
	BindForm(interface{}) error
	SetHeader(string, string)
	GetHeader(string) string
	Data(httpCode int, contentType string, data []byte)
	GetToken() string
	GetLang() *string
	GetRequestCtx() context.Context
	GetUserAgent() string
	GetTraceId() string
	GetHostname() string
	GetRequestId() string
	GetReferer() string
	GetGinContext() *gin.Context
}

type ginContext struct {
	*gin.Context
}

func (c *ginContext) GetHostname() string {
	host := c.Context.Request.Host
	if host != "" {
		return host
	}
	return "Unknown"
}

func (c *ginContext) GetPath() string {
	path := c.FullPath()
	if path != "" {
		return path
	}
	return "Unknown"
}

func (c *ginContext) GetReferer() string {
	referer := c.Context.GetHeader("A-Referer")
	if referer != "" {
		return referer
	}
	return "Unknown"
}

func (c *ginContext) GetRequestId() string {
	requestId := c.Context.GetHeader("A-Rid")
	if requestId != "" {
		return requestId
	}
	return uuid.New().String()
}

func (c *ginContext) GetTraceId() string {
	traceId := c.Context.GetHeader("A-Tid")
	if traceId != "" {
		return traceId
	}
	return uuid.New().String()
}

func (c *ginContext) GetUserAgent() string {
	userAgent := c.Context.GetHeader("User-Agent")
	if userAgent != "" {
		return userAgent
	}
	return "Unknown"
}

func (c *ginContext) GetRequestCtx() context.Context {
	return c.Request.Context()
}

func (c *ginContext) JSON(code int, obj interface{}) {
	c.Context.PureJSON(code, obj)
}

func (c *ginContext) Bind(obj interface{}) error {
	if err := c.Context.ShouldBindJSON(obj); err != nil {

		return &BindError{Message: err.Error()}
	}

	return nil
}

func (c *ginContext) GetQuery(key string) string {
	return c.Context.Query(key)
}

func (c *ginContext) Set(key string, value any) {
	c.Context.Set(key, value)
}

func (c *ginContext) Get(key string) (value any, exists bool) {
	return c.Context.Get(key)
}

func (c *ginContext) GetParamInt(key string) (int, error) {
	param := c.Context.Param(key)

	v, err := strconv.Atoi(param)
	if err != nil {
		return 0, err

	}

	return v, nil
}

func (c *ginContext) GetQueryInt(key string) (int, error) {
	param := c.Context.Query(key)

	v, err := strconv.Atoi(param)
	if err != nil {
		return 0, err

	}

	return v, nil
}

func (c *ginContext) GetParam(key string) string {
	return c.Context.Param(key)
}

func (c *ginContext) AttachError(err error) {
	c.Context.Error(err)
}

func (c *ginContext) SetHeader(key string, value string) {
	c.Header(key, value)
}

func (c *ginContext) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *ginContext) BindForm(v interface{}) error {
	return c.Context.ShouldBind(v)
}

func (c *ginContext) Data(httpCode int, contentType string, data []byte) {
	c.Context.Data(httpCode, contentType, data)
}

func (c *ginContext) GetToken() string {
	Auth := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(Auth, "Bearer ")
	return token
}

func (c *ginContext) GetLang() *string {
	lang := c.Context.GetString("lang")

	if lang == "" {
		return nil
	}

	return &lang
}

func (c *ginContext) GetGinContext() *gin.Context {
	return c.Context
}

func (c *ginContext) Redirect(code int, location string) {
	c.Context.Redirect(code, location)
}

func NewContext(c *gin.Context) Context {
	return &ginContext{Context: c}
}
