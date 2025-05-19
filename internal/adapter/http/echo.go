package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type EchoWrapper struct {
	*echo.Echo

	handler *HttpHandler
	cmw     *CustomMiddleware
}

func NewEchoWrapper(mw *CustomMiddleware, handler *HttpHandler) *EchoWrapper {
	e := echo.New()
	echoWrapper := &EchoWrapper{
		Echo:    e,
		handler: handler,
		cmw:     mw,
	}
	echoWrapper.initRoute()

	return echoWrapper
}

func (e *EchoWrapper) initRoute() {
	e.Use(e.cmw.loggingMiddleware())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, ")
	})

	e.GET("/hc", func(c echo.Context) error {

		return c.Redirect(http.StatusOK, "/")
	})
	e.GET("/health", func(c echo.Context) error {

		return c.Redirect(http.StatusFound, "/")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
