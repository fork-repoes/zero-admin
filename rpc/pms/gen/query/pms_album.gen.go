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

	"github.com/feihua/zero-admin/rpc/pms/gen/model"
)

func newPmsAlbum(db *gorm.DB, opts ...gen.DOOption) pmsAlbum {
	_pmsAlbum := pmsAlbum{}

	_pmsAlbum.pmsAlbumDo.UseDB(db, opts...)
	_pmsAlbum.pmsAlbumDo.UseModel(&model.PmsAlbum{})

	tableName := _pmsAlbum.pmsAlbumDo.TableName()
	_pmsAlbum.ALL = field.NewAsterisk(tableName)
	_pmsAlbum.ID = field.NewInt64(tableName, "id")
	_pmsAlbum.Name = field.NewString(tableName, "name")
	_pmsAlbum.CoverPic = field.NewString(tableName, "cover_pic")
	_pmsAlbum.PicCount = field.NewInt32(tableName, "pic_count")
	_pmsAlbum.Sort = field.NewInt32(tableName, "sort")
	_pmsAlbum.Description = field.NewString(tableName, "description")

	_pmsAlbum.fillFieldMap()

	return _pmsAlbum
}

// pmsAlbum 相册表
type pmsAlbum struct {
	pmsAlbumDo pmsAlbumDo

	ALL         field.Asterisk
	ID          field.Int64
	Name        field.String
	CoverPic    field.String
	PicCount    field.Int32
	Sort        field.Int32
	Description field.String

	fieldMap map[string]field.Expr
}

func (p pmsAlbum) Table(newTableName string) *pmsAlbum {
	p.pmsAlbumDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsAlbum) As(alias string) *pmsAlbum {
	p.pmsAlbumDo.DO = *(p.pmsAlbumDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsAlbum) updateTableName(table string) *pmsAlbum {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.Name = field.NewString(table, "name")
	p.CoverPic = field.NewString(table, "cover_pic")
	p.PicCount = field.NewInt32(table, "pic_count")
	p.Sort = field.NewInt32(table, "sort")
	p.Description = field.NewString(table, "description")

	p.fillFieldMap()

	return p
}

func (p *pmsAlbum) WithContext(ctx context.Context) IPmsAlbumDo { return p.pmsAlbumDo.WithContext(ctx) }

func (p pmsAlbum) TableName() string { return p.pmsAlbumDo.TableName() }

func (p pmsAlbum) Alias() string { return p.pmsAlbumDo.Alias() }

func (p pmsAlbum) Columns(cols ...field.Expr) gen.Columns { return p.pmsAlbumDo.Columns(cols...) }

func (p *pmsAlbum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsAlbum) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["cover_pic"] = p.CoverPic
	p.fieldMap["pic_count"] = p.PicCount
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["description"] = p.Description
}

func (p pmsAlbum) clone(db *gorm.DB) pmsAlbum {
	p.pmsAlbumDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsAlbum) replaceDB(db *gorm.DB) pmsAlbum {
	p.pmsAlbumDo.ReplaceDB(db)
	return p
}

type pmsAlbumDo struct{ gen.DO }

type IPmsAlbumDo interface {
	gen.SubQuery
	Debug() IPmsAlbumDo
	WithContext(ctx context.Context) IPmsAlbumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsAlbumDo
	WriteDB() IPmsAlbumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsAlbumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsAlbumDo
	Not(conds ...gen.Condition) IPmsAlbumDo
	Or(conds ...gen.Condition) IPmsAlbumDo
	Select(conds ...field.Expr) IPmsAlbumDo
	Where(conds ...gen.Condition) IPmsAlbumDo
	Order(conds ...field.Expr) IPmsAlbumDo
	Distinct(cols ...field.Expr) IPmsAlbumDo
	Omit(cols ...field.Expr) IPmsAlbumDo
	Join(table schema.Tabler, on ...field.Expr) IPmsAlbumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsAlbumDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsAlbumDo
	Group(cols ...field.Expr) IPmsAlbumDo
	Having(conds ...gen.Condition) IPmsAlbumDo
	Limit(limit int) IPmsAlbumDo
	Offset(offset int) IPmsAlbumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsAlbumDo
	Unscoped() IPmsAlbumDo
	Create(values ...*model.PmsAlbum) error
	CreateInBatches(values []*model.PmsAlbum, batchSize int) error
	Save(values ...*model.PmsAlbum) error
	First() (*model.PmsAlbum, error)
	Take() (*model.PmsAlbum, error)
	Last() (*model.PmsAlbum, error)
	Find() ([]*model.PmsAlbum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsAlbum, err error)
	FindInBatches(result *[]*model.PmsAlbum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsAlbum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsAlbumDo
	Assign(attrs ...field.AssignExpr) IPmsAlbumDo
	Joins(fields ...field.RelationField) IPmsAlbumDo
	Preload(fields ...field.RelationField) IPmsAlbumDo
	FirstOrInit() (*model.PmsAlbum, error)
	FirstOrCreate() (*model.PmsAlbum, error)
	FindByPage(offset int, limit int) (result []*model.PmsAlbum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsAlbumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsAlbumDo) Debug() IPmsAlbumDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsAlbumDo) WithContext(ctx context.Context) IPmsAlbumDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsAlbumDo) ReadDB() IPmsAlbumDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsAlbumDo) WriteDB() IPmsAlbumDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsAlbumDo) Session(config *gorm.Session) IPmsAlbumDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsAlbumDo) Clauses(conds ...clause.Expression) IPmsAlbumDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsAlbumDo) Returning(value interface{}, columns ...string) IPmsAlbumDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsAlbumDo) Not(conds ...gen.Condition) IPmsAlbumDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsAlbumDo) Or(conds ...gen.Condition) IPmsAlbumDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsAlbumDo) Select(conds ...field.Expr) IPmsAlbumDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsAlbumDo) Where(conds ...gen.Condition) IPmsAlbumDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsAlbumDo) Order(conds ...field.Expr) IPmsAlbumDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsAlbumDo) Distinct(cols ...field.Expr) IPmsAlbumDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsAlbumDo) Omit(cols ...field.Expr) IPmsAlbumDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsAlbumDo) Join(table schema.Tabler, on ...field.Expr) IPmsAlbumDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsAlbumDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsAlbumDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsAlbumDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsAlbumDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsAlbumDo) Group(cols ...field.Expr) IPmsAlbumDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsAlbumDo) Having(conds ...gen.Condition) IPmsAlbumDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsAlbumDo) Limit(limit int) IPmsAlbumDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsAlbumDo) Offset(offset int) IPmsAlbumDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsAlbumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsAlbumDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsAlbumDo) Unscoped() IPmsAlbumDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsAlbumDo) Create(values ...*model.PmsAlbum) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsAlbumDo) CreateInBatches(values []*model.PmsAlbum, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsAlbumDo) Save(values ...*model.PmsAlbum) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsAlbumDo) First() (*model.PmsAlbum, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbum), nil
	}
}

func (p pmsAlbumDo) Take() (*model.PmsAlbum, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbum), nil
	}
}

func (p pmsAlbumDo) Last() (*model.PmsAlbum, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbum), nil
	}
}

func (p pmsAlbumDo) Find() ([]*model.PmsAlbum, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsAlbum), err
}

func (p pmsAlbumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsAlbum, err error) {
	buf := make([]*model.PmsAlbum, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsAlbumDo) FindInBatches(result *[]*model.PmsAlbum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsAlbumDo) Attrs(attrs ...field.AssignExpr) IPmsAlbumDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsAlbumDo) Assign(attrs ...field.AssignExpr) IPmsAlbumDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsAlbumDo) Joins(fields ...field.RelationField) IPmsAlbumDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsAlbumDo) Preload(fields ...field.RelationField) IPmsAlbumDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsAlbumDo) FirstOrInit() (*model.PmsAlbum, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbum), nil
	}
}

func (p pmsAlbumDo) FirstOrCreate() (*model.PmsAlbum, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbum), nil
	}
}

func (p pmsAlbumDo) FindByPage(offset int, limit int) (result []*model.PmsAlbum, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pmsAlbumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsAlbumDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsAlbumDo) Delete(models ...*model.PmsAlbum) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsAlbumDo) withDO(do gen.Dao) *pmsAlbumDo {
	p.DO = *do.(*gen.DO)
	return p
}
