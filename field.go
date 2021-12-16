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
	return reflect.DeepEqual(cur, zero)
}

func (f *Field) Name() string {
	return f.field.Name
}

func (f *Field) Kind() reflect.Kind {
	return f.value.Kind()
}

func (f *Field) Set(val interface{}) error {
	if !f.IsExported() {
		return errNotExported
	}
	if !f.value.CanSet() {
		return errNotSettable
	}
	given := reflect.ValueOf(val)
	if f.value.Kind() != given.Kind() {
		return fmt.Errorf("wrong kind. got %s want %s", given.Kind(), f.value.Kind())
	}
	f.value.Set(given)
	return nil
}

func (f *Field) Zero() error {
	zero := reflect.Zero(f.value.Type()).Interface()
	return f.Set(zero)
}

func (f *Field) Fields() []*Field {
	return getFields(f.value, f.defaultTag)
}

func (f *Field) Field(name string) *Field {
	field, ok := f.FieldOk(name)
	if !ok {
		panic("field not found")
	}
	return field
}

func (f *Field) FieldOk(name string) (*Field, bool) {
	value := &f.value
	if f.value.Kind() != reflect.Ptr {
		a := f.value.Addr()
		value = &a
	}
	v := strctVal(value.Interface())
	t := v.Type()

	field, ok := t.FieldByName(name)
	if !ok {
		return nil, false
	}
	return &Field{
		value:      v.FieldByName(name),
		field:      field,
	}, true
}
