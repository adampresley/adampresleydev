package main

import (
	"context"
	"embed"
	"log/slog"
	"net/http"

	"github.com/adampresley/adampresleydev/cmd/adampresleydev/internal/configuration"
	"github.com/adampresley/adampresleydev/cmd/adampresleydev/internal/home"
	"github.com/adampresley/httphelpers/responses"
	"github.com/adampresley/mux"
	"github.com/adampresley/rendering"
)

var (
	Version string = "development"
	appName string = "adampresleydev"

	//go:embed app
	appFS embed.FS

	/* Services */
	renderer rendering.TemplateRenderer

	/* Controllers */
	homeController home.HomeHandlers
)

func main() {
	var (
		err error
	)

	config := configuration.LoadConfig()
	shutdownCtx, stopApp := context.WithCancel(context.Background())

	setupLogger(&config, Version)

	slog.Info("configuration loaded",
		"app", appName,
		"version", Version,
		"loglevel", config.LogLevel,
		"host", config.Host,
	)

	slog.Debug("setting up...")

	/*
	 * Setup services
	 */
	if renderer, err = rendering.NewGoTemplateRenderer(appFS); err != nil {
		panic(err)
	}

	/*
	 * Setup controllers
	 */
	homeController = home.NewHomeController(home.HomeControllerConfig{
		Config:   &config,
		Renderer: renderer,
	})

	/*
	 * Setup router and http server
	 */
	slog.Debug("setting up routes...")

	routes := []mux.Route{
		{Path: "GET /heartbeat", HandlerFunc: heartbeat},
		{Path: "GET /experience", HandlerFunc: homeController.ExperiencePage},
		{Path: "GET /", HandlerFunc: homeController.HomePage},
	}

	mux := mux.Setup(
		&config,
		routes,
		shutdownCtx,
		stopApp,

		mux.WithStaticContent("app", "/static/", appFS),
		mux.UseGzip(),
		mux.UseGzipForStaticFiles(),
		mux.WithDebug(Version == "development"),
	)

	slog.Info("server started")
	mux.Start()

	slog.Info("server stopped")
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	responses.TextOK(w, "OK")
}
