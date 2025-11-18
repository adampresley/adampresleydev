package home

import (
	"net/http"

	"github.com/adampresley/adamgokit/rendering"
	"github.com/adampresley/adampresleydev/cmd/adampresleydev/internal/configuration"
)

type HomeHandlers interface {
	HomePage(w http.ResponseWriter, r *http.Request)
	ExperiencePage(w http.ResponseWriter, r *http.Request)
}

type HomeControllerConfig struct {
	Config   *configuration.Config
	Renderer rendering.TemplateRenderer
}

type HomeController struct {
	config   *configuration.Config
	renderer rendering.TemplateRenderer
}

func NewHomeController(config HomeControllerConfig) HomeController {
	return HomeController{
		config:   config.Config,
		renderer: config.Renderer,
	}
}
