package schema

type TransactionSchema struct {
	ClientID    int    `json:"-" db:"cliente_id"`
	Data        []byte `json:"-" db:"data"`
	Type        string `json:"tipo" db:"tipo"`
	Value       int    `json:"valor" db:"valor"`
	Description string `json:"descricao" db:"descricao"`
	CreatedAt   string `json:"realizado_em" db:"realizado_em"`
}
