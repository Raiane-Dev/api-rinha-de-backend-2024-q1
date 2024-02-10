package model

import (
	"rinha_api/backend/entity"
	"rinha_api/backend/entity/schema"
	"rinha_api/backend/util"
)

type (
	ClientEntity schema.ClientSchema
)

func (data *ClientEntity) Create() (err error) {
	sql := &entity.Query{
		Table:   "clientes",
		Columns: []string{"limite", "saldo", "saldo_inicial"},
		Values:  "?, ?",
		Args:    []any{data.Limit, data.Balance, data.OpeningBalance},
	}
	_, err = util.Insert(sql)

	return
}

func (data *ClientEntity) Update(where string, args ...any) (err error) {
	sql := &entity.Query{
		Table:     "clientes",
		Condition: where,
		Values:    "limite = ?, saldo = ?, saldo_inicial = ?",
		Args:      append([]any{data.Limit, data.Balance, data.OpeningBalance}, args...),
	}
	_, err = util.Update(sql)

	return
}

func (ClientEntity) FindMany(where string, args ...any) (accounts []ClientEntity, err error) {
	query := &entity.Query{
		Table:     "clientes",
		Columns:   []string{"*"},
		Condition: where,
		Args:      args,
	}
	accounts, err = util.FindMany[[]ClientEntity](query)

	return
}

func (ClientEntity) FindBy(where string, args ...any) (data ClientEntity, err error) {
	query := &entity.Query{
		Table:     "clientes",
		Columns:   []string{"*"},
		Condition: where,
		Args:      args,
	}

	data, err = util.FindBy[ClientEntity](query)

	return
}
