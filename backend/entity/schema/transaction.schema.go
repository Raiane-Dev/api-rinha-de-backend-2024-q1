package schema

type TransactionSchema struct {
	ID          int    `json:"-" db:"id"`
	ClientID    int    `json:"-" db:"cliente_id"`
	Type        string `json:"tipo" db:"tipo"`
	Value       int    `json:"valor" db:"valor"`
	Description string `json:"descricao" db:"descricao"`
	CreatedAt   string `json:"realizado_em" db:"realizado_em"`
}
