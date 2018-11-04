package guyaml

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// // // // // // // // // // // // // // // // // // // // // // // // // //

type TestStruct struct {
	Foo string `yaml:"Foo"`
	Bar struct {
		Value0 int    `yaml:"Value0"`
		Value1 []int  `yaml:"Value1"`
		Value2 string `yaml:"Value2"`
	} `yaml:"Bar"`
}

func (t TestStruct) StructToYaml() {

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("---  dump: %s \n", string(d))

}

// // // // // // // // // // // // // // // // // // // // // // // // // //

type Config struct {
	Foo string
	Bar []string
}

func (c Config) YamlFileToStruct(tPath string) {

	source, err := ioutil.ReadFile(tPath)
	if err == nil {
		err = yaml.Unmarshal(source, &c)

		if err != nil {
			fmt.Println("err : ", err)
		} else {
			fmt.Println("=> ", c.Bar)
			fmt.Println("=> ", c.Foo)
		}
	}
}
