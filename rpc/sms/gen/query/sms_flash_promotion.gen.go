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

func newSmsFlashPromotion(db *gorm.DB, opts ...gen.DOOption) smsFlashPromotion {
	_smsFlashPromotion := smsFlashPromotion{}

	_smsFlashPromotion.smsFlashPromotionDo.UseDB(db, opts...)
	_smsFlashPromotion.smsFlashPromotionDo.UseModel(&model.SmsFlashPromotion{})

	tableName := _smsFlashPromotion.smsFlashPromotionDo.TableName()
	_smsFlashPromotion.ALL = field.NewAsterisk(tableName)
	_smsFlashPromotion.ID = field.NewInt64(tableName, "id")
	_smsFlashPromotion.Title = field.NewString(tableName, "title")
	_smsFlashPromotion.StartDate = field.NewTime(tableName, "start_date")
	_smsFlashPromotion.EndDate = field.NewTime(tableName, "end_date")
	_smsFlashPromotion.Status = field.NewInt32(tableName, "status")
	_smsFlashPromotion.CreateTime = field.NewTime(tableName, "create_time")

	_smsFlashPromotion.fillFieldMap()

	return _smsFlashPromotion
}

// smsFlashPromotion 限时购表
type smsFlashPromotion struct {
	smsFlashPromotionDo smsFlashPromotionDo

	ALL        field.Asterisk
	ID         field.Int64  // 编号
	Title      field.String // 标题
	StartDate  field.Time   // 开始日期
	EndDate    field.Time   // 结束日期
	Status     field.Int32  // 上下线状态
	CreateTime field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (s smsFlashPromotion) Table(newTableName string) *smsFlashPromotion {
	s.smsFlashPromotionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsFlashPromotion) As(alias string) *smsFlashPromotion {
	s.smsFlashPromotionDo.DO = *(s.smsFlashPromotionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsFlashPromotion) updateTableName(table string) *smsFlashPromotion {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Title = field.NewString(table, "title")
	s.StartDate = field.NewTime(table, "start_date")
	s.EndDate = field.NewTime(table, "end_date")
	s.Status = field.NewInt32(table, "status")
	s.CreateTime = field.NewTime(table, "create_time")

	s.fillFieldMap()

	return s
}

func (s *smsFlashPromotion) WithContext(ctx context.Context) ISmsFlashPromotionDo {
	return s.smsFlashPromotionDo.WithContext(ctx)
}

func (s smsFlashPromotion) TableName() string { return s.smsFlashPromotionDo.TableName() }

func (s smsFlashPromotion) Alias() string { return s.smsFlashPromotionDo.Alias() }

func (s smsFlashPromotion) Columns(cols ...field.Expr) gen.Columns {
	return s.smsFlashPromotionDo.Columns(cols...)
}

func (s *smsFlashPromotion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsFlashPromotion) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["title"] = s.Title
	s.fieldMap["start_date"] = s.StartDate
	s.fieldMap["end_date"] = s.EndDate
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_time"] = s.CreateTime
}

func (s smsFlashPromotion) clone(db *gorm.DB) smsFlashPromotion {
	s.smsFlashPromotionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsFlashPromotion) replaceDB(db *gorm.DB) smsFlashPromotion {
	s.smsFlashPromotionDo.ReplaceDB(db)
	return s
}

type smsFlashPromotionDo struct{ gen.DO }

type ISmsFlashPromotionDo interface {
	gen.SubQuery
	Debug() ISmsFlashPromotionDo
	WithContext(ctx context.Context) ISmsFlashPromotionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISmsFlashPromotionDo
	WriteDB() ISmsFlashPromotionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISmsFlashPromotionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISmsFlashPromotionDo
	Not(conds ...gen.Condition) ISmsFlashPromotionDo
	Or(conds ...gen.Condition) ISmsFlashPromotionDo
	Select(conds ...field.Expr) ISmsFlashPromotionDo
	Where(conds ...gen.Condition) ISmsFlashPromotionDo
	Order(conds ...field.Expr) ISmsFlashPromotionDo
	Distinct(cols ...field.Expr) ISmsFlashPromotionDo
	Omit(cols ...field.Expr) ISmsFlashPromotionDo
	Join(table schema.Tabler, on ...field.Expr) ISmsFlashPromotionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISmsFlashPromotionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISmsFlashPromotionDo
	Group(cols ...field.Expr) ISmsFlashPromotionDo
	Having(conds ...gen.Condition) ISmsFlashPromotionDo
	Limit(limit int) ISmsFlashPromotionDo
	Offset(offset int) ISmsFlashPromotionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsFlashPromotionDo
	Unscoped() ISmsFlashPromotionDo
	Create(values ...*model.SmsFlashPromotion) error
	CreateInBatches(values []*model.SmsFlashPromotion, batchSize int) error
	Save(values ...*model.SmsFlashPromotion) error
	First() (*model.SmsFlashPromotion, error)
	Take() (*model.SmsFlashPromotion, error)
	Last() (*model.SmsFlashPromotion, error)
	Find() ([]*model.SmsFlashPromotion, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsFlashPromotion, err error)
	FindInBatches(result *[]*model.SmsFlashPromotion, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SmsFlashPromotion) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISmsFlashPromotionDo
	Assign(attrs ...field.AssignExpr) ISmsFlashPromotionDo
	Joins(fields ...field.RelationField) ISmsFlashPromotionDo
	Preload(fields ...field.RelationField) ISmsFlashPromotionDo
	FirstOrInit() (*model.SmsFlashPromotion, error)
	FirstOrCreate() (*model.SmsFlashPromotion, error)
	FindByPage(offset int, limit int) (result []*model.SmsFlashPromotion, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISmsFlashPromotionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s smsFlashPromotionDo) Debug() ISmsFlashPromotionDo {
	return s.withDO(s.DO.Debug())
}

func (s smsFlashPromotionDo) WithContext(ctx context.Context) ISmsFlashPromotionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsFlashPromotionDo) ReadDB() ISmsFlashPromotionDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsFlashPromotionDo) WriteDB() ISmsFlashPromotionDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsFlashPromotionDo) Session(config *gorm.Session) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsFlashPromotionDo) Clauses(conds ...clause.Expression) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsFlashPromotionDo) Returning(value interface{}, columns ...string) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsFlashPromotionDo) Not(conds ...gen.Condition) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsFlashPromotionDo) Or(conds ...gen.Condition) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsFlashPromotionDo) Select(conds ...field.Expr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsFlashPromotionDo) Where(conds ...gen.Condition) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsFlashPromotionDo) Order(conds ...field.Expr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsFlashPromotionDo) Distinct(cols ...field.Expr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsFlashPromotionDo) Omit(cols ...field.Expr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsFlashPromotionDo) Join(table schema.Tabler, on ...field.Expr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsFlashPromotionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsFlashPromotionDo) RightJoin(table schema.Tabler, on ...field.Expr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsFlashPromotionDo) Group(cols ...field.Expr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsFlashPromotionDo) Having(conds ...gen.Condition) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsFlashPromotionDo) Limit(limit int) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsFlashPromotionDo) Offset(offset int) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsFlashPromotionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsFlashPromotionDo) Unscoped() ISmsFlashPromotionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsFlashPromotionDo) Create(values ...*model.SmsFlashPromotion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsFlashPromotionDo) CreateInBatches(values []*model.SmsFlashPromotion, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsFlashPromotionDo) Save(values ...*model.SmsFlashPromotion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsFlashPromotionDo) First() (*model.SmsFlashPromotion, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotion), nil
	}
}

func (s smsFlashPromotionDo) Take() (*model.SmsFlashPromotion, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotion), nil
	}
}

func (s smsFlashPromotionDo) Last() (*model.SmsFlashPromotion, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotion), nil
	}
}

func (s smsFlashPromotionDo) Find() ([]*model.SmsFlashPromotion, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsFlashPromotion), err
}

func (s smsFlashPromotionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsFlashPromotion, err error) {
	buf := make([]*model.SmsFlashPromotion, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsFlashPromotionDo) FindInBatches(result *[]*model.SmsFlashPromotion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsFlashPromotionDo) Attrs(attrs ...field.AssignExpr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsFlashPromotionDo) Assign(attrs ...field.AssignExpr) ISmsFlashPromotionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsFlashPromotionDo) Joins(fields ...field.RelationField) ISmsFlashPromotionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsFlashPromotionDo) Preload(fields ...field.RelationField) ISmsFlashPromotionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsFlashPromotionDo) FirstOrInit() (*model.SmsFlashPromotion, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotion), nil
	}
}

func (s smsFlashPromotionDo) FirstOrCreate() (*model.SmsFlashPromotion, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotion), nil
	}
}

func (s smsFlashPromotionDo) FindByPage(offset int, limit int) (result []*model.SmsFlashPromotion, count int64, err error) {
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

func (s smsFlashPromotionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsFlashPromotionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsFlashPromotionDo) Delete(models ...*model.SmsFlashPromotion) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsFlashPromotionDo) withDO(do gen.Dao) *smsFlashPromotionDo {
	s.DO = *do.(*gen.DO)
	return s
}
