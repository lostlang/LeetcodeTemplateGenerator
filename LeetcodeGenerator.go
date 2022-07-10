package main

import (
	"os"
	"strings"

	"text/template"
)

type FuncName struct {
	Normal string
	Upper  string
}

func generateName(words []string) string {
	var out string
	for _, word := range words {
		out = out + strings.Title(word)
	}
	return out
}

func main() {
	var arg []string = os.Args[1:]
	var funcName = arg[len(arg)-1]
	arg = arg[:len(arg)-2]

	var fileName = generateName(arg) + ".go"
	var fileNameTest = generateName(arg) + "_test.go"

	var fileTemplateMain, _ = os.ReadFile("main.template")
	var fileTemplateTest, _ = os.ReadFile("test.template")

	mainTemplate, _ := template.New("Main").Parse(string(fileTemplateMain))
	testTemplate, _ := template.New("Test").Parse(string(fileTemplateTest))

	var outMain, _ = os.Create(fileName)
	defer outMain.Close()

	var outTest, _ = os.Create(fileNameTest)
	defer outTest.Close()

	mainTemplate.Execute(outMain, funcName)
	testTemplate.Execute(outTest, FuncName{funcName, strings.Title(funcName)})
}
