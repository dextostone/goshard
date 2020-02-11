package queryparser

import (
	"fmt"
	"testing"
)

func TestParseQuery(t *testing.T) {
	s := "select a,b,c from db.dbname limit 1 orderby asc;"
	qp, er := ParseQuery(s)
	if er == nil {
		fmt.Println(qp.OrderBy)
	}

	s = "This is not a valid SQL query"
	qp, er = ParseQuery(s)
	if er == nil {
		fmt.Println(qp.OrderBy)
	}
}
