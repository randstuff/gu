package gujson

import (
	"encoding/json"
	"errors"
	"fmt"
)

// source : https://github.com/beyondns/gotips/
//
func ProcessJSON(s string) map[string]interface{} {
	var m map[string]interface{}

	if err := json.Unmarshal([]byte(s), &m); err != nil {
		panic(err)
	}

	if err := passM(m); err != nil {
		panic(err)
	}
	return m
}

// source : https://github.com/beyondns/gotips/
//
func passKV(k string, v interface{}) error {
	var ok bool

	if _, ok = v.(map[string]interface{}); ok {
		err := passM(v.(map[string]interface{}))
		if err != nil {
			return err
		}
	}

	if _, ok = v.([]interface{}); ok {
		err := passA(v.([]interface{}))
		if err != nil {
			return err
		}
	}

	fmt.Printf("%s : %v\n", k, v)

	return nil
}

// source : https://github.com/beyondns/gotips/
//
func passM(m map[string]interface{}) error {
	for k, v := range m {
		if err := passKV(k, v); err != nil {
			return err
		}
	}
	return nil
}

func passA(a []interface{}) error {
	for _, v := range a {
		switch v.(type) {
		case map[string]interface{}:
			if err := passM(v.(map[string]interface{})); err != nil {
				return err
			}
		case []interface{}:
			if err := passA(v.([]interface{})); err != nil {
				return err
			}
		default:
			return errors.New("wrong expression in array")
		}
	}
	return nil
}
