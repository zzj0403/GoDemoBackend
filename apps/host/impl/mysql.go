package impl

import (
	"GoDemoBackend/apps"
	"GoDemoBackend/apps/host"
	"GoDemoBackend/conf"
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现静态检查
// var _ host.Service = (*HostServiceImpl)(nil)

// 吧对象注册和对象的初始化 分开
var impl = &HostServiceImpl{}

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

func (i *HostServiceImpl) Config() {
	i.l = zap.L().Named("Host")
	i.db = conf.C().MySQL.GetDB()
}
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

func init() {
	apps.ImplRegistry(impl)
}
