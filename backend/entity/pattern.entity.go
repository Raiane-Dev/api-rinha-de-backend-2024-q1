package entity

type Query struct {
	Table     string
	Columns   string
	Condition string
	Args      []any
	Values    string
	RawQuery  string
}
