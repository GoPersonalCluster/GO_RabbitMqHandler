package config_test

import (
	"testing"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/config"
)

func TestGetEnvironmentConfig(t *testing.T) {
	envConfig := config.NewEnvironmentConfig()

	if envConfig == nil {
		t.Fatal("expected config, got nil")
	}
}
