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

	smtp, err := config.ReaderDB.Preparex(sql)
	if err != nil {
		return
	}
	smtp.Get(&data, query.Args[:]...)

	return
}

func FindMany[T any](query entity.Query) (data T, err error) {

	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s",
		query.Columns,
		query.Table,
		query.Condition,
	)

	smtp, err := config.ReaderDB.Preparex(sql)
	if err != nil {
		return
	}
	smtp.Select(&data, query.Args[:]...)
	return
}

func FindRaw[T any](query entity.Query) (data T, err error) {

	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s",
		query.Columns,
		query.Table,
		query.Condition,
	)

	smtp, err := config.ReaderDB.Preparex(sql)
	if err != nil {
		return
	}
	smtp.Select(&data, query.Args[:]...)
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

func InsertOrUpdate(query entity.Query) (res sql.Result, err error) {

	tx, err := config.WriterDB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return
	}

	sql := fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s) ON CONFLICT do UPDATE SET %s`,
		query.Table,
		query.Columns,
		query.Values,
		query.RawQuery,
	)

	res, err = tx.Exec(
		sql,
		query.Args[:]...,
	)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}
