package math

import (
	"testing"

	"go_rabbitmqhandler/internal/config"
)

func TestGetEnvironmentConfig(t *testing.T) {
	envConfig := config.NewEnvironmentConfig()

	if envConfig == nil {
		t.Fatal("expected config, got nil")
	}
}