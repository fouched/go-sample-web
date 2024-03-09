package handlers

import "github.com/fouched/go-sample-web/internal/config"

var Instance *HandlerConfig

type HandlerConfig struct {
	App *config.AppConfig
}

func NewConfig(a *config.AppConfig) *HandlerConfig {
	return &HandlerConfig{
		App: a,
	}
}

func NewHandlers(h *HandlerConfig) {
	Instance = h
}
