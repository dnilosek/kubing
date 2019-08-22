package web

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/dnilosek/kubing/app/lib/database"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	*Config
	router    *echo.Echo
	templates *template.Template
	database  *database.DB
}

func NewServer(cfg *Config, db *database.DB) *Server {

	// Create new server
	server := &Server{
		Config:    cfg,
		router:    echo.New(),
		templates: loadTemplates(cfg.WebDir),
		database:  db,
	}

	// Create endpoints
	server.router.GET("/", server.Index)
	server.router.GET("/reset", server.ResetCount)

	// Clear startup info printing
	server.router.HideBanner = true
	server.router.HidePort = true

	// Attach render function
	server.router.Renderer = server

	server.router.Use(middleware.Logger())
	return server
}

// Create renderer
func (server *Server) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return server.templates.ExecuteTemplate(w, name, data)
}

// Start serving
func (server *Server) Start() error {
	resetVisitorCount(server)
	return server.router.Start(fmt.Sprintf(":%d", server.Port))
}

// Stop serving
func (server *Server) Stop(ctx context.Context) error {
	return server.router.Shutdown(ctx)
}

// Serve the index page
func (server *Server) Index(context echo.Context) error {
	count, err := incrementVisitorCount(server)
	if err != nil {
		return err
	}
	return context.Render(http.StatusOK, "index.gohtml", count)
}

func (server *Server) ResetCount(context echo.Context) error {
	if err != nil {
		return err
	}
	return context.Redirect(http.StatusSeeOther, "/")
}

// Helper function to load templates from webdir
func loadTemplates(webDir string) *template.Template {
	return template.Must(template.ParseGlob(path.Join(webDir, "templates", "*.gohtml")))
}

// Helper function to increment visitor count
func incrementVisitorCount(server *Server) (int, error) {
	countStr, err := server.database.Get("visitor-count")
	if err != nil && err != redis.Nil {
		return 0, err
	}
	count := 0
	if countStr != "" {
		count, err = strconv.Atoi(countStr)
		if err != nil {
			return 0, err
		}
	}
	count += 1
	err = server.database.Set("visitor-count", strconv.Itoa(count))
	return count, err
}

func resetVisitorCount(server *Server) error {
	return server.database.Set("visitor-count", "0")
}
