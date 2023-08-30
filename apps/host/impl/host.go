package impl

import (
	"GoDemoBackend/apps/host"
	"context"
)

// CreateHost  业务处理层 controller
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 校验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	// 默认值填充
	ins.InjectDefault()
	// 有dao模块 负责 把对象入库
	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, ins *host.QueryHostRequest) (*host.SetHost, error) {

	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, ins *host.QueryHostRequest) (*host.Host, error) {

	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, ins *host.UpdateHostRequest) (*host.Host, error) {

	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, ins *host.DeleteHostRequest) (*host.Host, error) {

	return nil, nil
}
