package structs

import (
	"errors"
	"reflect"
	"fmt"
)
var (
	errNotExported = errors.New("field is not exported")
	errNotSettable = errors.New("field is not settable")
)

type Field struct {
	value reflect.Value
	field reflect.StructField
	defaultTag string
}

func (f *Field) Tag(key string) string {
	return f.field.Tag.Get(key)
}

func (f *Field) Value() interface{} {
	return f.value.Interface()
}

func (f *Field) IsEmbedded() bool {
	return f.field.Anonymous
}

func (f *Field) IsExported() bool {
	return f.field.PkgPath == ""
}

func (f *Field) IsZero() bool {
	zero := reflect.Zero(f.value.Type()).Interface()
	cur := f.Value()

	return
}