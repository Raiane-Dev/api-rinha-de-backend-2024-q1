package util

import (
	"database/sql"
	"fmt"
	"rinha_api/backend/config"
	"rinha_api/backend/entity"
	"strings"
)

func FindBy[T any](query *entity.Query) (data T, err error) {

	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s",
		strings.Join(query.Columns, ","),
		query.Table,
		query.Condition,
	)

	config.DatabaseInstance.Get(
		&data,
		sql,
		query.Args[:]...,
	)

	return
}

func FindMany[T any](query *entity.Query) (data T, err error) {

	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s",
		strings.Join(query.Columns, ","),
		query.Table,
		query.Condition,
	)

	err = config.DatabaseInstance.Select(
		&data,
		sql,
		query.Args[:]...,
	)

	return
}

func Insert(query *entity.Query) (res sql.Result, err error) {

	sql := fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s)`,
		query.Table,
		strings.Join(query.Columns, ","),
		query.Values,
	)

	res, err = config.DatabaseInstance.Exec(
		sql,
		query.Args[:]...,
	)

	return
}

func Update(query *entity.Query) (res sql.Result, err error) {

	sql := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		query.Table,
		query.Values,
		query.Condition,
	)

	res, err = config.DatabaseInstance.Exec(
		sql,
		query.Args[:]...,
	)

	return
}
