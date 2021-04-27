package web

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const withLogs = true

//go:embed static
var static embed.FS

// Run starts HTTP/1 service for scientific names verification.
func Run(port int) {
	var err error
	log.Printf("Starting the HTTP API server on port %d.", port)
	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	if withLogs {
		e.Use(middleware.Logger())
	}

	e.Renderer, err = NewTemplate()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", home())
	e.GET("/gnparser", gnparser())
	e.GET("/gnames", gnames())
	e.GET("/gnmatcher", gnmatcher())
	e.GET("/gnfinder", gnfinder())

	fs := http.FileServer(http.FS(static))
	e.GET("/static/*", echo.WrapHandler(fs))

	addr := fmt.Sprintf(":%d", port)
	s := &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}

type Data struct {
	DocJSON string
}

func home() func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{}
		return c.Render(http.StatusOK, "home", data)
	}
}

func gnfinder() func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{DocJSON: "static/gnfinder/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnparser() func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{DocJSON: "static/gnparser/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnmatcher() func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{DocJSON: "static/gnmatcher/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnames() func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{DocJSON: "static/gnames/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}
