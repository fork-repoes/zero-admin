package loginlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:07
*/
type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginLogAdd 添加登录日志
func (l *LoginLogAddLogic) LoginLogAdd(in *sysclient.LoginLogAddReq) (*sysclient.LoginLogAddResp, error) {
	err := query.SysLoginLog.WithContext(l.ctx).Create(&model.SysLoginLog{
		UserName:   in.UserName,
		Status:     in.Status,
		IP:         in.Ip,
		CreateTime: time.Now(),
		CreateBy:   in.CreateBy,
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.LoginLogAddResp{}, nil
}
