package httpserver

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type defaultRouter interface {
	GET(relativePath string, handlers ...interface{})
	POST(relativePath string, handlers ...interface{})
	PUT(relativePath string, handlers ...interface{})
	PATCH(relativePath string, handlers ...interface{})
	DELETE(relativePath string, handlers ...interface{})
	Use(middleware ...gin.HandlerFunc)
}

type Router interface {
	defaultRouter
	Group(relativePath string, handlers ...interface{}) RouterGroup
}

type RouterGroup interface {
	defaultRouter
	Group(relativePath string, handlers ...interface{}) RouterGroup
}

type ginRouter struct {
	engine *gin.Engine
}

type ginRouterGroup struct {
	engine *gin.RouterGroup
}

func NewRouter(isDebug bool, middlewares ...gin.HandlerFunc) *ginRouter {
	var r *gin.Engine

	if isDebug {
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	}

	// Add the custom middlewares
	r.Use(middlewares...)
	r.Use(Localization(false))
	r.Use(CORSMiddleware())

	// initial swagger doc
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &ginRouter{engine: r}
}

// ==== implements Router ==== //
func (r *ginRouter) GET(relativePath string, handlers ...interface{}) {
	r.engine.GET(relativePath, convertHandlersToGin(handlers)...)
}

func (r *ginRouter) POST(relativePath string, handlers ...interface{}) {
	r.engine.POST(relativePath, convertHandlersToGin(handlers)...)
}

func (r *ginRouter) PUT(relativePath string, handlers ...interface{}) {
	r.engine.PUT(relativePath, convertHandlersToGin(handlers)...)
}

func (rg *ginRouter) PATCH(relativePath string, handlers ...interface{}) {
	rg.engine.PATCH(relativePath, convertHandlersToGin(handlers)...)
}

func (r *ginRouter) DELETE(relativePath string, handlers ...interface{}) {
	r.engine.DELETE(relativePath, convertHandlersToGin(handlers)...)
}

func (r *ginRouter) Group(relativePath string, handlers ...interface{}) RouterGroup {
	group := r.engine.Group(relativePath, convertHandlersToGin(handlers)...)
	return &ginRouterGroup{engine: group}
}

func (r *ginRouter) Use(middleware ...gin.HandlerFunc) {
	r.engine.Use(middleware...)
}

// ==== End  implements Router ==== //

// ==== implements RouterGroup ==== //
func (rg *ginRouterGroup) GET(relativePath string, handlers ...interface{}) {
	rg.engine.GET(relativePath, convertHandlersToGin(handlers)...)
}

func (rg *ginRouterGroup) DELETE(relativePath string, handlers ...interface{}) {
	rg.engine.DELETE(relativePath, convertHandlersToGin(handlers)...)
}

func (rg *ginRouterGroup) POST(relativePath string, handlers ...interface{}) {
	rg.engine.POST(relativePath, convertHandlersToGin(handlers)...)
}

func (rg *ginRouterGroup) PUT(relativePath string, handlers ...interface{}) {
	rg.engine.PUT(relativePath, convertHandlersToGin(handlers)...)
}

func (rg *ginRouterGroup) PATCH(relativePath string, handlers ...interface{}) {
	rg.engine.PATCH(relativePath, convertHandlersToGin(handlers)...)
}

func (rg *ginRouterGroup) Use(middleware ...gin.HandlerFunc) {
	rg.engine.Use(middleware...)
}

func (rg *ginRouterGroup) Group(relativePath string, handlers ...interface{}) RouterGroup {
	group := rg.engine.Group(relativePath, convertHandlersToGin(handlers)...)
	return &ginRouterGroup{engine: group}
}

// ==== End implements RouterGroup ==== //

func convertHandlersToGin(handlers []interface{}) []gin.HandlerFunc {
	var ginHandlers []gin.HandlerFunc

	for _, h := range handlers {

		if handlerFunc, ok := h.(func(c *gin.Context)); ok {
			ginHandlers = append(ginHandlers, gin.HandlerFunc(handlerFunc))
		} else if ginHandler, ok := h.(gin.HandlerFunc); ok {
			ginHandlers = append(ginHandlers, ginHandler)
		} else if handlerFunc, ok := h.(func(c Context)); ok {
			ginHandlers = append(ginHandlers, convertToGinHandler(handlerFunc))
		} else {
			panic("unimplemented")
		}
	}

	return ginHandlers
}

func convertToGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&ginContext{Context: c})
	}
}
