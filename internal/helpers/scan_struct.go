package helpers

import (
	"errors"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

func ScanStruct(dest interface{}, r *sqlx.Rows, searchText string, replaceText string) error {
	v := reflect.ValueOf(dest)

	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}

	v = v.Elem()

	columns, err := r.Columns()
	if err != nil {
		return err
	}

	cols := make([]string, len(columns))
	for i, c := range columns {
		if strings.HasPrefix(c, searchText) {
			cols[i] = strings.Replace(c, searchText, replaceText, -1)
		} else {
			cols[i] = c
		}
	}

	m := r.Mapper
	fields := m.TraversalsByName(v.Type(), cols)
	values := make([]interface{}, len(columns))
	err = fieldsByTraversal(v, fields, values, true)
	if err != nil {
		return err
	}

	err = r.Scan(values...)
	if err != nil {
		return err
	}

	return r.Err()
}

func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}, ptrs bool) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("argument not a struct")
	}

	for i, traversal := range traversals {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}
		f := reflectx.FieldByIndexes(v, traversal)
		if ptrs {
			values[i] = f.Addr().Interface()
		} else {
			values[i] = f.Interface()
		}
	}
	return nil
}
