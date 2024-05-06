// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/logic/subjectproductrelationservice"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
)

type SubjectProductRelationServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedSubjectProductRelationServiceServer
}

func NewSubjectProductRelationServiceServer(svcCtx *svc.ServiceContext) *SubjectProductRelationServiceServer {
	return &SubjectProductRelationServiceServer{
		svcCtx: svcCtx,
	}
}

// 专题关联
func (s *SubjectProductRelationServiceServer) SubjectProductRelationAdd(ctx context.Context, in *cmsclient.SubjectProductRelationAddReq) (*cmsclient.SubjectProductRelationAddResp, error) {
	l := subjectproductrelationservicelogic.NewSubjectProductRelationAddLogic(ctx, s.svcCtx)
	return l.SubjectProductRelationAdd(in)
}

func (s *SubjectProductRelationServiceServer) SubjectProductRelationList(ctx context.Context, in *cmsclient.SubjectProductRelationListReq) (*cmsclient.SubjectProductRelationListResp, error) {
	l := subjectproductrelationservicelogic.NewSubjectProductRelationListLogic(ctx, s.svcCtx)
	return l.SubjectProductRelationList(in)
}
