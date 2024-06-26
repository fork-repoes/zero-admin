package userservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReSetPasswordLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:12
*/
type ReSetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReSetPassword 重置用户密码
func (l *ReSetPasswordLogic) ReSetPassword(in *sysclient.ReSetPasswordReq) (*sysclient.ReSetPasswordResp, error) {

	q := query.SysUser
	now := time.Now()
	_, err := q.WithContext(l.ctx).Updates(&model.SysUser{
		ID:         in.Id,
		Password:   "123456",
		Salt:       "123456",
		UpdateBy:   &in.UpdateBy,
		UpdateTime: &now,
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.ReSetPasswordResp{}, nil
}
