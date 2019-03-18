package gujson

import (
	"testing"
)

func TestSetDefaultValues(t *testing.T) {

	n := MyStructA{}
	n.SetDefaultValues()
	n.Print()
}

func TestPrint(t *testing.T) {

	n := MyStructA{}
	n.SetDefaultValues()
	n.Print()
}

func TestAddItem(t *testing.T) {

	c := MyStructA{}
	c.SetDefaultValues()

	d := MyStructA{}
	d.SetDefaultValues()

	arr := ArrMyStructA{}
	arr.AddItem(c)
	arr.AddItem(d)
	arr.Print()
}

func GetTestArray() ArrMyStructA {

	c := MyStructA{}
	c.SetDefaultValues()

	d := MyStructA{}
	d.SetDefaultValues()

	arr := ArrMyStructA{}
	arr.AddItem(c)
	arr.AddItem(d)

	return arr
}
