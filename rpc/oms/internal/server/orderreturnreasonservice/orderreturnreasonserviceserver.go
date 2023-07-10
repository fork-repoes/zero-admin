// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package server

import (
	"context"

	"zero-admin/rpc/oms/internal/logic/orderreturnreasonservice"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"
)

type OrderReturnReasonServiceServer struct {
	svcCtx *svc.ServiceContext
	omsclient.UnimplementedOrderReturnReasonServiceServer
}

func NewOrderReturnReasonServiceServer(svcCtx *svc.ServiceContext) *OrderReturnReasonServiceServer {
	return &OrderReturnReasonServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderReturnReasonServiceServer) OrderReturnReasonAdd(ctx context.Context, in *omsclient.OrderReturnReasonAddReq) (*omsclient.OrderReturnReasonAddResp, error) {
	l := orderreturnreasonservicelogic.NewOrderReturnReasonAddLogic(ctx, s.svcCtx)
	return l.OrderReturnReasonAdd(in)
}

func (s *OrderReturnReasonServiceServer) OrderReturnReasonList(ctx context.Context, in *omsclient.OrderReturnReasonListReq) (*omsclient.OrderReturnReasonListResp, error) {
	l := orderreturnreasonservicelogic.NewOrderReturnReasonListLogic(ctx, s.svcCtx)
	return l.OrderReturnReasonList(in)
}

func (s *OrderReturnReasonServiceServer) OrderReturnReasonUpdate(ctx context.Context, in *omsclient.OrderReturnReasonUpdateReq) (*omsclient.OrderReturnReasonUpdateResp, error) {
	l := orderreturnreasonservicelogic.NewOrderReturnReasonUpdateLogic(ctx, s.svcCtx)
	return l.OrderReturnReasonUpdate(in)
}

func (s *OrderReturnReasonServiceServer) OrderReturnReasonDelete(ctx context.Context, in *omsclient.OrderReturnReasonDeleteReq) (*omsclient.OrderReturnReasonDeleteResp, error) {
	l := orderreturnreasonservicelogic.NewOrderReturnReasonDeleteLogic(ctx, s.svcCtx)
	return l.OrderReturnReasonDelete(in)
}
