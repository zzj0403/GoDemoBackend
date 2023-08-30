package host

import "context"

type Service interface {
	CreateHost(ctx context.Context, host *Host) (*Host, error)
	QueryHost(ctx context.Context, req *QueryHostRequest) (*SetHost, error)
	DescribeHost(ctx context.Context, req *QueryHostRequest) (*Host, error)
	UpdateHost(ctx context.Context, req *UpdateHostRequest) (*Host, error)
	DeleteHost(ctx context.Context, req *DeleteHostRequest) (*Host, error)
}
