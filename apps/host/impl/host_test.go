package impl_test

import (
	"GoDemoBackend/conf"
	"GoDemoBackend/pkg/host"
	"GoDemoBackend/pkg/host/impl"
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	service host.Service
)

func TestCreateHost(t *testing.T) {

	should := assert.New(t)
	ins := host.NewHost()
	ins.Name = "test"
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func init() {

	// 初始化日志
	zap.DevelopmentSetup()

	// 初始化配置
	conf.LoadConfigFromToml("../impl/host/pkg/etc/demo.toml")
	service = impl.NewHostServiceImpl()
}
