package directusgosdk

type Filter struct {
	FieldName string
	Operator  string
	Value     string
}

type Query struct {
	Fields []string
	Filter []Filter
	Search string
	Sort   []string
	Limit  uint
	Offset uint
	Page   uint
}

func NewQuery() *Query {
	return &Query{
		Fields: []string{},
		Filter: []Filter{},
		Search: "",
		Sort:   []string{},
		Limit:  100,
		Offset: 0,
		Page:   1,
	}
}
