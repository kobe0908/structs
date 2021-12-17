package structs

import (
	//"reflect"
	"testing"
)
type Foo struct {
	A string
	B int `structs:"y"`
	C bool `json:"c"`
	d string
	E *Baz
	x string `xml:"x"`
	Y []string
	Z map[string]interface{}
	*Bar
}

type Baz struct {
	A string
	B int
}

type Bar struct {
	E string
	F int
	g []string
}

func newStruct() *Struct {
	b := &Bar{
		E: "example",
		F: 2,
		g: []string{"zeynep","faith"},
	}
	f := &Foo {
		A: "gopher",
		C: true,
		d: "small",
		E: nil,
		Y: []string{"example"},
		Z: nil,
	}
	f.Bar = b
	return New(f)
}

func TestField_Set(t *testing.T) {
	s := newStruct()

	f := s.Field("A")
	err := f.Set("fatih")
	if err != nil {
		t.Error(err)
	}

	if f.Value().(string) != "fatih" {
		t.Errorf("want fatih")
	}

	f = s.Field("Y")
	err = f.Set([]string{"over","wih","op"})
	if err != nil {
		t.Error(err)
	}

	sliceLen := len(f.Value().([]string))
	if sliceLen != 3 {
		t.Errorf("want 3")
	}
}

func TestField_Tag(t *testing.T) {
	s := newStruct()

	v := s.Field("B").Tag("json")
	if v != "" {
		t.Errorf("should empty")
	}
}

func TestField_IsEmbedded(t *testing.T) {
	s := newStruct()
	if !s.Field("Bar").IsEmbedded() {
		t.Error("is an embed")
	}
}

func TestField_IsExported(t *testing.T) {
	s := newStruct()

	if s.Field("d").IsExported() {
		t.Errorf("1")
	}
}

func TestField_Name(t *testing.T) {
	s := newStruct()

	if s.Field("A").Name() != "A" {
		t.Errorf("name")
	}
}

func TestField_Field(t *testing.T) {
	s := newStruct()

	e := s.Field("Bar").Field("E")

	val, ok := e.Value().(string)
	if !ok {
		t.Error("val")
	}
	t.Log(val)

	defer func() {
		err := recover()
		if err == nil {
			t.Error("1")
		}
	}()
	_ = s.Field("Bar").Field("e")
}

func TestField_Fields(t *testing.T) {
	s := newStruct()

	fields := s.Field("Bar").Fields()
	if len(fields) != 3 {
		t.Error("")
	}
}


