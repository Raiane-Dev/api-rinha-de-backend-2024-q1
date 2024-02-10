package output

type ExtractOutput[T any] struct {
	Balance           Balance `json:"saldo"`
	LatestTransaction []T     `json:"ultimas_transacoes"`
}

type Balance struct {
	Total       int    `json:"total"`
	ExtractDate string `json:"data_extrato"`
	Limit       int    `json:"limite"`
}
