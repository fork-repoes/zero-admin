// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"zero-admin/rpc/pms/internal/logic/commentreplayservice"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"
)

type CommentReplayServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedCommentReplayServiceServer
}

func NewCommentReplayServiceServer(svcCtx *svc.ServiceContext) *CommentReplayServiceServer {
	return &CommentReplayServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentReplayServiceServer) CommentReplayAdd(ctx context.Context, in *pmsclient.CommentReplayAddReq) (*pmsclient.CommentReplayAddResp, error) {
	l := commentreplayservicelogic.NewCommentReplayAddLogic(ctx, s.svcCtx)
	return l.CommentReplayAdd(in)
}

func (s *CommentReplayServiceServer) CommentReplayList(ctx context.Context, in *pmsclient.CommentReplayListReq) (*pmsclient.CommentReplayListResp, error) {
	l := commentreplayservicelogic.NewCommentReplayListLogic(ctx, s.svcCtx)
	return l.CommentReplayList(in)
}

func (s *CommentReplayServiceServer) CommentReplayUpdate(ctx context.Context, in *pmsclient.CommentReplayUpdateReq) (*pmsclient.CommentReplayUpdateResp, error) {
	l := commentreplayservicelogic.NewCommentReplayUpdateLogic(ctx, s.svcCtx)
	return l.CommentReplayUpdate(in)
}

func (s *CommentReplayServiceServer) CommentReplayDelete(ctx context.Context, in *pmsclient.CommentReplayDeleteReq) (*pmsclient.CommentReplayDeleteResp, error) {
	l := commentreplayservicelogic.NewCommentReplayDeleteLogic(ctx, s.svcCtx)
	return l.CommentReplayDelete(in)
}
