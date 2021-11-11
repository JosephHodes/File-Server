package main

import (
	"./Tests/FuncTests"
	"log"
	"os"
	"strings"
)

func main() {
	switch strings.ToUpper(os.Getenv("Mode")){
	case "TEST":
		err := FuncTests.FuncTest()
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Fatalln("fatal error couldn't get mode it was defined as: "+os.Getenv())
}
