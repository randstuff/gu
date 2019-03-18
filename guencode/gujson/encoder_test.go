package gujson

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {

	n := MyStructA{}
	n.SetDefaultValues()

	if out, err := n.StructToJSON(); err == nil {
		fmt.Println("out :", string(out))
	}
}

func TestDecode(t *testing.T) {

	tjson := "{\"Action\":\"hello\",\"Status\":\"hello\",\"Priority\":-42,\"StartDate\":\"hello\",\"URL\":\"hello\",\"IP\":\"hello\",\"Domain\":\"hello\",\"FQDN\":\"hello\",\"Options\":true}"

	c := MyStructA{}
	c.SetDefaultValues()
	c.BinaryToStruct([]byte(tjson))
	c.Print()
}

func TestEncodeArr(t *testing.T) {

	arr := GetTestArray()

	if out, err := arr.StructToJSON(); err == nil {
		fmt.Println("out :", string(out))
	}
}

func TestDecodeArr(t *testing.T) {

	tjson := "{\"Items\":[{\"Action\":\"N/A\",\"Status\":\"N/A\",\"Priority\":-1,\"StartDate\":\"N/A\",\"URL\":\"N/A\",\"IP\":\"N/A\",\"Domain\":\"N/A\",\"FQDN\":\"N/A\",\"Options\":null},{\"Action\":\"N/A\",\"Status\":\"N/A\",\"Priority\":-1,\"StartDate\":\"N/A\",\"URL\":\"N/A\",\"IP\":\"N/A\",\"Domain\":\"N/A\",\"FQDN\":\"N/A\",\"Options\":null}]}"

	arr := ArrMyStructA{}
	arr.BinaryToStruct([]byte(tjson))
	arr.Print()
}

/////////////////////////////////////////////////////////////////////////////////

func TestGenericStructToJson(t *testing.T) {

	arr := GetTestArray()
	if out, err := GenericStructToJSON(arr); err == nil {
		fmt.Println(" out : ", string(out))
	}

	c := MyStructA{}
	c.SetDefaultValues()

	if out, err := GenericStructToJSON(c); err == nil {
		fmt.Println(" out : ", string(out))
	}

}
