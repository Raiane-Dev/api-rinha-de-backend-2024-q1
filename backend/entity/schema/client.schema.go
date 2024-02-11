package schema

type ClientSchema struct {
	ID             int    `db:"id"`
	Name           string `db:"nome"`
	Limit          int    `db:"limite"`
	Balance        int    `db:"saldo"`
	OpeningBalance int    `db:"saldo_inicial"`
}
