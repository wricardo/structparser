package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"

	"github.com/wricardo/structparser"
)

func main() {
	filter := func(fi fs.FileInfo) bool {
		return fi.Name() == "simple_struct.go"
	}
	parsed, err := structparser.ParseDirectoryWithFilter("./example", filter)
	if err != nil {
		log.Fatal(err)
	}

	pretty, _ := json.MarshalIndent(parsed, "", "\t")
	fmt.Println("parsed", string(pretty))
}
