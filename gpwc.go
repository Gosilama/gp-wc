package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("arguments missing")
	}
	args := os.Args[1:]

	// if args[0] != "ccwc" {
	// 	log.Fatal("unrecognized command")
	// }

	if len(args) < 2 {
		log.Fatal("missing arguments, please check command")
	}

	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal("could not open file")
	}

	data := make([]byte, 1024)
	n, err := file.Read(data)

	if err != nil {
		log.Fatal("could not read file")
	}

	switch args[0] {
	case "-c":
		getBytes(args[1], n)
	}
}

func getBytes(f string, n int) {
	fmt.Printf("%v %v", n, f)
}
