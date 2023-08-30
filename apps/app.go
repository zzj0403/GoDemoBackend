package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// IOC 容器层，管理服务实例

var (
	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

func ImplRegistry(svc ImplService) {
	// 服务实力注册到 implApps ，map当中
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s is exist", svc.Name()))
	}
	implApps[svc.Name()] = svc
	//if v, ok := svc.(host.Service); ok {
	//	HostService = v
	//}
}
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}
	return nil
}

func GinRegistry(svc GinService) {
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s is exist", svc.Name()))
	}
	ginApps[svc.Name()] = svc

}

// 已经加载完成的Gin App由哪些
func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}

	return
}

type ImplService interface {
	Config()
	Name() string
}

func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}

}

func InitGin(r gin.IRouter) {
	// 先初始化好所有对象
	for _, v := range ginApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

type GinService interface {
	Registry(r gin.IRouter)
	Name() string
	Config()
}
