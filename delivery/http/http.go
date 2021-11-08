package http

import (
	"strings"

	"github.com/Auth-Service/config"
	"github.com/Auth-Service/delivery/http/v1/auth"
	//"github.com/Auth-Service/delivery/http/auth"
	"github.com/cnson19700/pkg/middleware"
	"github.com/labstack/echo/v4"

	"github.com/Auth-Service/repository"
	"github.com/Auth-Service/usecase"
)

// NewHTTPHandler .
func NewHTTPHandler(repo *repository.Repository, ucase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	cfg := config.GetConfig()

	skipper := func(c echo.Context) bool {
		p := c.Request().URL.Path

		return strings.Contains(p, "/health_check") ||
			strings.Contains(p, "/login") ||
			strings.Contains(p, "/register")
	}

	// loggerCfg := middleware.DefaultLoggerConfig
	// loggerCfg.Skipper = skipper

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	Skipper:      middleware.DefaultSkipper,
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{
	// 		http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions,
	// 	},
	// }))
	// e.Use(middleware.LoggerWithConfig(loggerCfg))
	// e.Use(middleware.Recover())
	// e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(sentryecho.New(sentryecho.Options{
	// 	Repanic: true,
	// }))
	e.Use(middleware.Auth(cfg.Jwt.Key, skipper, false))

	// if cfg.Endpoints.DatadogAgentEndpoint != "" {
	// 	e.Use(myMiddleware.DataDogTrace("hus-echo"))
	// }

	// e.GET("/health_check", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "ok")
	// })

	//e.GET("/docs/*", echoSwagger.WrapHandler)

	apiV1 := e.Group("/v1")

	auth.Init(apiV1.Group("/auth"), ucase)

	return e
}
