// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                                = new(Query)
	SmsCoupon                        *smsCoupon
	SmsCouponHistory                 *smsCouponHistory
	SmsCouponProductCategoryRelation *smsCouponProductCategoryRelation
	SmsCouponProductRelation         *smsCouponProductRelation
	SmsFlashPromotion                *smsFlashPromotion
	SmsFlashPromotionLog             *smsFlashPromotionLog
	SmsFlashPromotionProductRelation *smsFlashPromotionProductRelation
	SmsFlashPromotionSession         *smsFlashPromotionSession
	SmsHomeAdvertise                 *smsHomeAdvertise
	SmsHomeBrand                     *smsHomeBrand
	SmsHomeNewProduct                *smsHomeNewProduct
	SmsHomeRecommendProduct          *smsHomeRecommendProduct
	SmsHomeRecommendSubject          *smsHomeRecommendSubject
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	SmsCoupon = &Q.SmsCoupon
	SmsCouponHistory = &Q.SmsCouponHistory
	SmsCouponProductCategoryRelation = &Q.SmsCouponProductCategoryRelation
	SmsCouponProductRelation = &Q.SmsCouponProductRelation
	SmsFlashPromotion = &Q.SmsFlashPromotion
	SmsFlashPromotionLog = &Q.SmsFlashPromotionLog
	SmsFlashPromotionProductRelation = &Q.SmsFlashPromotionProductRelation
	SmsFlashPromotionSession = &Q.SmsFlashPromotionSession
	SmsHomeAdvertise = &Q.SmsHomeAdvertise
	SmsHomeBrand = &Q.SmsHomeBrand
	SmsHomeNewProduct = &Q.SmsHomeNewProduct
	SmsHomeRecommendProduct = &Q.SmsHomeRecommendProduct
	SmsHomeRecommendSubject = &Q.SmsHomeRecommendSubject
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                               db,
		SmsCoupon:                        newSmsCoupon(db, opts...),
		SmsCouponHistory:                 newSmsCouponHistory(db, opts...),
		SmsCouponProductCategoryRelation: newSmsCouponProductCategoryRelation(db, opts...),
		SmsCouponProductRelation:         newSmsCouponProductRelation(db, opts...),
		SmsFlashPromotion:                newSmsFlashPromotion(db, opts...),
		SmsFlashPromotionLog:             newSmsFlashPromotionLog(db, opts...),
		SmsFlashPromotionProductRelation: newSmsFlashPromotionProductRelation(db, opts...),
		SmsFlashPromotionSession:         newSmsFlashPromotionSession(db, opts...),
		SmsHomeAdvertise:                 newSmsHomeAdvertise(db, opts...),
		SmsHomeBrand:                     newSmsHomeBrand(db, opts...),
		SmsHomeNewProduct:                newSmsHomeNewProduct(db, opts...),
		SmsHomeRecommendProduct:          newSmsHomeRecommendProduct(db, opts...),
		SmsHomeRecommendSubject:          newSmsHomeRecommendSubject(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	SmsCoupon                        smsCoupon
	SmsCouponHistory                 smsCouponHistory
	SmsCouponProductCategoryRelation smsCouponProductCategoryRelation
	SmsCouponProductRelation         smsCouponProductRelation
	SmsFlashPromotion                smsFlashPromotion
	SmsFlashPromotionLog             smsFlashPromotionLog
	SmsFlashPromotionProductRelation smsFlashPromotionProductRelation
	SmsFlashPromotionSession         smsFlashPromotionSession
	SmsHomeAdvertise                 smsHomeAdvertise
	SmsHomeBrand                     smsHomeBrand
	SmsHomeNewProduct                smsHomeNewProduct
	SmsHomeRecommendProduct          smsHomeRecommendProduct
	SmsHomeRecommendSubject          smsHomeRecommendSubject
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                               db,
		SmsCoupon:                        q.SmsCoupon.clone(db),
		SmsCouponHistory:                 q.SmsCouponHistory.clone(db),
		SmsCouponProductCategoryRelation: q.SmsCouponProductCategoryRelation.clone(db),
		SmsCouponProductRelation:         q.SmsCouponProductRelation.clone(db),
		SmsFlashPromotion:                q.SmsFlashPromotion.clone(db),
		SmsFlashPromotionLog:             q.SmsFlashPromotionLog.clone(db),
		SmsFlashPromotionProductRelation: q.SmsFlashPromotionProductRelation.clone(db),
		SmsFlashPromotionSession:         q.SmsFlashPromotionSession.clone(db),
		SmsHomeAdvertise:                 q.SmsHomeAdvertise.clone(db),
		SmsHomeBrand:                     q.SmsHomeBrand.clone(db),
		SmsHomeNewProduct:                q.SmsHomeNewProduct.clone(db),
		SmsHomeRecommendProduct:          q.SmsHomeRecommendProduct.clone(db),
		SmsHomeRecommendSubject:          q.SmsHomeRecommendSubject.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                               db,
		SmsCoupon:                        q.SmsCoupon.replaceDB(db),
		SmsCouponHistory:                 q.SmsCouponHistory.replaceDB(db),
		SmsCouponProductCategoryRelation: q.SmsCouponProductCategoryRelation.replaceDB(db),
		SmsCouponProductRelation:         q.SmsCouponProductRelation.replaceDB(db),
		SmsFlashPromotion:                q.SmsFlashPromotion.replaceDB(db),
		SmsFlashPromotionLog:             q.SmsFlashPromotionLog.replaceDB(db),
		SmsFlashPromotionProductRelation: q.SmsFlashPromotionProductRelation.replaceDB(db),
		SmsFlashPromotionSession:         q.SmsFlashPromotionSession.replaceDB(db),
		SmsHomeAdvertise:                 q.SmsHomeAdvertise.replaceDB(db),
		SmsHomeBrand:                     q.SmsHomeBrand.replaceDB(db),
		SmsHomeNewProduct:                q.SmsHomeNewProduct.replaceDB(db),
		SmsHomeRecommendProduct:          q.SmsHomeRecommendProduct.replaceDB(db),
		SmsHomeRecommendSubject:          q.SmsHomeRecommendSubject.replaceDB(db),
	}
}

type queryCtx struct {
	SmsCoupon                        ISmsCouponDo
	SmsCouponHistory                 ISmsCouponHistoryDo
	SmsCouponProductCategoryRelation ISmsCouponProductCategoryRelationDo
	SmsCouponProductRelation         ISmsCouponProductRelationDo
	SmsFlashPromotion                ISmsFlashPromotionDo
	SmsFlashPromotionLog             ISmsFlashPromotionLogDo
	SmsFlashPromotionProductRelation ISmsFlashPromotionProductRelationDo
	SmsFlashPromotionSession         ISmsFlashPromotionSessionDo
	SmsHomeAdvertise                 ISmsHomeAdvertiseDo
	SmsHomeBrand                     ISmsHomeBrandDo
	SmsHomeNewProduct                ISmsHomeNewProductDo
	SmsHomeRecommendProduct          ISmsHomeRecommendProductDo
	SmsHomeRecommendSubject          ISmsHomeRecommendSubjectDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		SmsCoupon:                        q.SmsCoupon.WithContext(ctx),
		SmsCouponHistory:                 q.SmsCouponHistory.WithContext(ctx),
		SmsCouponProductCategoryRelation: q.SmsCouponProductCategoryRelation.WithContext(ctx),
		SmsCouponProductRelation:         q.SmsCouponProductRelation.WithContext(ctx),
		SmsFlashPromotion:                q.SmsFlashPromotion.WithContext(ctx),
		SmsFlashPromotionLog:             q.SmsFlashPromotionLog.WithContext(ctx),
		SmsFlashPromotionProductRelation: q.SmsFlashPromotionProductRelation.WithContext(ctx),
		SmsFlashPromotionSession:         q.SmsFlashPromotionSession.WithContext(ctx),
		SmsHomeAdvertise:                 q.SmsHomeAdvertise.WithContext(ctx),
		SmsHomeBrand:                     q.SmsHomeBrand.WithContext(ctx),
		SmsHomeNewProduct:                q.SmsHomeNewProduct.WithContext(ctx),
		SmsHomeRecommendProduct:          q.SmsHomeRecommendProduct.WithContext(ctx),
		SmsHomeRecommendSubject:          q.SmsHomeRecommendSubject.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
