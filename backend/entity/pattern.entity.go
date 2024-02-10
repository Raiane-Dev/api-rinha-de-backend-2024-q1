package entity

type Query struct {
	Table     string
	Columns   []string
	Condition string
	Args      []any
	Values    string
	GroupBy   string
	Join      []Inner
	Structure interface{}
}

type Inner struct {
	Table      string
	Union      bool
	Compar     []string
	Structure  interface{}
	InnerChild InnerChild
}

type InnerChild struct {
	Table     string
	Compar    []string
	Structure interface{}
}
