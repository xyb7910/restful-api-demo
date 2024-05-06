package conf

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal(config.App.Name, "restful-api-demo")
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("APP_NAME", "restful-api-demo")
	err := LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal(config.App.Name, "restful-api-demo")
	}
}
