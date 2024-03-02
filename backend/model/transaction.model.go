package model

import (
	"log"
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

func FindManyTransactions(where string, args ...any) (transactions []string) {
	transactions, err := util.FindMany[[]string](entity.Query{
		Table: "transacoes",
		Columns: `json_group_array( 
			json_object(
			'tipo', tipo, 
			'valor', valor,
			'realizado_em', realizado_em ,
			'descricao', descricao
			)
		)`,
		Condition: where,
		Args:      args,
	})
	if err != nil {
		log.Println("err read db", err)
	}

	return
}
