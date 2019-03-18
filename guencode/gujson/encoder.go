package gujson

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
)

func GenericStructToJSON(t interface{}) ([]byte, error) {

	var err error
	var tmp []byte

	switch t.(type) {

	case MyStructA:
		tmp, err = json.Marshal(t)
	case ArrMyStructA:
		tmp, err = json.Marshal(t)
		//	case MyStruct1:
		//		tmp, err = json.Marshal(t)
		//	case MyStruct2:
		//		tmp, err = json.Marshal(t)
		//	case MyStruct3:
		//		tmp, err = json.Marshal(t)
	}

	return tmp, err
}

func GenericBinaryToStruct(t []byte, tDest interface{}) (interface{}, error) {

	r := bytes.NewReader(t)
	b, err := ioutil.ReadAll(r)

	if err == nil {

		switch tDest.(type) {
		case MyStructA:
			err = json.Unmarshal(b, &tDest)
		case ArrMyStructA:
			err = json.Unmarshal(b, &tDest)

		}
	}
	return tDest, err
}

// // // // // // // // // // // // MyStructA // // // // // // // // // // // //

func (t *MyStructA) StructToJSON() ([]byte, error) {
	var tmp []byte
	tmp, err := json.Marshal(t)

	//	tmp, err := GenericStructToJSON(t)

	if err != nil {
		fmt.Println("ERR :: MyStructA.StructToJSON ::", err)
	}
	return tmp, err
}

func (data *MyStructA) BinaryToStruct(t []byte) error {

	r := bytes.NewReader(t)
	b, err := ioutil.ReadAll(r)

	if err == nil {
		errJson := json.Unmarshal(b, &data)
		if errJson != nil {
			fmt.Println("ERR :: MyStructA.BinaryToStruct :: ", errJson)
		}
	}
	return err
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //

func (t *ArrMyStructA) StructToJSON() ([]byte, error) {
	var tmp []byte
	tmp, err := json.Marshal(t)

	if err != nil {
		fmt.Println("ERR :: ArrMyStructA.StructToJSON ::", err)
	}
	return tmp, err
}

func (data *ArrMyStructA) BinaryToStruct(t []byte) error {

	r := bytes.NewReader(t)
	b, err := ioutil.ReadAll(r)

	if err == nil {
		errJson := json.Unmarshal(b, &data)
		if errJson != nil {
			fmt.Println("ERR :: ArrMyStructA.BinaryToStruct :: ", errJson)
		}
	}
	return err
}
