package main

import (
	"./Tests/FuncTests"
	"log"
	"os"
	"strings"
)

func main() {
	mode := strings.ToUpper(os.Getenv("Mode"))
	switch mode {
	case "TEST":
		err := FuncTests.FuncTest()
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Fatalln("fatal error couldn't get mode it was defined as: "+mode)
}
