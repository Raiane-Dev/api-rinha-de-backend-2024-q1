package model

import (
	"rinha_api/backend/entity"
	"rinha_api/backend/entity/schema"
	"rinha_api/backend/util"
)

type (
	TransactionEntity schema.TransactionSchema
)

func (data TransactionEntity) Create() (err error) {

	_, err = util.Insert(entity.Query{
		Table:   "transacoes",
		Columns: "tipo, cliente_id, valor, descricao",
		Values:  "?, ?, ?, ?",
		Args:    []any{data.Type, data.ClientID, data.Value, data.Description},
	})

	return
}

func FindManyTransactions(where string, args ...any) (accounts []map[string]string) {
	accounts, _ = util.FindMany[[]map[string]string](entity.Query{
		Table:     "transacoes",
		Columns:   "*",
		Condition: where,
		Args:      args,
	})

	return
}
