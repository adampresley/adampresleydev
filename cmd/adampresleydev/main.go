package main

import (
	"context"
	"embed"
	"log/slog"
	"net/http"

	"github.com/adampresley/adamgokit/httphelpers"
	"github.com/adampresley/adamgokit/mux2"
	"github.com/adampresley/adamgokit/rendering"
	"github.com/adampresley/adampresleydev/cmd/adampresleydev/internal/configuration"
	"github.com/adampresley/adampresleydev/cmd/adampresleydev/internal/home"
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

	routes := []mux2.Route{
		{Path: "GET /heartbeat", HandlerFunc: heartbeat},
		{Path: "GET /experience", HandlerFunc: homeController.ExperiencePage},
		{Path: "GET /", HandlerFunc: homeController.HomePage},
	}

	mux := mux2.Setup(
		&config,
		routes,
		shutdownCtx,
		stopApp,

		mux2.WithStaticContent("app", "/static/", appFS),
		mux2.UseGzip(),
		mux2.UseGzipForStaticFiles(),
		mux2.WithDebug(Version == "development"),
	)

	slog.Info("server started")
	mux.Start()

	slog.Info("server stopped")
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	httphelpers.TextOK(w, "OK")
}
