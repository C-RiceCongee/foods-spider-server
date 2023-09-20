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

func newYkFoodnutrition(db *gorm.DB, opts ...gen.DOOption) ykFoodnutrition {
	_ykFoodnutrition := ykFoodnutrition{}

	_ykFoodnutrition.ykFoodnutritionDo.UseDB(db, opts...)
	_ykFoodnutrition.ykFoodnutritionDo.UseModel(&models.YkFoodnutrition{})

	tableName := _ykFoodnutrition.ykFoodnutritionDo.TableName()
	_ykFoodnutrition.ALL = field.NewAsterisk(tableName)
	_ykFoodnutrition.ID = field.NewInt32(tableName, "id")
	_ykFoodnutrition.Docv = field.NewString(tableName, "docv")
	_ykFoodnutrition.Leixing = field.NewInt32(tableName, "leixing")
	_ykFoodnutrition.Name = field.NewString(tableName, "name")
	_ykFoodnutrition.Xuhao = field.NewString(tableName, "xuhao")
	_ykFoodnutrition.Shibu = field.NewInt32(tableName, "shibu")
	_ykFoodnutrition.Clcnengliang = field.NewInt32(tableName, "clcnengliang")
	_ykFoodnutrition.Nengliang = field.NewInt32(tableName, "nengliang")
	_ykFoodnutrition.Shuifen = field.NewFloat64(tableName, "shuifen")
	_ykFoodnutrition.Danbai = field.NewFloat64(tableName, "danbai")
	_ykFoodnutrition.Zhifang = field.NewFloat64(tableName, "zhifang")
	_ykFoodnutrition.Tanshui = field.NewFloat64(tableName, "tanshui")
	_ykFoodnutrition.Xianwei = field.NewFloat64(tableName, "xianwei")
	_ykFoodnutrition.Huifen = field.NewFloat64(tableName, "huifen")
	_ykFoodnutrition.VA = field.NewFloat64(tableName, "VA")
	_ykFoodnutrition.B1 = field.NewFloat64(tableName, "B1")
	_ykFoodnutrition.B2 = field.NewFloat64(tableName, "B2")
	_ykFoodnutrition.B3 = field.NewFloat64(tableName, "B3")
	_ykFoodnutrition.VC = field.NewFloat64(tableName, "VC")
	_ykFoodnutrition.VE = field.NewFloat64(tableName, "VE")
	_ykFoodnutrition.Jia = field.NewFloat64(tableName, "jia")
	_ykFoodnutrition.Na = field.NewFloat64(tableName, "na")
	_ykFoodnutrition.Gai = field.NewFloat64(tableName, "gai")
	_ykFoodnutrition.Mei = field.NewFloat64(tableName, "mei")
	_ykFoodnutrition.Tie = field.NewFloat64(tableName, "tie")
	_ykFoodnutrition.Meng = field.NewFloat64(tableName, "meng")
	_ykFoodnutrition.Xin = field.NewFloat64(tableName, "xin")
	_ykFoodnutrition.Tong = field.NewFloat64(tableName, "tong")
	_ykFoodnutrition.Lin = field.NewFloat64(tableName, "lin")
	_ykFoodnutrition.Xi = field.NewFloat64(tableName, "xi")
	_ykFoodnutrition.Leibie = field.NewInt32(tableName, "leibie")
	_ykFoodnutrition.Lei = field.NewInt32(tableName, "lei")
	_ykFoodnutrition.Danguchun = field.NewFloat64(tableName, "danguchun")

	_ykFoodnutrition.fillFieldMap()

	return _ykFoodnutrition
}

type ykFoodnutrition struct {
	ykFoodnutritionDo

	ALL          field.Asterisk
	ID           field.Int32
	Docv         field.String
	Leixing      field.Int32
	Name         field.String
	Xuhao        field.String
	Shibu        field.Int32
	Clcnengliang field.Int32
	Nengliang    field.Int32
	Shuifen      field.Float64
	Danbai       field.Float64
	Zhifang      field.Float64
	Tanshui      field.Float64
	Xianwei      field.Float64
	Huifen       field.Float64
	VA           field.Float64
	B1           field.Float64
	B2           field.Float64
	B3           field.Float64
	VC           field.Float64
	VE           field.Float64
	Jia          field.Float64
	Na           field.Float64
	Gai          field.Float64
	Mei          field.Float64
	Tie          field.Float64
	Meng         field.Float64
	Xin          field.Float64
	Tong         field.Float64
	Lin          field.Float64
	Xi           field.Float64
	Leibie       field.Int32
	Lei          field.Int32
	Danguchun    field.Float64

	fieldMap map[string]field.Expr
}

func (y ykFoodnutrition) Table(newTableName string) *ykFoodnutrition {
	y.ykFoodnutritionDo.UseTable(newTableName)
	return y.updateTableName(newTableName)
}

func (y ykFoodnutrition) As(alias string) *ykFoodnutrition {
	y.ykFoodnutritionDo.DO = *(y.ykFoodnutritionDo.As(alias).(*gen.DO))
	return y.updateTableName(alias)
}

func (y *ykFoodnutrition) updateTableName(table string) *ykFoodnutrition {
	y.ALL = field.NewAsterisk(table)
	y.ID = field.NewInt32(table, "id")
	y.Docv = field.NewString(table, "docv")
	y.Leixing = field.NewInt32(table, "leixing")
	y.Name = field.NewString(table, "name")
	y.Xuhao = field.NewString(table, "xuhao")
	y.Shibu = field.NewInt32(table, "shibu")
	y.Clcnengliang = field.NewInt32(table, "clcnengliang")
	y.Nengliang = field.NewInt32(table, "nengliang")
	y.Shuifen = field.NewFloat64(table, "shuifen")
	y.Danbai = field.NewFloat64(table, "danbai")
	y.Zhifang = field.NewFloat64(table, "zhifang")
	y.Tanshui = field.NewFloat64(table, "tanshui")
	y.Xianwei = field.NewFloat64(table, "xianwei")
	y.Huifen = field.NewFloat64(table, "huifen")
	y.VA = field.NewFloat64(table, "VA")
	y.B1 = field.NewFloat64(table, "B1")
	y.B2 = field.NewFloat64(table, "B2")
	y.B3 = field.NewFloat64(table, "B3")
	y.VC = field.NewFloat64(table, "VC")
	y.VE = field.NewFloat64(table, "VE")
	y.Jia = field.NewFloat64(table, "jia")
	y.Na = field.NewFloat64(table, "na")
	y.Gai = field.NewFloat64(table, "gai")
	y.Mei = field.NewFloat64(table, "mei")
	y.Tie = field.NewFloat64(table, "tie")
	y.Meng = field.NewFloat64(table, "meng")
	y.Xin = field.NewFloat64(table, "xin")
	y.Tong = field.NewFloat64(table, "tong")
	y.Lin = field.NewFloat64(table, "lin")
	y.Xi = field.NewFloat64(table, "xi")
	y.Leibie = field.NewInt32(table, "leibie")
	y.Lei = field.NewInt32(table, "lei")
	y.Danguchun = field.NewFloat64(table, "danguchun")

	y.fillFieldMap()

	return y
}

func (y *ykFoodnutrition) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := y.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (y *ykFoodnutrition) fillFieldMap() {
	y.fieldMap = make(map[string]field.Expr, 33)
	y.fieldMap["id"] = y.ID
	y.fieldMap["docv"] = y.Docv
	y.fieldMap["leixing"] = y.Leixing
	y.fieldMap["name"] = y.Name
	y.fieldMap["xuhao"] = y.Xuhao
	y.fieldMap["shibu"] = y.Shibu
	y.fieldMap["clcnengliang"] = y.Clcnengliang
	y.fieldMap["nengliang"] = y.Nengliang
	y.fieldMap["shuifen"] = y.Shuifen
	y.fieldMap["danbai"] = y.Danbai
	y.fieldMap["zhifang"] = y.Zhifang
	y.fieldMap["tanshui"] = y.Tanshui
	y.fieldMap["xianwei"] = y.Xianwei
	y.fieldMap["huifen"] = y.Huifen
	y.fieldMap["VA"] = y.VA
	y.fieldMap["B1"] = y.B1
	y.fieldMap["B2"] = y.B2
	y.fieldMap["B3"] = y.B3
	y.fieldMap["VC"] = y.VC
	y.fieldMap["VE"] = y.VE
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
	y.fieldMap["leibie"] = y.Leibie
	y.fieldMap["lei"] = y.Lei
	y.fieldMap["danguchun"] = y.Danguchun
}

func (y ykFoodnutrition) clone(db *gorm.DB) ykFoodnutrition {
	y.ykFoodnutritionDo.ReplaceConnPool(db.Statement.ConnPool)
	return y
}

func (y ykFoodnutrition) replaceDB(db *gorm.DB) ykFoodnutrition {
	y.ykFoodnutritionDo.ReplaceDB(db)
	return y
}

type ykFoodnutritionDo struct{ gen.DO }

func (y ykFoodnutritionDo) Debug() *ykFoodnutritionDo {
	return y.withDO(y.DO.Debug())
}

func (y ykFoodnutritionDo) WithContext(ctx context.Context) *ykFoodnutritionDo {
	return y.withDO(y.DO.WithContext(ctx))
}

func (y ykFoodnutritionDo) ReadDB() *ykFoodnutritionDo {
	return y.Clauses(dbresolver.Read)
}

func (y ykFoodnutritionDo) WriteDB() *ykFoodnutritionDo {
	return y.Clauses(dbresolver.Write)
}

func (y ykFoodnutritionDo) Session(config *gorm.Session) *ykFoodnutritionDo {
	return y.withDO(y.DO.Session(config))
}

func (y ykFoodnutritionDo) Clauses(conds ...clause.Expression) *ykFoodnutritionDo {
	return y.withDO(y.DO.Clauses(conds...))
}

func (y ykFoodnutritionDo) Returning(value interface{}, columns ...string) *ykFoodnutritionDo {
	return y.withDO(y.DO.Returning(value, columns...))
}

func (y ykFoodnutritionDo) Not(conds ...gen.Condition) *ykFoodnutritionDo {
	return y.withDO(y.DO.Not(conds...))
}

func (y ykFoodnutritionDo) Or(conds ...gen.Condition) *ykFoodnutritionDo {
	return y.withDO(y.DO.Or(conds...))
}

func (y ykFoodnutritionDo) Select(conds ...field.Expr) *ykFoodnutritionDo {
	return y.withDO(y.DO.Select(conds...))
}

func (y ykFoodnutritionDo) Where(conds ...gen.Condition) *ykFoodnutritionDo {
	return y.withDO(y.DO.Where(conds...))
}

func (y ykFoodnutritionDo) Order(conds ...field.Expr) *ykFoodnutritionDo {
	return y.withDO(y.DO.Order(conds...))
}

func (y ykFoodnutritionDo) Distinct(cols ...field.Expr) *ykFoodnutritionDo {
	return y.withDO(y.DO.Distinct(cols...))
}

func (y ykFoodnutritionDo) Omit(cols ...field.Expr) *ykFoodnutritionDo {
	return y.withDO(y.DO.Omit(cols...))
}

func (y ykFoodnutritionDo) Join(table schema.Tabler, on ...field.Expr) *ykFoodnutritionDo {
	return y.withDO(y.DO.Join(table, on...))
}

func (y ykFoodnutritionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *ykFoodnutritionDo {
	return y.withDO(y.DO.LeftJoin(table, on...))
}

func (y ykFoodnutritionDo) RightJoin(table schema.Tabler, on ...field.Expr) *ykFoodnutritionDo {
	return y.withDO(y.DO.RightJoin(table, on...))
}

func (y ykFoodnutritionDo) Group(cols ...field.Expr) *ykFoodnutritionDo {
	return y.withDO(y.DO.Group(cols...))
}

func (y ykFoodnutritionDo) Having(conds ...gen.Condition) *ykFoodnutritionDo {
	return y.withDO(y.DO.Having(conds...))
}

func (y ykFoodnutritionDo) Limit(limit int) *ykFoodnutritionDo {
	return y.withDO(y.DO.Limit(limit))
}

func (y ykFoodnutritionDo) Offset(offset int) *ykFoodnutritionDo {
	return y.withDO(y.DO.Offset(offset))
}

func (y ykFoodnutritionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *ykFoodnutritionDo {
	return y.withDO(y.DO.Scopes(funcs...))
}

func (y ykFoodnutritionDo) Unscoped() *ykFoodnutritionDo {
	return y.withDO(y.DO.Unscoped())
}

func (y ykFoodnutritionDo) Create(values ...*models.YkFoodnutrition) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Create(values)
}

func (y ykFoodnutritionDo) CreateInBatches(values []*models.YkFoodnutrition, batchSize int) error {
	return y.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (y ykFoodnutritionDo) Save(values ...*models.YkFoodnutrition) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Save(values)
}

func (y ykFoodnutritionDo) First() (*models.YkFoodnutrition, error) {
	if result, err := y.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkFoodnutrition), nil
	}
}

func (y ykFoodnutritionDo) Take() (*models.YkFoodnutrition, error) {
	if result, err := y.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkFoodnutrition), nil
	}
}

func (y ykFoodnutritionDo) Last() (*models.YkFoodnutrition, error) {
	if result, err := y.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkFoodnutrition), nil
	}
}

func (y ykFoodnutritionDo) Find() ([]*models.YkFoodnutrition, error) {
	result, err := y.DO.Find()
	return result.([]*models.YkFoodnutrition), err
}

func (y ykFoodnutritionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.YkFoodnutrition, err error) {
	buf := make([]*models.YkFoodnutrition, 0, batchSize)
	err = y.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (y ykFoodnutritionDo) FindInBatches(result *[]*models.YkFoodnutrition, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return y.DO.FindInBatches(result, batchSize, fc)
}

func (y ykFoodnutritionDo) Attrs(attrs ...field.AssignExpr) *ykFoodnutritionDo {
	return y.withDO(y.DO.Attrs(attrs...))
}

func (y ykFoodnutritionDo) Assign(attrs ...field.AssignExpr) *ykFoodnutritionDo {
	return y.withDO(y.DO.Assign(attrs...))
}

func (y ykFoodnutritionDo) Joins(fields ...field.RelationField) *ykFoodnutritionDo {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Joins(_f))
	}
	return &y
}

func (y ykFoodnutritionDo) Preload(fields ...field.RelationField) *ykFoodnutritionDo {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Preload(_f))
	}
	return &y
}

func (y ykFoodnutritionDo) FirstOrInit() (*models.YkFoodnutrition, error) {
	if result, err := y.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkFoodnutrition), nil
	}
}

func (y ykFoodnutritionDo) FirstOrCreate() (*models.YkFoodnutrition, error) {
	if result, err := y.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.YkFoodnutrition), nil
	}
}

func (y ykFoodnutritionDo) FindByPage(offset int, limit int) (result []*models.YkFoodnutrition, count int64, err error) {
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

func (y ykFoodnutritionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = y.Count()
	if err != nil {
		return
	}

	err = y.Offset(offset).Limit(limit).Scan(result)
	return
}

func (y ykFoodnutritionDo) Scan(result interface{}) (err error) {
	return y.DO.Scan(result)
}

func (y ykFoodnutritionDo) Delete(models ...*models.YkFoodnutrition) (result gen.ResultInfo, err error) {
	return y.DO.Delete(models)
}

func (y *ykFoodnutritionDo) withDO(do gen.Dao) *ykFoodnutritionDo {
	y.DO = *do.(*gen.DO)
	return y
}
