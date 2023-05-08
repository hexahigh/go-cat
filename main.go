package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) > 1 || len(os.Args[1:]) == 0 {
		fmt.Println("Too many or too few arguments specified")
	} else {
		path := os.Args[1]
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
