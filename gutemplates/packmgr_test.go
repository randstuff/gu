package gutemplates

import (
	"fmt"

	"testing"
)

func TestYamlFileToStruct(t *testing.T) {

	c := DefaultConfig()
	c = YamlFileToStruct("package.yaml")
	fmt.Println(" ::: ", c)

}

func TestTemplateTest(t *testing.T) {

	TemplateTest("package.yaml", "package.template", "files/package.template")
	TemplateTest("package.yaml", "package_test.template", "files/package_test.template")
}
