package gutemplates

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PackageName  string `yaml:"package_name"`
	TestFuncName string `yaml:"test_func_name"`
	FuncName     string `yaml:"func_name"`
	BasePath     string `yaml:"base_path"`

	// package_name: Flipfront
	// test_func_name: TestFuncName
	// func_name: FuncName
	// base_path: /tmp/go.utils/src/
}

func DefaultConfig() Config {

	c := Config{}
	c.PackageName = "Default"
	c.TestFuncName = "TestDefault"
	c.FuncName = "Default"
	c.BasePath = "/" + c.FuncName + "/"

	return c
}
func (c Config) Print() {

	fmt.Println(":: ", c.PackageName)
	fmt.Println(":: ", c.TestFuncName)
	fmt.Println(":: ", c.FuncName)
	fmt.Println(":: ", c.BasePath)
}

func YamlFileToStruct(tPath string) Config {

	c := DefaultConfig()
	source, err := ioutil.ReadFile(tPath)

	if err == nil {
		err = yaml.Unmarshal(source, &c)

		if err != nil {
			fmt.Println("err : ", err)
		} else {
			fmt.Println("--> ", c)
			return c
		}
	} else {
		fmt.Println("err : ", err)
	}
	return c
}

func TemplateTest(tYamlFilePath string, tTemplateName string, tTemplateFilePath string) {

	c := YamlFileToStruct(tYamlFilePath)

	m := make(map[string]interface{})

	m["PackageName"] = c.PackageName
	m["TestFuncName"] = c.TestFuncName
	m["FuncName"] = c.FuncName
	m["BasePath"] = c.BasePath

	tmpl, err := template.New(tTemplateName).ParseFiles(tTemplateFilePath)

	if err == nil {
		err = tmpl.Execute(os.Stdout, m)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
