package loginlogservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:07
*/
type LoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginLogList 登录日志列表
func (l *LoginLogListLogic) LoginLogList(in *sysclient.LoginLogListReq) (*sysclient.LoginLogListResp, error) {
	q := query.SysLoginLog.WithContext(l.ctx)
	if len(in.UserName) > 0 {
		q = q.Where(query.SysLoginLog.UserName.Like("%" + in.UserName + "%"))
	}
	if len(in.Ip) > 0 {
		q = q.Where(query.SysLoginLog.UserName.Like("%" + in.Ip + "%"))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		in, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询登录记录列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*sysclient.LoginLogListData
	for _, log := range result {
		list = append(list, &sysclient.LoginLogListData{
			Id:         log.ID,
			UserName:   log.UserName,
			Status:     log.Status,
			Ip:         log.IP,
			CreateBy:   log.CreateBy,
			CreateTime: log.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询登录记录列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.LoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
