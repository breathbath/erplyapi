package migrations

type Query struct {
	sql  string
	args []interface{}
}

func NewQuery(sql string, args ...interface{}) (Query) {
	return Query{
		sql,
		args,
	}
}