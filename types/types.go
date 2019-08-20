package gdstype

import (
	"fmt"
	"reflect"
)

// Types struct
type Types struct {
	typ *reflect.Type
}

// WatchType func
func (t *Types) WatchType(val interface{}) error {
	valType := reflect.TypeOf(val)
	a := 1
	if t.typ == nil {
		t.typ = &valType
		a++
	} else {
		if *t.typ != valType {
			return fmt.Errorf("%v should be of type %s", val, *t.typ)
		}
	}
	return nil
}
