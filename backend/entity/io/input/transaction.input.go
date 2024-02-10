package input

type TransactionInput struct {
	Value       int    `json:"valor" validate:"required,min=1"`
	Type        string `json:"tipo" validate:"required,oneof=d c"`
	Description string `json:"descricao" validate:"required,min=1,max=10"`
}
