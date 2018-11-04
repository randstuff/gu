package gujson

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {

	n := Cmd{}
	n.SetDefaultValues()

	if out, err := n.StructToJSON(); err == nil {
		fmt.Println("out :", string(out))
	}
}

func TestDecode(t *testing.T) {

	tjson := `"{\"Action\":\"hello\",
			\"Status\":\"hello\",
			\"Priority\":-42,
			\"AAAA\":\"hello\",
			\"BBBB\":\"hello\",
			\"CCCC\":\"hello\",
			\"DDDD\":\"hello\",
			\"Options\":true}"`

	c := Cmd{}
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

	tjson := `"{\"Items\":[{
						\"Action\":\"N/A\",
						\"Status\":\"N/A\",
						\"Priority\":-1,
						\"AAAA\":\"N/A\",
						\"BBBB\":\"N/A\",
						\"CCCC\":\"N/A\",
						\"DDDD\":\"N/A\",
						\"Options\":null},
						{
						\"Action\":\"N/A\",
						\"Status\":\"N/A\",
						\"Priority\":-1,
						\"AAAA\":\"N/A\",
						\"BBBB\":\"N/A\",
						\"CCCC\":\"N/A\",
						\"DDDD\":\"N/A\",
						\"Options\":null}
						]}"`

	arr := ArrCmd{}
	arr.BinaryToStruct([]byte(tjson))
	arr.Print()
}

/////////////////////////////////////////////////////////////////////////////////

func TestGenericStructToJson(t *testing.T) {

	arr := GetTestArray()
	if out, err := GenericStructToJSON(arr); err == nil {
		fmt.Println(" out : ", string(out))
	}

	c := Cmd{}
	c.SetDefaultValues()

	if out, err := GenericStructToJSON(c); err == nil {
		fmt.Println(" out : ", string(out))
	}

}

func TestGenericGenericJsonToStruct(t *testing.T) {

	tjson := "{\"Action\":\"hello\",\"Status\":\"hello\",\"Priority\":-42,\"AAAA\":\"hello\",\"BBBB\":\"hello\",\"CCCC\":\"hello\",\"DDDD\":\"hello\",\"Options\":true}"
	out := ProcessJSON(tjson)
	fmt.Println(out)

}
