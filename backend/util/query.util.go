package util

import (
	"database/sql"
	"fmt"
	"rinha_api/backend/config"
	"rinha_api/backend/entity"
)

func FindBy[T any](query entity.Query) (data T, err error) {

	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s",
		query.Columns,
		query.Table,
		query.Condition,
	)

	config.ReaderDB.Get(
		&data,
		sql,
		query.Args[:]...,
	)

	return
}

func FindMany[T any](query entity.Query) (data T, err error) {

	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s",
		query.Columns,
		query.Table,
		query.Condition,
	)

	err = config.ReaderDB.Select(
		&data,
		sql,
		query.Args[:]...,
	)

	return
}

func Insert(query entity.Query) (res sql.Result, err error) {
	tx, err := config.WriterDB.Beginx()
	defer tx.Rollback()

	sql := fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s)`,
		query.Table,
		query.Columns,
		query.Values,
	)

	res, err = tx.Exec(
		sql,
		query.Args[:]...,
	)

	err = tx.Commit()

	return
}

func Update(query entity.Query) (res sql.Result, err error) {
	tx, err := config.WriterDB.Beginx()
	defer tx.Rollback()

	sql := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		query.Table,
		query.Values,
		query.Condition,
	)

	res, err = tx.Exec(
		sql,
		query.Args[:]...,
	)

	err = tx.Commit()

	return
}
