package conf_test

import (
	"GoDemoBackend/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("demo", conf.C().App.Name)
	}
}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()
	}
}
