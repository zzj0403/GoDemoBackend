package http

import (
	"GoDemoBackend/apps"
	"GoDemoBackend/apps/host"
	"github.com/gin-gonic/gin"
)

var handler = &Handler{}

// 面向接口, 真正Service的实现, 在服务实例化的时候传递进行
// 也就是(CLI)  Start时候

type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	h.svc = apps.GetImpl(host.AppName).(host.Service)
}
func (h *Handler) Name() string {
	return host.AppName
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}

func init() {
	apps.GinRegistry(handler)
}
