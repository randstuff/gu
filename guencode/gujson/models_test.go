package gujson

import (
	"testing"
)

func TestSetDefaultValues(t *testing.T) {

	n := Cmd{}
	n.SetDefaultValues()
	n.Print()
}

func TestPrint(t *testing.T) {

	n := Cmd{}
	n.SetDefaultValues()
	n.Print()
}

func TestAddItem(t *testing.T) {

	c := Cmd{}
	c.SetDefaultValues()

	d := Cmd{}
	d.SetDefaultValues()

	arr := ArrCmd{}
	arr.AddItem(c)
	arr.AddItem(d)
	arr.Print()
}

func GetTestArray() ArrCmd {

	c := Cmd{}
	c.SetDefaultValues()

	d := Cmd{}
	d.SetDefaultValues()

	arr := ArrCmd{}
	arr.AddItem(c)
	arr.AddItem(d)

	return arr
}
