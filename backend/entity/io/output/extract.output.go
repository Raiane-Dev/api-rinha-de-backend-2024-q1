package output

import "encoding/json"

type ExtractOutput struct {
	Balance           Balance         `json:"saldo"`
	LatestTransaction json.RawMessage `json:"ultimas_transacoes"`
}

type Balance struct {
	Total       int    `json:"total"`
	ExtractDate string `json:"data_extrato"`
	Limit       int    `json:"limite"`
}
