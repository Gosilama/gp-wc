package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("arguments missing")
	}
	args := os.Args[1:]

	// if args[0] != "ccwc" {
	// 	log.Fatal("unrecognized command")
	// }

	var filename string
	var cmd string
	if len(args) < 2 {
		filename = args[0]
	} else {
		filename = args[1]
		cmd = args[0]
	}

	file, err := os.ReadFile(filename)
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

	switch cmd {
	case "-c":
		printBytes(args[1], file)
	case "-l":
		printLines(args[1], file)
	case "-w":
		printWords(args[1], file)
	case "-m":
		printCharacters(args[1], file)
	default:
		fmt.Printf("%v %v %v %v", len(strings.Split(string(file), "\n")), len(strings.Split(string(file), " ")), len(file), filename)
	}
}

func printWords(filename string, file []byte) {
	fmt.Printf("%v %v", len(strings.Split(string(file), " ")), filename)
}

func printBytes(filename string, b []byte) {
	fmt.Printf("%v %v", len(b), filename)
}

func printLines(filename string, b []byte) {
	fmt.Printf("%v %v", len(strings.Split(string(b), "\n")), filename)
}

func printCharacters(filename string, b []byte) {
	if utf8.Valid(b) {
		fmt.Printf("%v %v", len(strings.Split(string(b), "")), filename)
	} else {
		printBytes(filename, b)
	}
}
