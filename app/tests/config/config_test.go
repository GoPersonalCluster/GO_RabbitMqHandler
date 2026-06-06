package math

import (
	"embed"
	"go_rabbitmqhandler/internal/config"
	"runtime"
	"testing"
)

//go:embed config/*
var cfgFiles embed.FS

func TestGetConfig(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	println("Current file:", filename)

	config := config.Config{}
	config.GetConfig()

}
