package schema

type ClientSchema struct {
	ID             int `db:"id"`
	Limit          int `db:"limite"`
	Balance        int `db:"saldo"`
	OpeningBalance int `db:"saldo_inicial"`
}
