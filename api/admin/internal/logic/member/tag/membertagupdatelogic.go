package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagUpdateLogic {
	return MemberTagUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagUpdateLogic) MemberTagUpdate(req types.UpdateMemberTagReq) (*types.UpdateMemberTagResp, error) {
	_, err := l.svcCtx.MemberTagService.MemberTagUpdate(l.ctx, &umsclient.MemberTagUpdateReq{
		Id:                req.Id,
		Name:              req.Name,
		FinishOrderCount:  req.FinishOrderCount,
		FinishOrderAmount: req.FinishOrderAmount,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员标签信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新会员标签失败")
	}

	return &types.UpdateMemberTagResp{
		Code:    "000000",
		Message: "更新会员标签成功",
	}, nil
}
