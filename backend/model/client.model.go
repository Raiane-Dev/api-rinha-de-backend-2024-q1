package model

import (
	"rinha_api/backend/entity"
	"rinha_api/backend/entity/schema"
	"rinha_api/backend/util"
)

type (
	ClientEntity schema.ClientSchema
)

func (data ClientEntity) Create() (err error) {

	_, err = util.Insert(entity.Query{
		Table:   "clientes",
		Columns: "limite, saldo, saldo_inicial",
		Values:  "?, ?, ?",
		Args:    []any{data.Limit, data.Balance, data.OpeningBalance},
	})

	return
}

func (data ClientEntity) Update(where string, args ...any) (err error) {

	_, err = util.Update(entity.Query{
		Table:     "clientes",
		Condition: where,
		Values:    "limite = ?, saldo = ?, saldo_inicial = ?",
		Args:      append([]any{data.Limit, data.Balance, data.OpeningBalance}, args...),
	})

	return
}

func FindByClient(where string, args ...any) (data ClientEntity, err error) {

	data, err = util.FindBy[ClientEntity](entity.Query{
		Table:     "clientes",
		Columns:   "*",
		Condition: where,
		Args:      args,
	})

	return
}
