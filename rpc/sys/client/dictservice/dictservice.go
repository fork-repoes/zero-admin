// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package dictservice

import (
	"context"

	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConfigAddReq          = sysclient.ConfigAddReq
	ConfigAddResp         = sysclient.ConfigAddResp
	ConfigDeleteReq       = sysclient.ConfigDeleteReq
	ConfigDeleteResp      = sysclient.ConfigDeleteResp
	ConfigListData        = sysclient.ConfigListData
	ConfigListReq         = sysclient.ConfigListReq
	ConfigListResp        = sysclient.ConfigListResp
	ConfigUpdateReq       = sysclient.ConfigUpdateReq
	ConfigUpdateResp      = sysclient.ConfigUpdateResp
	DeptAddReq            = sysclient.DeptAddReq
	DeptAddResp           = sysclient.DeptAddResp
	DeptDeleteReq         = sysclient.DeptDeleteReq
	DeptDeleteResp        = sysclient.DeptDeleteResp
	DeptListData          = sysclient.DeptListData
	DeptListReq           = sysclient.DeptListReq
	DeptListResp          = sysclient.DeptListResp
	DeptUpdateReq         = sysclient.DeptUpdateReq
	DeptUpdateResp        = sysclient.DeptUpdateResp
	DictAddReq            = sysclient.DictAddReq
	DictAddResp           = sysclient.DictAddResp
	DictDeleteReq         = sysclient.DictDeleteReq
	DictDeleteResp        = sysclient.DictDeleteResp
	DictListData          = sysclient.DictListData
	DictListReq           = sysclient.DictListReq
	DictListResp          = sysclient.DictListResp
	DictUpdateReq         = sysclient.DictUpdateReq
	DictUpdateResp        = sysclient.DictUpdateResp
	InfoReq               = sysclient.InfoReq
	InfoResp              = sysclient.InfoResp
	JobAddReq             = sysclient.JobAddReq
	JobAddResp            = sysclient.JobAddResp
	JobDeleteReq          = sysclient.JobDeleteReq
	JobDeleteResp         = sysclient.JobDeleteResp
	JobListData           = sysclient.JobListData
	JobListReq            = sysclient.JobListReq
	JobListResp           = sysclient.JobListResp
	JobUpdateReq          = sysclient.JobUpdateReq
	JobUpdateResp         = sysclient.JobUpdateResp
	LoginLogAddReq        = sysclient.LoginLogAddReq
	LoginLogAddResp       = sysclient.LoginLogAddResp
	LoginLogDeleteReq     = sysclient.LoginLogDeleteReq
	LoginLogDeleteResp    = sysclient.LoginLogDeleteResp
	LoginLogListData      = sysclient.LoginLogListData
	LoginLogListReq       = sysclient.LoginLogListReq
	LoginLogListResp      = sysclient.LoginLogListResp
	LoginReq              = sysclient.LoginReq
	LoginResp             = sysclient.LoginResp
	MenuAddReq            = sysclient.MenuAddReq
	MenuAddResp           = sysclient.MenuAddResp
	MenuDeleteReq         = sysclient.MenuDeleteReq
	MenuDeleteResp        = sysclient.MenuDeleteResp
	MenuListData          = sysclient.MenuListData
	MenuListReq           = sysclient.MenuListReq
	MenuListResp          = sysclient.MenuListResp
	MenuListTree          = sysclient.MenuListTree
	MenuUpdateReq         = sysclient.MenuUpdateReq
	MenuUpdateResp        = sysclient.MenuUpdateResp
	QueryMenuByRoleIdReq  = sysclient.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp = sysclient.QueryMenuByRoleIdResp
	ReSetPasswordReq      = sysclient.ReSetPasswordReq
	ReSetPasswordResp     = sysclient.ReSetPasswordResp
	RoleAddReq            = sysclient.RoleAddReq
	RoleAddResp           = sysclient.RoleAddResp
	RoleDeleteReq         = sysclient.RoleDeleteReq
	RoleDeleteResp        = sysclient.RoleDeleteResp
	RoleListData          = sysclient.RoleListData
	RoleListReq           = sysclient.RoleListReq
	RoleListResp          = sysclient.RoleListResp
	RoleUpdateReq         = sysclient.RoleUpdateReq
	RoleUpdateResp        = sysclient.RoleUpdateResp
	SysLogAddReq          = sysclient.SysLogAddReq
	SysLogAddResp         = sysclient.SysLogAddResp
	SysLogDeleteReq       = sysclient.SysLogDeleteReq
	SysLogDeleteResp      = sysclient.SysLogDeleteResp
	SysLogListData        = sysclient.SysLogListData
	SysLogListReq         = sysclient.SysLogListReq
	SysLogListResp        = sysclient.SysLogListResp
	UpdateConfigRoleReq   = sysclient.UpdateConfigRoleReq
	UpdateConfigRoleResp  = sysclient.UpdateConfigRoleResp
	UpdateMenuRoleReq     = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp    = sysclient.UpdateMenuRoleResp
	UpdateRoleRoleReq     = sysclient.UpdateRoleRoleReq
	UpdateRoleRoleResp    = sysclient.UpdateRoleRoleResp
	UserAddReq            = sysclient.UserAddReq
	UserAddResp           = sysclient.UserAddResp
	UserDeleteReq         = sysclient.UserDeleteReq
	UserDeleteResp        = sysclient.UserDeleteResp
	UserListData          = sysclient.UserListData
	UserListReq           = sysclient.UserListReq
	UserListResp          = sysclient.UserListResp
	UserStatusReq         = sysclient.UserStatusReq
	UserStatusResp        = sysclient.UserStatusResp
	UserUpdateReq         = sysclient.UserUpdateReq
	UserUpdateResp        = sysclient.UserUpdateResp

	DictService interface {
		DictAdd(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error)
		DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error)
		DictUpdate(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error)
		DictDelete(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error)
	}

	defaultDictService struct {
		cli zrpc.Client
	}
)

func NewDictService(cli zrpc.Client) DictService {
	return &defaultDictService{
		cli: cli,
	}
}

func (m *defaultDictService) DictAdd(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.DictAdd(ctx, in, opts...)
}

func (m *defaultDictService) DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.DictList(ctx, in, opts...)
}

func (m *defaultDictService) DictUpdate(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.DictUpdate(ctx, in, opts...)
}

func (m *defaultDictService) DictDelete(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.DictDelete(ctx, in, opts...)
}
