// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package flashpromotionservice

import (
	"context"

	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CouponAddReq                            = smsclient.CouponAddReq
	CouponAddResp                           = smsclient.CouponAddResp
	CouponDeleteReq                         = smsclient.CouponDeleteReq
	CouponDeleteResp                        = smsclient.CouponDeleteResp
	CouponFindByIdReq                       = smsclient.CouponFindByIdReq
	CouponFindByIdResp                      = smsclient.CouponFindByIdResp
	CouponFindByIdsReq                      = smsclient.CouponFindByIdsReq
	CouponFindByIdsResp                     = smsclient.CouponFindByIdsResp
	CouponHistoryAddReq                     = smsclient.CouponHistoryAddReq
	CouponHistoryAddResp                    = smsclient.CouponHistoryAddResp
	CouponHistoryDeleteReq                  = smsclient.CouponHistoryDeleteReq
	CouponHistoryDeleteResp                 = smsclient.CouponHistoryDeleteResp
	CouponHistoryListData                   = smsclient.CouponHistoryListData
	CouponHistoryListReq                    = smsclient.CouponHistoryListReq
	CouponHistoryListResp                   = smsclient.CouponHistoryListResp
	CouponHistoryUpdateReq                  = smsclient.CouponHistoryUpdateReq
	CouponHistoryUpdateResp                 = smsclient.CouponHistoryUpdateResp
	CouponListData                          = smsclient.CouponListData
	CouponListReq                           = smsclient.CouponListReq
	CouponListResp                          = smsclient.CouponListResp
	CouponProductCategoryRelationAddReq     = smsclient.CouponProductCategoryRelationAddReq
	CouponProductCategoryRelationAddResp    = smsclient.CouponProductCategoryRelationAddResp
	CouponProductCategoryRelationDeleteReq  = smsclient.CouponProductCategoryRelationDeleteReq
	CouponProductCategoryRelationDeleteResp = smsclient.CouponProductCategoryRelationDeleteResp
	CouponProductCategoryRelationListData   = smsclient.CouponProductCategoryRelationListData
	CouponProductCategoryRelationListReq    = smsclient.CouponProductCategoryRelationListReq
	CouponProductCategoryRelationListResp   = smsclient.CouponProductCategoryRelationListResp
	CouponProductCategoryRelationUpdateReq  = smsclient.CouponProductCategoryRelationUpdateReq
	CouponProductCategoryRelationUpdateResp = smsclient.CouponProductCategoryRelationUpdateResp
	CouponProductRelationAddReq             = smsclient.CouponProductRelationAddReq
	CouponProductRelationAddResp            = smsclient.CouponProductRelationAddResp
	CouponProductRelationDeleteReq          = smsclient.CouponProductRelationDeleteReq
	CouponProductRelationDeleteResp         = smsclient.CouponProductRelationDeleteResp
	CouponProductRelationListData           = smsclient.CouponProductRelationListData
	CouponProductRelationListReq            = smsclient.CouponProductRelationListReq
	CouponProductRelationListResp           = smsclient.CouponProductRelationListResp
	CouponProductRelationUpdateReq          = smsclient.CouponProductRelationUpdateReq
	CouponProductRelationUpdateResp         = smsclient.CouponProductRelationUpdateResp
	CouponUpdateReq                         = smsclient.CouponUpdateReq
	CouponUpdateResp                        = smsclient.CouponUpdateResp
	FlashPromotionAddReq                    = smsclient.FlashPromotionAddReq
	FlashPromotionAddResp                   = smsclient.FlashPromotionAddResp
	FlashPromotionDeleteReq                 = smsclient.FlashPromotionDeleteReq
	FlashPromotionDeleteResp                = smsclient.FlashPromotionDeleteResp
	FlashPromotionListByDateReq             = smsclient.FlashPromotionListByDateReq
	FlashPromotionListByDateResp            = smsclient.FlashPromotionListByDateResp
	FlashPromotionListData                  = smsclient.FlashPromotionListData
	FlashPromotionListReq                   = smsclient.FlashPromotionListReq
	FlashPromotionListResp                  = smsclient.FlashPromotionListResp
	FlashPromotionLogAddReq                 = smsclient.FlashPromotionLogAddReq
	FlashPromotionLogAddResp                = smsclient.FlashPromotionLogAddResp
	FlashPromotionLogDeleteReq              = smsclient.FlashPromotionLogDeleteReq
	FlashPromotionLogDeleteResp             = smsclient.FlashPromotionLogDeleteResp
	FlashPromotionLogListData               = smsclient.FlashPromotionLogListData
	FlashPromotionLogListReq                = smsclient.FlashPromotionLogListReq
	FlashPromotionLogListResp               = smsclient.FlashPromotionLogListResp
	FlashPromotionLogUpdateReq              = smsclient.FlashPromotionLogUpdateReq
	FlashPromotionLogUpdateResp             = smsclient.FlashPromotionLogUpdateResp
	FlashPromotionProductRelationAddReq     = smsclient.FlashPromotionProductRelationAddReq
	FlashPromotionProductRelationAddResp    = smsclient.FlashPromotionProductRelationAddResp
	FlashPromotionProductRelationDeleteReq  = smsclient.FlashPromotionProductRelationDeleteReq
	FlashPromotionProductRelationDeleteResp = smsclient.FlashPromotionProductRelationDeleteResp
	FlashPromotionProductRelationListData   = smsclient.FlashPromotionProductRelationListData
	FlashPromotionProductRelationListReq    = smsclient.FlashPromotionProductRelationListReq
	FlashPromotionProductRelationListResp   = smsclient.FlashPromotionProductRelationListResp
	FlashPromotionProductRelationUpdateReq  = smsclient.FlashPromotionProductRelationUpdateReq
	FlashPromotionProductRelationUpdateResp = smsclient.FlashPromotionProductRelationUpdateResp
	FlashPromotionSessionAddReq             = smsclient.FlashPromotionSessionAddReq
	FlashPromotionSessionAddResp            = smsclient.FlashPromotionSessionAddResp
	FlashPromotionSessionByTimeReq          = smsclient.FlashPromotionSessionByTimeReq
	FlashPromotionSessionByTimeResp         = smsclient.FlashPromotionSessionByTimeResp
	FlashPromotionSessionDeleteReq          = smsclient.FlashPromotionSessionDeleteReq
	FlashPromotionSessionDeleteResp         = smsclient.FlashPromotionSessionDeleteResp
	FlashPromotionSessionListData           = smsclient.FlashPromotionSessionListData
	FlashPromotionSessionListReq            = smsclient.FlashPromotionSessionListReq
	FlashPromotionSessionListResp           = smsclient.FlashPromotionSessionListResp
	FlashPromotionSessionUpdateReq          = smsclient.FlashPromotionSessionUpdateReq
	FlashPromotionSessionUpdateResp         = smsclient.FlashPromotionSessionUpdateResp
	FlashPromotionUpdateReq                 = smsclient.FlashPromotionUpdateReq
	FlashPromotionUpdateResp                = smsclient.FlashPromotionUpdateResp
	HomeAdvertiseAddReq                     = smsclient.HomeAdvertiseAddReq
	HomeAdvertiseAddResp                    = smsclient.HomeAdvertiseAddResp
	HomeAdvertiseDeleteReq                  = smsclient.HomeAdvertiseDeleteReq
	HomeAdvertiseDeleteResp                 = smsclient.HomeAdvertiseDeleteResp
	HomeAdvertiseListData                   = smsclient.HomeAdvertiseListData
	HomeAdvertiseListReq                    = smsclient.HomeAdvertiseListReq
	HomeAdvertiseListResp                   = smsclient.HomeAdvertiseListResp
	HomeAdvertiseUpdateReq                  = smsclient.HomeAdvertiseUpdateReq
	HomeAdvertiseUpdateResp                 = smsclient.HomeAdvertiseUpdateResp
	HomeBrandAddData                        = smsclient.HomeBrandAddData
	HomeBrandAddReq                         = smsclient.HomeBrandAddReq
	HomeBrandAddResp                        = smsclient.HomeBrandAddResp
	HomeBrandDeleteReq                      = smsclient.HomeBrandDeleteReq
	HomeBrandDeleteResp                     = smsclient.HomeBrandDeleteResp
	HomeBrandListData                       = smsclient.HomeBrandListData
	HomeBrandListReq                        = smsclient.HomeBrandListReq
	HomeBrandListResp                       = smsclient.HomeBrandListResp
	HomeBrandUpdateReq                      = smsclient.HomeBrandUpdateReq
	HomeBrandUpdateResp                     = smsclient.HomeBrandUpdateResp
	HomeNewProductAddData                   = smsclient.HomeNewProductAddData
	HomeNewProductAddReq                    = smsclient.HomeNewProductAddReq
	HomeNewProductAddResp                   = smsclient.HomeNewProductAddResp
	HomeNewProductDeleteReq                 = smsclient.HomeNewProductDeleteReq
	HomeNewProductDeleteResp                = smsclient.HomeNewProductDeleteResp
	HomeNewProductListData                  = smsclient.HomeNewProductListData
	HomeNewProductListReq                   = smsclient.HomeNewProductListReq
	HomeNewProductListResp                  = smsclient.HomeNewProductListResp
	HomeNewProductUpdateReq                 = smsclient.HomeNewProductUpdateReq
	HomeNewProductUpdateResp                = smsclient.HomeNewProductUpdateResp
	HomeRecommendProductAddData             = smsclient.HomeRecommendProductAddData
	HomeRecommendProductAddReq              = smsclient.HomeRecommendProductAddReq
	HomeRecommendProductAddResp             = smsclient.HomeRecommendProductAddResp
	HomeRecommendProductDeleteReq           = smsclient.HomeRecommendProductDeleteReq
	HomeRecommendProductDeleteResp          = smsclient.HomeRecommendProductDeleteResp
	HomeRecommendProductListData            = smsclient.HomeRecommendProductListData
	HomeRecommendProductListReq             = smsclient.HomeRecommendProductListReq
	HomeRecommendProductListResp            = smsclient.HomeRecommendProductListResp
	HomeRecommendProductUpdateReq           = smsclient.HomeRecommendProductUpdateReq
	HomeRecommendProductUpdateResp          = smsclient.HomeRecommendProductUpdateResp
	HomeRecommendSubjectAddData             = smsclient.HomeRecommendSubjectAddData
	HomeRecommendSubjectAddReq              = smsclient.HomeRecommendSubjectAddReq
	HomeRecommendSubjectAddResp             = smsclient.HomeRecommendSubjectAddResp
	HomeRecommendSubjectDeleteReq           = smsclient.HomeRecommendSubjectDeleteReq
	HomeRecommendSubjectDeleteResp          = smsclient.HomeRecommendSubjectDeleteResp
	HomeRecommendSubjectListData            = smsclient.HomeRecommendSubjectListData
	HomeRecommendSubjectListReq             = smsclient.HomeRecommendSubjectListReq
	HomeRecommendSubjectListResp            = smsclient.HomeRecommendSubjectListResp
	HomeRecommendSubjectUpdateReq           = smsclient.HomeRecommendSubjectUpdateReq
	HomeRecommendSubjectUpdateResp          = smsclient.HomeRecommendSubjectUpdateResp

	FlashPromotionService interface {
		FlashPromotionAdd(ctx context.Context, in *FlashPromotionAddReq, opts ...grpc.CallOption) (*FlashPromotionAddResp, error)
		FlashPromotionList(ctx context.Context, in *FlashPromotionListReq, opts ...grpc.CallOption) (*FlashPromotionListResp, error)
		FlashPromotionUpdate(ctx context.Context, in *FlashPromotionUpdateReq, opts ...grpc.CallOption) (*FlashPromotionUpdateResp, error)
		FlashPromotionDelete(ctx context.Context, in *FlashPromotionDeleteReq, opts ...grpc.CallOption) (*FlashPromotionDeleteResp, error)
		FlashPromotionListByDate(ctx context.Context, in *FlashPromotionListByDateReq, opts ...grpc.CallOption) (*FlashPromotionListByDateResp, error)
	}

	defaultFlashPromotionService struct {
		cli zrpc.Client
	}
)

func NewFlashPromotionService(cli zrpc.Client) FlashPromotionService {
	return &defaultFlashPromotionService{
		cli: cli,
	}
}

func (m *defaultFlashPromotionService) FlashPromotionAdd(ctx context.Context, in *FlashPromotionAddReq, opts ...grpc.CallOption) (*FlashPromotionAddResp, error) {
	client := smsclient.NewFlashPromotionServiceClient(m.cli.Conn())
	return client.FlashPromotionAdd(ctx, in, opts...)
}

func (m *defaultFlashPromotionService) FlashPromotionList(ctx context.Context, in *FlashPromotionListReq, opts ...grpc.CallOption) (*FlashPromotionListResp, error) {
	client := smsclient.NewFlashPromotionServiceClient(m.cli.Conn())
	return client.FlashPromotionList(ctx, in, opts...)
}

func (m *defaultFlashPromotionService) FlashPromotionUpdate(ctx context.Context, in *FlashPromotionUpdateReq, opts ...grpc.CallOption) (*FlashPromotionUpdateResp, error) {
	client := smsclient.NewFlashPromotionServiceClient(m.cli.Conn())
	return client.FlashPromotionUpdate(ctx, in, opts...)
}

func (m *defaultFlashPromotionService) FlashPromotionDelete(ctx context.Context, in *FlashPromotionDeleteReq, opts ...grpc.CallOption) (*FlashPromotionDeleteResp, error) {
	client := smsclient.NewFlashPromotionServiceClient(m.cli.Conn())
	return client.FlashPromotionDelete(ctx, in, opts...)
}

func (m *defaultFlashPromotionService) FlashPromotionListByDate(ctx context.Context, in *FlashPromotionListByDateReq, opts ...grpc.CallOption) (*FlashPromotionListByDateResp, error) {
	client := smsclient.NewFlashPromotionServiceClient(m.cli.Conn())
	return client.FlashPromotionListByDate(ctx, in, opts...)
}