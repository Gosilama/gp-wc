package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type wcFunc func(string, []byte)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("arguments missing")
	}
	args := os.Args[1:]

	// if args[0] != "ccwc" {
	// 	log.Fatal("unrecognized command")
	// }
	argsMap := map[string]wcFunc{
		"-c": printBytes,
		"-l": printLines,
		"-w": printWords,
		"-m": printCharacters,
	}

	var filename string
	var cmd string
	var file []byte
	var err error

	if len(args) < 2 {
		if argsMap[args[0]] != nil {
			file, err = reader()
			filename = ""
			cmd = args[0]
		} else {
			filename = args[0]
			file, err = os.ReadFile(filename)
		}
	} else {
		filename = args[1]
		file, err = os.ReadFile(filename)
		cmd = args[0]
	}

	if err != nil {
		log.Fatal("could not read file")
	}

	if argsMap[cmd] != nil {
		argsMap[cmd](filename, file)
	} else {
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

func reader() ([]byte, error) {
	var bb bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Bytes()

		if len(text) < 0 {
			break
		}

		bb.Write(text)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bb.Bytes(), nil
}
