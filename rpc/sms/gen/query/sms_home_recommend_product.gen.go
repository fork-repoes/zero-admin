// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/feihua/zero-admin/rpc/sms/gen/model"
)

func newSmsHomeRecommendProduct(db *gorm.DB, opts ...gen.DOOption) smsHomeRecommendProduct {
	_smsHomeRecommendProduct := smsHomeRecommendProduct{}

	_smsHomeRecommendProduct.smsHomeRecommendProductDo.UseDB(db, opts...)
	_smsHomeRecommendProduct.smsHomeRecommendProductDo.UseModel(&model.SmsHomeRecommendProduct{})

	tableName := _smsHomeRecommendProduct.smsHomeRecommendProductDo.TableName()
	_smsHomeRecommendProduct.ALL = field.NewAsterisk(tableName)
	_smsHomeRecommendProduct.ID = field.NewInt64(tableName, "id")
	_smsHomeRecommendProduct.ProductID = field.NewInt64(tableName, "product_id")
	_smsHomeRecommendProduct.ProductName = field.NewString(tableName, "product_name")
	_smsHomeRecommendProduct.RecommendStatus = field.NewInt32(tableName, "recommend_status")
	_smsHomeRecommendProduct.Sort = field.NewInt32(tableName, "sort")

	_smsHomeRecommendProduct.fillFieldMap()

	return _smsHomeRecommendProduct
}

// smsHomeRecommendProduct 人气推荐商品表
type smsHomeRecommendProduct struct {
	smsHomeRecommendProductDo smsHomeRecommendProductDo

	ALL             field.Asterisk
	ID              field.Int64
	ProductID       field.Int64  // 商品id
	ProductName     field.String // 商品名称
	RecommendStatus field.Int32  // 推荐状态：0->不推荐;1->推荐
	Sort            field.Int32  // 排序

	fieldMap map[string]field.Expr
}

func (s smsHomeRecommendProduct) Table(newTableName string) *smsHomeRecommendProduct {
	s.smsHomeRecommendProductDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsHomeRecommendProduct) As(alias string) *smsHomeRecommendProduct {
	s.smsHomeRecommendProductDo.DO = *(s.smsHomeRecommendProductDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsHomeRecommendProduct) updateTableName(table string) *smsHomeRecommendProduct {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.ProductID = field.NewInt64(table, "product_id")
	s.ProductName = field.NewString(table, "product_name")
	s.RecommendStatus = field.NewInt32(table, "recommend_status")
	s.Sort = field.NewInt32(table, "sort")

	s.fillFieldMap()

	return s
}

func (s *smsHomeRecommendProduct) WithContext(ctx context.Context) ISmsHomeRecommendProductDo {
	return s.smsHomeRecommendProductDo.WithContext(ctx)
}

func (s smsHomeRecommendProduct) TableName() string { return s.smsHomeRecommendProductDo.TableName() }

func (s smsHomeRecommendProduct) Alias() string { return s.smsHomeRecommendProductDo.Alias() }

func (s smsHomeRecommendProduct) Columns(cols ...field.Expr) gen.Columns {
	return s.smsHomeRecommendProductDo.Columns(cols...)
}

func (s *smsHomeRecommendProduct) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsHomeRecommendProduct) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id"] = s.ID
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["product_name"] = s.ProductName
	s.fieldMap["recommend_status"] = s.RecommendStatus
	s.fieldMap["sort"] = s.Sort
}

func (s smsHomeRecommendProduct) clone(db *gorm.DB) smsHomeRecommendProduct {
	s.smsHomeRecommendProductDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsHomeRecommendProduct) replaceDB(db *gorm.DB) smsHomeRecommendProduct {
	s.smsHomeRecommendProductDo.ReplaceDB(db)
	return s
}

type smsHomeRecommendProductDo struct{ gen.DO }

type ISmsHomeRecommendProductDo interface {
	gen.SubQuery
	Debug() ISmsHomeRecommendProductDo
	WithContext(ctx context.Context) ISmsHomeRecommendProductDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISmsHomeRecommendProductDo
	WriteDB() ISmsHomeRecommendProductDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISmsHomeRecommendProductDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISmsHomeRecommendProductDo
	Not(conds ...gen.Condition) ISmsHomeRecommendProductDo
	Or(conds ...gen.Condition) ISmsHomeRecommendProductDo
	Select(conds ...field.Expr) ISmsHomeRecommendProductDo
	Where(conds ...gen.Condition) ISmsHomeRecommendProductDo
	Order(conds ...field.Expr) ISmsHomeRecommendProductDo
	Distinct(cols ...field.Expr) ISmsHomeRecommendProductDo
	Omit(cols ...field.Expr) ISmsHomeRecommendProductDo
	Join(table schema.Tabler, on ...field.Expr) ISmsHomeRecommendProductDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISmsHomeRecommendProductDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISmsHomeRecommendProductDo
	Group(cols ...field.Expr) ISmsHomeRecommendProductDo
	Having(conds ...gen.Condition) ISmsHomeRecommendProductDo
	Limit(limit int) ISmsHomeRecommendProductDo
	Offset(offset int) ISmsHomeRecommendProductDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsHomeRecommendProductDo
	Unscoped() ISmsHomeRecommendProductDo
	Create(values ...*model.SmsHomeRecommendProduct) error
	CreateInBatches(values []*model.SmsHomeRecommendProduct, batchSize int) error
	Save(values ...*model.SmsHomeRecommendProduct) error
	First() (*model.SmsHomeRecommendProduct, error)
	Take() (*model.SmsHomeRecommendProduct, error)
	Last() (*model.SmsHomeRecommendProduct, error)
	Find() ([]*model.SmsHomeRecommendProduct, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsHomeRecommendProduct, err error)
	FindInBatches(result *[]*model.SmsHomeRecommendProduct, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SmsHomeRecommendProduct) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISmsHomeRecommendProductDo
	Assign(attrs ...field.AssignExpr) ISmsHomeRecommendProductDo
	Joins(fields ...field.RelationField) ISmsHomeRecommendProductDo
	Preload(fields ...field.RelationField) ISmsHomeRecommendProductDo
	FirstOrInit() (*model.SmsHomeRecommendProduct, error)
	FirstOrCreate() (*model.SmsHomeRecommendProduct, error)
	FindByPage(offset int, limit int) (result []*model.SmsHomeRecommendProduct, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISmsHomeRecommendProductDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s smsHomeRecommendProductDo) Debug() ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Debug())
}

func (s smsHomeRecommendProductDo) WithContext(ctx context.Context) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsHomeRecommendProductDo) ReadDB() ISmsHomeRecommendProductDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsHomeRecommendProductDo) WriteDB() ISmsHomeRecommendProductDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsHomeRecommendProductDo) Session(config *gorm.Session) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsHomeRecommendProductDo) Clauses(conds ...clause.Expression) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsHomeRecommendProductDo) Returning(value interface{}, columns ...string) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsHomeRecommendProductDo) Not(conds ...gen.Condition) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsHomeRecommendProductDo) Or(conds ...gen.Condition) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsHomeRecommendProductDo) Select(conds ...field.Expr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsHomeRecommendProductDo) Where(conds ...gen.Condition) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsHomeRecommendProductDo) Order(conds ...field.Expr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsHomeRecommendProductDo) Distinct(cols ...field.Expr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsHomeRecommendProductDo) Omit(cols ...field.Expr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsHomeRecommendProductDo) Join(table schema.Tabler, on ...field.Expr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsHomeRecommendProductDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsHomeRecommendProductDo) RightJoin(table schema.Tabler, on ...field.Expr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsHomeRecommendProductDo) Group(cols ...field.Expr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsHomeRecommendProductDo) Having(conds ...gen.Condition) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsHomeRecommendProductDo) Limit(limit int) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsHomeRecommendProductDo) Offset(offset int) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsHomeRecommendProductDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsHomeRecommendProductDo) Unscoped() ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsHomeRecommendProductDo) Create(values ...*model.SmsHomeRecommendProduct) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsHomeRecommendProductDo) CreateInBatches(values []*model.SmsHomeRecommendProduct, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsHomeRecommendProductDo) Save(values ...*model.SmsHomeRecommendProduct) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsHomeRecommendProductDo) First() (*model.SmsHomeRecommendProduct, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeRecommendProduct), nil
	}
}

func (s smsHomeRecommendProductDo) Take() (*model.SmsHomeRecommendProduct, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeRecommendProduct), nil
	}
}

func (s smsHomeRecommendProductDo) Last() (*model.SmsHomeRecommendProduct, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeRecommendProduct), nil
	}
}

func (s smsHomeRecommendProductDo) Find() ([]*model.SmsHomeRecommendProduct, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsHomeRecommendProduct), err
}

func (s smsHomeRecommendProductDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsHomeRecommendProduct, err error) {
	buf := make([]*model.SmsHomeRecommendProduct, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsHomeRecommendProductDo) FindInBatches(result *[]*model.SmsHomeRecommendProduct, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsHomeRecommendProductDo) Attrs(attrs ...field.AssignExpr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsHomeRecommendProductDo) Assign(attrs ...field.AssignExpr) ISmsHomeRecommendProductDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsHomeRecommendProductDo) Joins(fields ...field.RelationField) ISmsHomeRecommendProductDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsHomeRecommendProductDo) Preload(fields ...field.RelationField) ISmsHomeRecommendProductDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsHomeRecommendProductDo) FirstOrInit() (*model.SmsHomeRecommendProduct, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeRecommendProduct), nil
	}
}

func (s smsHomeRecommendProductDo) FirstOrCreate() (*model.SmsHomeRecommendProduct, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeRecommendProduct), nil
	}
}

func (s smsHomeRecommendProductDo) FindByPage(offset int, limit int) (result []*model.SmsHomeRecommendProduct, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s smsHomeRecommendProductDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsHomeRecommendProductDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsHomeRecommendProductDo) Delete(models ...*model.SmsHomeRecommendProduct) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsHomeRecommendProductDo) withDO(do gen.Dao) *smsHomeRecommendProductDo {
	s.DO = *do.(*gen.DO)
	return s
}
