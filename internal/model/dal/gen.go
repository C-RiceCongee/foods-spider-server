// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                    = new(Query)
	File                 *file
	User                 *user
	YkFoodnutrition      *ykFoodnutrition
	YkFoodnutritionCopy1 *ykFoodnutritionCopy1
	YkYycfb              *ykYycfb
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	File = &Q.File
	User = &Q.User
	YkFoodnutrition = &Q.YkFoodnutrition
	YkFoodnutritionCopy1 = &Q.YkFoodnutritionCopy1
	YkYycfb = &Q.YkYycfb
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		File:                 newFile(db, opts...),
		User:                 newUser(db, opts...),
		YkFoodnutrition:      newYkFoodnutrition(db, opts...),
		YkFoodnutritionCopy1: newYkFoodnutritionCopy1(db, opts...),
		YkYycfb:              newYkYycfb(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	File                 file
	User                 user
	YkFoodnutrition      ykFoodnutrition
	YkFoodnutritionCopy1 ykFoodnutritionCopy1
	YkYycfb              ykYycfb
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		File:                 q.File.clone(db),
		User:                 q.User.clone(db),
		YkFoodnutrition:      q.YkFoodnutrition.clone(db),
		YkFoodnutritionCopy1: q.YkFoodnutritionCopy1.clone(db),
		YkYycfb:              q.YkYycfb.clone(db),
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
		db:                   db,
		File:                 q.File.replaceDB(db),
		User:                 q.User.replaceDB(db),
		YkFoodnutrition:      q.YkFoodnutrition.replaceDB(db),
		YkFoodnutritionCopy1: q.YkFoodnutritionCopy1.replaceDB(db),
		YkYycfb:              q.YkYycfb.replaceDB(db),
	}
}

type queryCtx struct {
	File                 *fileDo
	User                 *userDo
	YkFoodnutrition      *ykFoodnutritionDo
	YkFoodnutritionCopy1 *ykFoodnutritionCopy1Do
	YkYycfb              *ykYycfbDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		File:                 q.File.WithContext(ctx),
		User:                 q.User.WithContext(ctx),
		YkFoodnutrition:      q.YkFoodnutrition.WithContext(ctx),
		YkFoodnutritionCopy1: q.YkFoodnutritionCopy1.WithContext(ctx),
		YkYycfb:              q.YkYycfb.WithContext(ctx),
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