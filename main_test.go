package main

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	assert.Equal(t, viper.AllSettings(), map[string]interface{}{
		"serve_host":  "localhost",
		"serve_port":  5002,
		"debug":       false,
		"serve_neo4j": "http://localhost:7474/db/data/",
	})
}
