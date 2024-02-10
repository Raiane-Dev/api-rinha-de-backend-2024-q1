package model

import (
	"rinha_api/backend/entity"
	"rinha_api/backend/entity/schema"
	"rinha_api/backend/util"
)

type (
	TransactionEntity schema.TransactionSchema
)

func (data *TransactionEntity) Create() (err error) {
	sql := &entity.Query{
		Table:   "transacoes",
		Columns: []string{"tipo", "cliente_id", "valor", "descricao"},
		Values:  "?, ?, ?, ?",
		Args:    []any{data.Type, data.ClientID, data.Value, data.Description},
	}
	_, err = util.Insert(sql)

	return
}

func (TransactionEntity) FindMany(where string, args ...any) (accounts []TransactionEntity, err error) {
	query := &entity.Query{
		Table:     "transacoes",
		Columns:   []string{"*"},
		Condition: where,
		Args:      args,
	}
	accounts, err = util.FindMany[[]TransactionEntity](query)

	return
}

func (TransactionEntity) FindBy(where string, args ...any) (data TransactionEntity, err error) {
	query := &entity.Query{
		Table:     "transacoes",
		Columns:   []string{"*"},
		Condition: where,
		Args:      args,
	}

	data, err = util.FindBy[TransactionEntity](query)

	return
}
