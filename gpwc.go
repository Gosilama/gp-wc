package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

	file, err := os.ReadFile(args[1])
	// file, err := os.Open(args[1])

	// if err != nil {
	// 	log.Fatal("could not open file")
	// }

	// data := make([]byte, 1024)
	// n, err := file.Read(data)

	if err != nil {
		log.Fatal("could not read file")
	}

	// argsMap := map[string]interface{}{
	// 	"-c": func(string, int)printBytes(args[1], n),
	// }

	// argsMap := map[string]func(string, string)

	switch args[0] {
	case "-c":
		printBytes(args[1], file)
	case "-l":
		printLines(string(args[1]), file)
	}
}

func printBytes(filename string, b []byte) {
	fmt.Printf("%v %v", len(b), filename)
}

func printLines(filename string, b []byte) {
	fmt.Printf("%v %v", len(strings.Split(string(b), "\n")), filename)
}
