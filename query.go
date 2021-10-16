package directusgosdk

import (
	"fmt"
	"strconv"
	"strings"
)

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
	Limit  uint64
	Offset uint64
	Page   uint64
	Deep   string
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

func (q *Query) ToQueryString() string {
	qs := []string{}
	if len(q.Fields) > 0 {
		qs = append(qs, "fields="+strings.Join(q.Fields, ","))
	}

	if q.Search != "" {
		qs = append(qs, "search="+q.Search)
	}

	if len(q.Sort) > 0 {
		qs = append(qs, "sort="+strings.Join(q.Sort, ","))
	}

	if q.Limit > 0 {
		qs = append(qs, "limit="+strconv.FormatUint(q.Limit, 10))
	}

	if q.Offset > 0 {
		qs = append(qs, "offset="+strconv.FormatUint(q.Offset, 10))
	}

	if q.Page > 0 {
		qs = append(qs, "page="+strconv.FormatUint(q.Page, 10))
	}

	if len(q.Filter) > 0 {
		fqs := []string{}

		for _, f := range q.Filter {
			fqs = append(fqs, "filter["+f.FieldName+"]["+f.Operator+"]="+f.Value)
		}

		qs = append(qs, strings.Join(fqs, "&"))
	}

	if q.Deep != "" {
		qs = append(qs, fmt.Sprintf("deep=%s", q.Deep))
	}

	return strings.Join(qs, "&")
}
