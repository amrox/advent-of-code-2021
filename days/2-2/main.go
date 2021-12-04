package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	log.Println("aoc 2-2")

	fileArg := os.Args[1]

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(fileArg)

	if err != nil {
		log.Fatalf("failed to open")
	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	// and then a loop iterates through
	// and prints each of the slice values

	var aim = 0
	var hPos = 0
	var depth = 0

	for _, each_ln := range text {

		s := strings.Fields(each_ln)
		dir := s[0]
		val, _ := strconv.Atoi(s[1])

		switch dir {
		case "forward":
			hPos += val
			depth += (val * aim)
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	fmt.Println(hPos * depth)
}
