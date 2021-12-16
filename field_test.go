package structs

import (
	"reflect"
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

