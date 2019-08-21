package web

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	*Config
	router    *echo.Echo
	templates *template.Template
}

func NewServer(cfg *Config) *Server {

	// Create new server
	server := &Server{
		Config:    cfg,
		router:    echo.New(),
		templates: loadTemplates(cfg.WebDir),
	}

	// Create endpoints
	server.router.GET("/", server.Index)

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
	return server.router.Start(fmt.Sprintf(":%d", server.Port))
}

// Stop serving
func (server *Server) Stop(ctx context.Context) error {
	return server.router.Shutdown(ctx)
}

// Serve the index page
func (server *Server) Index(context echo.Context) error {
	return context.Render(http.StatusOK, "index.gohtml", nil)
}

// Helper function to load templates from webdir
func loadTemplates(webDir string) *template.Template {
	return template.Must(template.ParseGlob(path.Join(webDir, "templates", "*.gohtml")))
}
