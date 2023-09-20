// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"foods-spider-server/internal/model/model/models"
)

func newYkYycfb(db *gorm.DB, opts ...gen.DOOption) ykYycfb {
	_ykYycfb := ykYycfb{}

	_ykYycfb.ykYycfbDo.UseDB(db, opts...)
	_ykYycfb.ykYycfbDo.UseModel(&models.YkYycfb{})

	tableName := _ykYycfb.ykYycfbDo.TableName()
	_ykYycfb.ALL = field.NewAsterisk(tableName)
	_ykYycfb.ID = field.NewInt32(tableName, "id")
	_ykYycfb.Name = field.NewString(tableName, "name")
	_ykYycfb.Reliang = field.NewInt32(tableName, "reliang")
	_ykYycfb.Danbai = field.NewInt32(tableName, "danbai")
	_ykYycfb.Zhifang = field.NewInt32(tableName, "zhifang")
	_ykYycfb.Tanshui = field.NewInt32(tableName, "tanshui")
	_ykYycfb.Xianwei = field.NewInt32(tableName, "xianwei")
	_ykYycfb.Va = field.NewInt32(tableName, "va")
	_ykYycfb.Huluobosu = field.NewInt32(tableName, "huluobosu")
	_ykYycfb.Shihuangchun = field.NewInt32(tableName, "shihuangchun")
	_ykYycfb.Liuansu = field.NewInt32(tableName, "liuansu")
	_ykYycfb.Hehuangsu = field.NewInt32(tableName, "hehuangsu")
	_ykYycfb.Yansuan = field.NewInt32(tableName, "yansuan")
	_ykYycfb.Vc = field.NewInt32(tableName, "vc")
	_ykYycfb.Ve = field.NewInt32(tableName, "ve")
	_ykYycfb.Danguchun = field.NewInt32(tableName, "danguchun")
	_ykYycfb.Jia = field.NewInt32(tableName, "jia")
	_ykYycfb.Na = field.NewInt32(tableName, "na")
	_ykYycfb.Gai = field.NewInt32(tableName, "gai")
	_ykYycfb.Mei = field.NewInt32(tableName, "mei")
	_ykYycfb.Tie = field.NewInt32(tableName, "tie")
	_ykYycfb.Meng = field.NewInt32(tableName, "meng")
	_ykYycfb.Xin = field.NewInt32(tableName, "xin")
	_ykYycfb.Tong = field.NewInt32(tableName, "tong")
	_ykYycfb.Lin = field.NewInt32(tableName, "lin")
	_ykYycfb.Xi = field.NewInt32(tableName, "xi")

	_ykYycfb.fillFieldMap()

	return _ykYycfb
}

type ykYycfb struct {
	ykYycfbDo

	ALL          field.Asterisk
	ID           field.Int32
	Name         field.String
	Reliang      field.Int32
	Danbai       field.Int32
	Zhifang      field.Int32
	Tanshui      field.Int32
	Xianwei      field.Int32
	Va           field.Int32
	Huluobosu    field.Int32
	Shihuangchun field.Int32
	Liuansu      field.Int32
	Hehuangsu    field.Int32
	Yansuan      field.Int32
	Vc           field.Int32
	Ve           field.Int32
	Danguchun    field.Int32
	Jia          field.Int32
	Na           field.Int32
	Gai          field.Int32
	Mei          field.Int32
	Tie          field.Int32
	Meng         field.Int32
	Xin          field.Int32
	Tong         field.Int32
	Lin          field.Int32
	Xi           field.Int32

	fieldMap map[string]field.Expr
}

func (y ykYycfb) Table(newTableName string) *ykYycfb {
	y.ykYycfbDo.UseTable(newTableName)
	return y.updateTableName(newTableName)
}

func (y ykYycfb) As(alias string) *ykYycfb {
	y.ykYycfbDo.DO = *(y.ykYycfbDo.As(alias).(*gen.DO))
	return y.updateTableName(alias)
}

func (y *ykYycfb) updateTableName(table string) *ykYycfb {
	y.ALL = field.NewAsterisk(table)
	y.ID = field.NewInt32(table, "id")
	y.Name = field.NewString(table, "name")
	y.Reliang = field.NewInt32(table, "reliang")
	y.Danbai = field.NewInt32(table, "danbai")
	y.Zhifang = field.NewInt32(table, "zhifang")
	y.Tanshui = field.NewInt32(table, "tanshui")
	y.Xianwei = field.NewInt32(table, "xianwei")
	y.Va = field.NewInt32(table, "va")
	y.Huluobosu = field.NewInt32(table, "huluobosu")
	y.Shihuangchun = field.NewInt32(table, "shihuangchun")
	y.Liuansu = field.NewInt32(table, "liuansu")
	y.Hehuangsu = field.NewInt32(table, "hehuangsu")
	y.Yansuan = field.NewInt32(table, "yansuan")
	y.Vc = field.NewInt32(table, "vc")
	y.Ve = field.NewInt32(table, "ve")
	y.Danguchun = field.NewInt32(table, "danguchun")
	y.Jia = field.NewInt32(table, "jia")
	y.Na = field.NewInt32(table, "na")
	y.Gai = field.NewInt32(table, "gai")
	y.Mei = field.NewInt32(table, "mei")
	y.Tie = field.NewInt32(table, "tie")
	y.Meng = field.NewInt32(table, "meng")
	y.Xin = field.NewInt32(table, "xin")
	y.Tong = field.NewInt32(table, "tong")
	y.Lin = field.NewInt32(table, "lin")
	y.Xi = field.NewInt32(table, "xi")

	y.fillFieldMap()

	return y
}

func (y *ykYycfb) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := y.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (y *ykYycfb) fillFieldMap() {
	y.fieldMap = make(map[string]field.Expr, 26)
	y.fieldMap["id"] = y.ID
	y.fieldMap["name"] = y.Name
	y.fieldMap["reliang"] = y.Reliang
	y.fieldMap["danbai"] = y.Danbai
	y.fieldMap["zhifang"] = y.Zhifang
	y.fieldMap["tanshui"] = y.Tanshui
	y.fieldMap["xianwei"] = y.Xianwei
	y.fieldMap["va"] = y.Va
	y.fieldMap["huluobosu"] = y.Huluobosu
	y.fieldMap["shihuangchun"] = y.Shihuangchun
	y.fieldMap["liuansu"] = y.Liuansu
	y.fieldMap["hehuangsu"] = y.Hehuangsu
	y.fieldMap["yansuan"] = y.Yansuan
	y.fieldMap["vc"] = y.Vc
	y.fieldMap["ve"] = y.Ve
	y.fieldMap["danguchun"] = y.Danguchun
	y.fieldMap["jia"] = y.Jia
	y.fieldMap["na"] = y.Na
	y.fieldMap["gai"] = y.Gai
	y.fieldMap["mei"] = y.Mei
	y.fieldMap["tie"] = y.Tie
	y.fieldMap["meng"] = y.Meng
	y.fieldMap["xin"] = y.Xin
	y.fieldMap["tong"] = y.Tong
	y.fieldMap["lin"] = y.Lin
	y.fieldMap["xi"] = y.Xi
}

func (y ykYycfb) clone(db *gorm.DB) ykYycfb {
	y.ykYycfbDo.ReplaceConnPool(db.Statement.ConnPool)
	return y
}

func (y ykYycfb) replaceDB(db *gorm.DB) ykYycfb {
	y.ykYycfbDo.ReplaceDB(db)
	return y
}

type ykYycfbDo struct{ gen.DO }

func (y ykYycfbDo) Debug() *ykYycfbDo {
	return y.withDO(y.DO.Debug())
}

func (y ykYycfbDo) WithContext(ctx context.Context) *ykYycfbDo {
	return y.withDO(y.DO.WithContext(ctx))
}

func (y ykYycfbDo) ReadDB() *ykYycfbDo {
	return y.Clauses(dbresolver.Read)
}

func (y ykYycfbDo) WriteDB() *ykYycfbDo {
	return y.Clauses(dbresolver.Write)
}

func (y ykYycfbDo) Session(config *gorm.Session) *ykYycfbDo {
	return y.withDO(y.DO.Session(config))
}

func (y ykYycfbDo) Clauses(conds ...clause.Expression) *ykYycfbDo {
	return y.withDO(y.DO.Clauses(conds...))
}

func (y ykYycfbDo) Returning(value interface{}, columns ...string) *ykYycfbDo {
	return y.withDO(y.DO.Returning(value, columns...))
}

func (y ykYycfbDo) Not(conds ...gen.Condition) *ykYycfbDo {
	return y.withDO(y.DO.Not(conds...))
}

func (y ykYycfbDo) Or(conds ...gen.Condition) *ykYycfbDo {
	return y.withDO(y.DO.Or(conds...))
}

func (y ykYycfbDo) Select(conds ...field.Expr) *ykYycfbDo {
	return y.withDO(y.DO.Select(conds...))
}

func (y ykYycfbDo) Where(conds ...gen.Condition) *ykYycfbDo {
	return y.withDO(y.DO.Where(conds...))
}

func (y ykYycfbDo) Order(conds ...field.Expr) *ykYycfbDo {
	return y.withDO(y.DO.Order(conds...))
}

func (y ykYycfbDo) Distinct(cols ...field.Expr) *ykYycfbDo {
	return y.withDO(y.DO.Distinct(cols...))
}

func (y ykYycfbDo) Omit(cols ...field.Expr) *ykYycfbDo {
	return y.withDO(y.DO.Omit(cols...))
}

func (y ykYycfbDo) Join(table schema.Tabler, on ...field.Expr) *ykYycfbDo {
	return y.withDO(y.DO.Join(table, on...))
}

func (y ykYycfbDo) LeftJoin(table schema.Tabler, on ...field.Expr) *ykYycfbDo {
	return y.withDO(y.DO.LeftJoin(table, on...))
}

func (y ykYycfbDo) RightJoin(table schema.Tabler, on ...field.Expr) *ykYycfbDo {
	return y.withDO(y.DO.RightJoin(table, on...))
}

func (y ykYycfbDo) Group(cols ...field.Expr) *ykYycfbDo {
	return y.withDO(y.DO.Group(cols...))
}

func (y ykYycfbDo) Having(conds ...gen.Condition) *ykYycfbDo {
	return y.withDO(y.DO.Having(conds...))
}

func (y ykYycfbDo) Limit(limit int) *ykYycfbDo {
	return y.withDO(y.DO.Limit(limit))
}

func (y ykYycfbDo) Offset(offset int) *ykYycfbDo {
	return y.withDO(y.DO.Offset(offset))
}

func (y ykYycfbDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *ykYycfbDo {
	return y.withDO(y.DO.Scopes(funcs...))
}

func (y ykYycfbDo) Unscoped() *ykYycfbDo {
	return y.withDO(y.DO.Unscoped())
}

func (y ykYycfbDo) Create(values ...*models.YkYycfb) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Create(values)
}

func (y ykYycfbDo) CreateInBatches(values []*models.YkYycfb, batchSize int) error {
	return y.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (y ykYycfbDo) Save(values ...*models.YkYycfb) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Save(values)
}

func (y ykYycfbDo) First() (*models.YkYycfb, error) {
	if result, err := y.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkYycfb), nil
	}
}

func (y ykYycfbDo) Take() (*models.YkYycfb, error) {
	if result, err := y.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkYycfb), nil
	}
}

func (y ykYycfbDo) Last() (*models.YkYycfb, error) {
	if result, err := y.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkYycfb), nil
	}
}

func (y ykYycfbDo) Find() ([]*models.YkYycfb, error) {
	result, err := y.DO.Find()
	return result.([]*models.YkYycfb), err
}

func (y ykYycfbDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.YkYycfb, err error) {
	buf := make([]*models.YkYycfb, 0, batchSize)
	err = y.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (y ykYycfbDo) FindInBatches(result *[]*models.YkYycfb, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return y.DO.FindInBatches(result, batchSize, fc)
}

func (y ykYycfbDo) Attrs(attrs ...field.AssignExpr) *ykYycfbDo {
	return y.withDO(y.DO.Attrs(attrs...))
}

func (y ykYycfbDo) Assign(attrs ...field.AssignExpr) *ykYycfbDo {
	return y.withDO(y.DO.Assign(attrs...))
}

func (y ykYycfbDo) Joins(fields ...field.RelationField) *ykYycfbDo {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Joins(_f))
	}
	return &y
}

func (y ykYycfbDo) Preload(fields ...field.RelationField) *ykYycfbDo {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Preload(_f))
	}
	return &y
}

func (y ykYycfbDo) FirstOrInit() (*models.YkYycfb, error) {
	if result, err := y.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkYycfb), nil
	}
}

func (y ykYycfbDo) FirstOrCreate() (*models.YkYycfb, error) {
	if result, err := y.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkYycfb), nil
	}
}

func (y ykYycfbDo) FindByPage(offset int, limit int) (result []*models.YkYycfb, count int64, err error) {
	result, err = y.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = y.Offset(-1).Limit(-1).Count()
	return
}

func (y ykYycfbDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = y.Count()
	if err != nil {
		return
	}

	err = y.Offset(offset).Limit(limit).Scan(result)
	return
}

func (y ykYycfbDo) Scan(result interface{}) (err error) {
	return y.DO.Scan(result)
}

func (y ykYycfbDo) Delete(models ...*models.YkYycfb) (result gen.ResultInfo, err error) {
	return y.DO.Delete(models)
}

func (y *ykYycfbDo) withDO(do gen.Dao) *ykYycfbDo {
	y.DO = *do.(*gen.DO)
	return y
}
