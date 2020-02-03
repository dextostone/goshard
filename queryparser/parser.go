package queryparser

import (
	"errors"
	"regexp"
	"strings"
)

// QueryParams will have cols,values corresponding sql commands
type QueryParams struct {
	Select  string
	From    string
	OrderBy string
}

// ValidateQuery will check is the query complies with the SQL standards
func ValidateQuery(query string) (bool, error) {

	// simple sql pattern
	pattern := "(select|SELECT) ([a-zA-Z]+(,[a-zA-Z]+)*) (from|FROM) [a-zA-Z]+(\\.[a-zA-Z]+)* ((limit|LIMIT) [0-9]+)? ((orderby|ORDERBY) (asc|desc|ASC|DESC))?;"
	matched, err := regexp.MatchString(pattern, query)
	if err != nil {
		return false, err
	}
	return matched, nil
}

// ParseQuery will parse the SQL string and return a class
func ParseQuery(query string) (QueryParams, error) {
	match, _ := ValidateQuery(query)
	if !match {
		return QueryParams{}, errors.New("Not a valid SQL query")
	}
	words := strings.Fields(query[0 : len(query)-1])
	selectVal := words[1]
	fromVal := words[3]
	orderByVal := ""
	for i := 0; i < len(words); i++ {
		if strings.EqualFold(words[i], "orderby") {
			orderByVal = strings.ToLower(words[i+1])
			break
		}
	}

	return QueryParams{Select: selectVal, From: fromVal, OrderBy: orderByVal}, nil
}
