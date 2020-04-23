package querybuilder

import (
	"fmt"
	"strings"
	"regexp"
)

type QueryBuilder struct {
	_select []string 
	_from string
	_where [][]string 
	_orwhere [][]string 
	_orderBy string 
	_groupBy string
}

func (q *QueryBuilder)Select(fields... string) *QueryBuilder{	
	q._select = fields
	return q
}

func (q *QueryBuilder) AddSelect(fields... string) *QueryBuilder{	
	q._select = append(q._select, fields...)
	return q
}

func (q *QueryBuilder)From(table string) *QueryBuilder{
	q._from = table
	return q
}

func (q *QueryBuilder)Where(conditions... string) *QueryBuilder{
	q._where = append(q._where, conditions)
	return q
}

func (q *QueryBuilder) AndWhere(conditions... string) *QueryBuilder{
	return q.Where(conditions...)
}

func (q *QueryBuilder) OrWhere(conditions... string) *QueryBuilder{
	q._orwhere = append(q._orwhere, conditions)
	return q
}

func (q *QueryBuilder)OrderBy(sort string, direction string) *QueryBuilder{
	q._orderBy = sort + " " + direction
	return q
}

func (q *QueryBuilder) GetQueryString() string {

	selectSlice := make([]string, len(q._select))
	for i, v := range q._select {
		 selectSlice[i] = fmt.Sprintf("`%s`", trim(v))
	}

	// build where
	var whereSlice []string
	for _, v := range q._where {
		whereSlice = append(whereSlice, strings.Join(v, ""))
	}
	
	// build OrWhere
	var orWhereSlice []string
	for _, v := range q._orwhere {
		orWhereSlice = append(orWhereSlice, strings.Join(v, ""))
	}

	var whereStr string = strings.Join(whereSlice, " AND ")

	if len(orWhereSlice) > 0 {
		whereStr = fmt.Sprintf("(%s) OR %s", whereStr, strings.Join(orWhereSlice, " AND ")) 
	}

	fmt.Println(orWhereSlice)

	qr := fmt.Sprintf("SELECT %s FROM %s WHERE %s", 
		strings.Join(selectSlice, ","),
		fmt.Sprintf("`%s`", trim(q._from)),		
		whereStr,
	)
	return trim(qr)
}
func trim(str string) string {
	re := regexp.MustCompile(`[\s]+`)

	// replace nhiều khoảng trắng = 1 khoảng trắng
	str = re.ReplaceAllString(str, " ")

	// trim chuỗi
	return strings.TrimSpace(str)
}

///////////////////////////////////////////////////////////////////////////////////////

type ORM struct {
	table string
}

func Table(tbl string) ORM{
	return ORM{table: tbl}
}