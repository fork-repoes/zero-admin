// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"zero-admin/rpc/sms/internal/logic/couponproductcategoryrelationservice"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"
)

type CouponProductCategoryRelationServiceServer struct {
	svcCtx *svc.ServiceContext
	smsclient.UnimplementedCouponProductCategoryRelationServiceServer
}

func NewCouponProductCategoryRelationServiceServer(svcCtx *svc.ServiceContext) *CouponProductCategoryRelationServiceServer {
	return &CouponProductCategoryRelationServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CouponProductCategoryRelationServiceServer) CouponProductCategoryRelationAdd(ctx context.Context, in *smsclient.CouponProductCategoryRelationAddReq) (*smsclient.CouponProductCategoryRelationAddResp, error) {
	l := couponproductcategoryrelationservicelogic.NewCouponProductCategoryRelationAddLogic(ctx, s.svcCtx)
	return l.CouponProductCategoryRelationAdd(in)
}

func (s *CouponProductCategoryRelationServiceServer) CouponProductCategoryRelationList(ctx context.Context, in *smsclient.CouponProductCategoryRelationListReq) (*smsclient.CouponProductCategoryRelationListResp, error) {
	l := couponproductcategoryrelationservicelogic.NewCouponProductCategoryRelationListLogic(ctx, s.svcCtx)
	return l.CouponProductCategoryRelationList(in)
}

func (s *CouponProductCategoryRelationServiceServer) CouponProductCategoryRelationUpdate(ctx context.Context, in *smsclient.CouponProductCategoryRelationUpdateReq) (*smsclient.CouponProductCategoryRelationUpdateResp, error) {
	l := couponproductcategoryrelationservicelogic.NewCouponProductCategoryRelationUpdateLogic(ctx, s.svcCtx)
	return l.CouponProductCategoryRelationUpdate(in)
}

func (s *CouponProductCategoryRelationServiceServer) CouponProductCategoryRelationDelete(ctx context.Context, in *smsclient.CouponProductCategoryRelationDeleteReq) (*smsclient.CouponProductCategoryRelationDeleteResp, error) {
	l := couponproductcategoryrelationservicelogic.NewCouponProductCategoryRelationDeleteLogic(ctx, s.svcCtx)
	return l.CouponProductCategoryRelationDelete(in)
}
