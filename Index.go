package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	switch strings.ToUpper(os.Getenv("Mode")){
	case "TEST":
		
	}
	log.Fatalln("fatal error couldn't get mode it was defined as: "+os.Getenv())
}
