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
func Run(domain string, port int) {
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

	e.GET("/", home(domain))
	e.GET("/gnparser", gnparser(domain))
	e.GET("/gnames", gnames(domain))
	e.GET("/gnames-beta", gnamesBeta(domain))
	e.GET("/gnmatcher", gnmatcher(domain))
	e.GET("/gnmatcher-beta", gnmatcherBeta(domain))
	e.GET("/gnfinder", gnfinder(domain))
	e.GET("/gnfinder-beta", gnfinderBeta(domain))

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
	Domain, DocJSON string
}

func home(domain string) func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{Domain: domain}
		return c.Render(http.StatusOK, "home", data)
	}
}

func gnfinder(domain string) func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{Domain: domain, DocJSON: "static/gnfinder/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnfinderBeta(domain string) func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{Domain: domain, DocJSON: "static/gnfinder-beta/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnparser(domain string) func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{Domain: domain, DocJSON: "static/gnparser/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnmatcher(domain string) func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{Domain: domain, DocJSON: "static/gnmatcher/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnmatcherBeta(domain string) func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{Domain: domain, DocJSON: "static/gnmatcher-beta/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnames(domain string) func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{Domain: domain, DocJSON: "static/gnames/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}

func gnamesBeta(domain string) func(echo.Context) error {
	return func(c echo.Context) error {
		data := Data{Domain: domain, DocJSON: "static/gnames-beta/openapi.json"}
		return c.Render(http.StatusOK, "api", data)
	}
}
