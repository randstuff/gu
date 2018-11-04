package gujson

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
)

func GenericStructToJSON(t interface{}) ([]byte, error) {

	fmt.Println(" :: GenericStructToJSON ")

	var err error
	var tmp []byte

	switch t.(type) {

	case Cmd:
		tmp, err = json.Marshal(t)
	case ArrCmd:
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

// // // // // // // // // // // // // // // // // // // // // // // // // //

func (t *Cmd) StructToJSON() ([]byte, error) {
	var tmp []byte

	tmp, err := GenericStructToJSON(t)

	if err != nil {
		fmt.Println("ERR :: CmdDetails.StructToJSON ::", err)
	}
	return tmp, err
}

func (data *Cmd) BinaryToStruct(t []byte) error {

	r := bytes.NewReader(t)
	b, err := ioutil.ReadAll(r)

	if err == nil {
		errJson := json.Unmarshal(b, &data)
		if errJson != nil {
			//			Logger.Warning.Println("ERR :: CmdDetails.BinaryToStruct :: ", errJson)
			fmt.Println("ERR :: CmdDetails.BinaryToStruct :: ", errJson)
		}
	}
	return err
}

// // // // // // // // // // // // // // // // // // // // // // // // // //

func (t *ArrCmd) StructToJSON() ([]byte, error) {
	var tmp []byte
	tmp, err := json.Marshal(t)

	if err != nil {
		//		Logger.Warning.Println("ERR :: ArrCmdDetails.StructToJSON ::", err)
		fmt.Println("ERR :: ArrCmdDetails.StructToJSON ::", err)
	}
	return tmp, err
}

func (data *ArrCmd) BinaryToStruct(t []byte) error {

	r := bytes.NewReader(t)
	b, err := ioutil.ReadAll(r)

	if err == nil {
		errJson := json.Unmarshal(b, &data)
		if errJson != nil {
			//			Logger.Warning.Println("ERR :: ArrCmdDetails.BinaryToStruct :: ", errJson)
			fmt.Println("ERR :: ArrCmdDetails.BinaryToStruct :: ", errJson)
		}
	}
	return err
}
