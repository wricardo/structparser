package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wricardo/structparser"
)

func main() {
	parsed, err := structparser.ParseDirectoryWithFilter("./", nil)
	if err != nil {
		log.Fatal(err)
	}
	encoded, _ := json.Marshal(parsed)
	fmt.Println(string(encoded))
}
