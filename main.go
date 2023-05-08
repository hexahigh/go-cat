package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

var hexFlag bool

func init() {
	flag.BoolVar(&hexFlag, "hex", false, "Usage: --hex to use hexadecimal for the output")
}

func main() {
	flag.Parse()
	pathOpt := flag.Arg(1)
	fmt.Println(pathOpt)
	if len(os.Args[1:]) == 0 {
		fmt.Println("Too few arguments specified")
	} else if hexFlag {
		path := flag.Arg(0)
		readLinesHex(path)
	} else {
		path := flag.Arg(0)
		readLines(path)
	}

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return lines, scanner.Err()
}

func readLinesHex(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// Read the file in chunks
	buf := make([]byte, 256) // Adjust the chunk size as needed
	for {
		_, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		// Convert the chunk to hexadecimal and print it
		fmt.Printf("%s", hex.Dump(buf))
	}
	return nil, nil
}
