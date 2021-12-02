package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	log.Println("aoc 1-1")

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

	var prevWindowSum int
	var increaseCount int = 0
	var lineNo = 0
	var window [3]int
	for _, each_ln := range text {
		curMeasurement, _ := strconv.Atoi(each_ln)

		if lineNo < 3 {
			window[lineNo] = curMeasurement

		} else {

			prevWindowSum = window[0] + window[1] + window[2]
			window[0] = window[1]
			window[1] = window[2]
			window[2] = curMeasurement

			var curWindowSum = window[0] + window[1] + window[2]

			if curWindowSum > prevWindowSum {
				increaseCount++
			}
		}
		lineNo++
	}
	log.Println(increaseCount)
}
