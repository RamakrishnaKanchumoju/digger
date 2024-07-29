// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/next/model"
)

func newProjectComment(db *gorm.DB, opts ...gen.DOOption) projectComment {
	_projectComment := projectComment{}

	_projectComment.projectCommentDo.UseDB(db, opts...)
	_projectComment.projectCommentDo.UseModel(&model.ProjectComment{})

	tableName := _projectComment.projectCommentDo.TableName()
	_projectComment.ALL = field.NewAsterisk(tableName)
	_projectComment.ID = field.NewInt64(tableName, "id")
	_projectComment.CreatedAt = field.NewTime(tableName, "created_at")
	_projectComment.Text = field.NewString(tableName, "text")
	_projectComment.UserID = field.NewString(tableName, "user_id")
	_projectComment.InReplyTo = field.NewInt64(tableName, "in_reply_to")
	_projectComment.ProjectID = field.NewString(tableName, "project_id")

	_projectComment.fillFieldMap()

	return _projectComment
}

type projectComment struct {
	projectCommentDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	Text      field.String
	UserID    field.String
	InReplyTo field.Int64
	ProjectID field.String

	fieldMap map[string]field.Expr
}

func (p projectComment) Table(newTableName string) *projectComment {
	p.projectCommentDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectComment) As(alias string) *projectComment {
	p.projectCommentDo.DO = *(p.projectCommentDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectComment) updateTableName(table string) *projectComment {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.Text = field.NewString(table, "text")
	p.UserID = field.NewString(table, "user_id")
	p.InReplyTo = field.NewInt64(table, "in_reply_to")
	p.ProjectID = field.NewString(table, "project_id")

	p.fillFieldMap()

	return p
}

func (p *projectComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectComment) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["text"] = p.Text
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["in_reply_to"] = p.InReplyTo
	p.fieldMap["project_id"] = p.ProjectID
}

func (p projectComment) clone(db *gorm.DB) projectComment {
	p.projectCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectComment) replaceDB(db *gorm.DB) projectComment {
	p.projectCommentDo.ReplaceDB(db)
	return p
}

type projectCommentDo struct{ gen.DO }

type IProjectCommentDo interface {
	gen.SubQuery
	Debug() IProjectCommentDo
	WithContext(ctx context.Context) IProjectCommentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProjectCommentDo
	WriteDB() IProjectCommentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProjectCommentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProjectCommentDo
	Not(conds ...gen.Condition) IProjectCommentDo
	Or(conds ...gen.Condition) IProjectCommentDo
	Select(conds ...field.Expr) IProjectCommentDo
	Where(conds ...gen.Condition) IProjectCommentDo
	Order(conds ...field.Expr) IProjectCommentDo
	Distinct(cols ...field.Expr) IProjectCommentDo
	Omit(cols ...field.Expr) IProjectCommentDo
	Join(table schema.Tabler, on ...field.Expr) IProjectCommentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProjectCommentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProjectCommentDo
	Group(cols ...field.Expr) IProjectCommentDo
	Having(conds ...gen.Condition) IProjectCommentDo
	Limit(limit int) IProjectCommentDo
	Offset(offset int) IProjectCommentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectCommentDo
	Unscoped() IProjectCommentDo
	Create(values ...*model.ProjectComment) error
	CreateInBatches(values []*model.ProjectComment, batchSize int) error
	Save(values ...*model.ProjectComment) error
	First() (*model.ProjectComment, error)
	Take() (*model.ProjectComment, error)
	Last() (*model.ProjectComment, error)
	Find() ([]*model.ProjectComment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectComment, err error)
	FindInBatches(result *[]*model.ProjectComment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProjectComment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProjectCommentDo
	Assign(attrs ...field.AssignExpr) IProjectCommentDo
	Joins(fields ...field.RelationField) IProjectCommentDo
	Preload(fields ...field.RelationField) IProjectCommentDo
	FirstOrInit() (*model.ProjectComment, error)
	FirstOrCreate() (*model.ProjectComment, error)
	FindByPage(offset int, limit int) (result []*model.ProjectComment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProjectCommentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p projectCommentDo) Debug() IProjectCommentDo {
	return p.withDO(p.DO.Debug())
}

func (p projectCommentDo) WithContext(ctx context.Context) IProjectCommentDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectCommentDo) ReadDB() IProjectCommentDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectCommentDo) WriteDB() IProjectCommentDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectCommentDo) Session(config *gorm.Session) IProjectCommentDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectCommentDo) Clauses(conds ...clause.Expression) IProjectCommentDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectCommentDo) Returning(value interface{}, columns ...string) IProjectCommentDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectCommentDo) Not(conds ...gen.Condition) IProjectCommentDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectCommentDo) Or(conds ...gen.Condition) IProjectCommentDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectCommentDo) Select(conds ...field.Expr) IProjectCommentDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectCommentDo) Where(conds ...gen.Condition) IProjectCommentDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectCommentDo) Order(conds ...field.Expr) IProjectCommentDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectCommentDo) Distinct(cols ...field.Expr) IProjectCommentDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectCommentDo) Omit(cols ...field.Expr) IProjectCommentDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectCommentDo) Join(table schema.Tabler, on ...field.Expr) IProjectCommentDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProjectCommentDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) IProjectCommentDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectCommentDo) Group(cols ...field.Expr) IProjectCommentDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectCommentDo) Having(conds ...gen.Condition) IProjectCommentDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectCommentDo) Limit(limit int) IProjectCommentDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectCommentDo) Offset(offset int) IProjectCommentDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectCommentDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectCommentDo) Unscoped() IProjectCommentDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectCommentDo) Create(values ...*model.ProjectComment) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectCommentDo) CreateInBatches(values []*model.ProjectComment, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectCommentDo) Save(values ...*model.ProjectComment) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectCommentDo) First() (*model.ProjectComment, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectComment), nil
	}
}

func (p projectCommentDo) Take() (*model.ProjectComment, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectComment), nil
	}
}

func (p projectCommentDo) Last() (*model.ProjectComment, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectComment), nil
	}
}

func (p projectCommentDo) Find() ([]*model.ProjectComment, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProjectComment), err
}

func (p projectCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectComment, err error) {
	buf := make([]*model.ProjectComment, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectCommentDo) FindInBatches(result *[]*model.ProjectComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectCommentDo) Attrs(attrs ...field.AssignExpr) IProjectCommentDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectCommentDo) Assign(attrs ...field.AssignExpr) IProjectCommentDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectCommentDo) Joins(fields ...field.RelationField) IProjectCommentDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectCommentDo) Preload(fields ...field.RelationField) IProjectCommentDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectCommentDo) FirstOrInit() (*model.ProjectComment, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectComment), nil
	}
}

func (p projectCommentDo) FirstOrCreate() (*model.ProjectComment, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectComment), nil
	}
}

func (p projectCommentDo) FindByPage(offset int, limit int) (result []*model.ProjectComment, count int64, err error) {
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

func (p projectCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectCommentDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectCommentDo) Delete(models ...*model.ProjectComment) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectCommentDo) withDO(do gen.Dao) *projectCommentDo {
	p.DO = *do.(*gen.DO)
	return p
}