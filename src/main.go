package main

import (
	"journalist"
	"flag"
	"io/ioutil"
	"log"
	"fmt"
	"parser"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "text", "File to read")
	flag.Parse()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	chatlog := parser.ParseIRCLogBytes(data)
	article := journalist.GenerateArticle(chatlog)
	for _, s := range article {
		fmt.Println(s)
	}
}
