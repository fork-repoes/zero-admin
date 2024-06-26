package memberreceiveaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReceiveAddressAddLogic 收货地址
/*
Author: LiuFeiHua
Date: 2023/11/30 11:27
*/
type MemberReceiveAddressAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressAddLogic {
	return &MemberReceiveAddressAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberReceiveAddressAdd 添加会员收货地址
func (l *MemberReceiveAddressAddLogic) MemberReceiveAddressAdd(in *umsclient.MemberReceiveAddressAddReq) (*umsclient.MemberReceiveAddressAddResp, error) {
	q := query.UmsMemberReceiveAddress

	//如果新增的地址为默认地址,则需要把之前的默认地址去除默认标识
	if in.DefaultStatus == 1 {
		//查询会员所有有地址
		memberList, _ := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Find()

		for _, address := range memberList {
			//判断是否为默认,如果是,则修改
			if address.DefaultStatus == 1 {
				address.DefaultStatus = 0
				_, _ = q.WithContext(l.ctx).Where(q.ID.Eq(address.ID)).Updates(address)
			}
		}
	}

	err := q.WithContext(l.ctx).Create(&model.UmsMemberReceiveAddress{
		MemberID:      in.MemberId,
		Name:          in.Name,
		PhoneNumber:   in.PhoneNumber,
		DefaultStatus: in.DefaultStatus,
		PostCode:      in.PostCode,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReceiveAddressAddResp{}, nil
}
