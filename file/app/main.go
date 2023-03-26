package main

import (
	"github.com/MenciusCheng/go-text-template/parse"
	"io/fs"
	"io/ioutil"
)

func main() {
	textStr, err := ioutil.ReadFile("file/json/hello.json")
	if err != nil {
		panic(err)
	}

	textMap, err := parse.StringToMap(string(textStr))
	if err != nil {
		panic(err)
	}

	tmplStr, err := ioutil.ReadFile("file/tmpl/hello.tmpl")
	if err != nil {
		panic(err)
	}

	res, err := parse.JsonMap(string(tmplStr), textMap)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("out/hello.txt", []byte(res), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
