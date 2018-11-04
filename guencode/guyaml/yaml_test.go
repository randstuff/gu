package guyaml

import (
	"fmt"
	"testing"
)

func TestNewReadYamlFile(t *testing.T) {

	var c Config

	c.YamlFileToStruct("test.yaml")
	fmt.Println(":: ", c.Foo)
}

func TestDecodeByte(t *testing.T) {

	c := TestStruct{}

	c.Foo = "test"

	c.Bar.Value0 = 1
	c.Bar.Value1 = []int{42, 42}
	c.Bar.Value2 = "hello"

	c.StructToYaml()

}
