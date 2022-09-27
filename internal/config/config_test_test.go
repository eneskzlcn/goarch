package config_test

import (
	config "<path-to-config>"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	environment := "test"
	config, err := config.LoadConfig[config.Config]("<path-to-.dev/>", environment, "yaml")
	assert.Nil(t, err)
	assert.Equal(t, "<should be same with the username defined on <path-to-.dev/test.yaml> >", config.Db.Username)
}
