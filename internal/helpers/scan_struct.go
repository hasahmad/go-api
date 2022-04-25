package helpers

import (
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

func ScanStruct(dest interface{}, r *sqlx.Rows, searchText string, replaceText string) error {
	var v reflect.Value
	v = reflect.ValueOf(dest)
	v = v.Elem()

	columns, err := r.Columns()
	if err != nil {
		return err
	}
	m := r.Mapper
	values := make([]interface{}, len(columns))
	cols := make([]string, len(columns))
	for i, c := range columns {
		if strings.HasPrefix(c, searchText) {
			cols[i] = strings.Replace(c, searchText, replaceText, -1)
		} else {
			cols[i] = c
		}
	}
	fields := m.TraversalsByName(v.Type(), cols)
	v = reflect.Indirect(v)
	for i, traversal := range fields {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}
		f := reflectx.FieldByIndexes(v, traversal)
		values[i] = f.Addr().Interface()
	}

	err = r.Scan(values...)
	if err != nil {
		return err
	}

	return nil
}
