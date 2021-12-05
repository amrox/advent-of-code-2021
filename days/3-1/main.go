package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// var mask = 0x0
	var counts [12]int
	for _, each_ln := range text {
		for i := 0; i < 12; i++ {
			if each_ln[i] == '1' {
				counts[i]++
			}

		}
	}
	var gamma = 0
	fmt.Println(counts)

	for i := 0; i < 12; i++ {
		if counts[i] > len(text)/2 {
			mask := 0x1 << (12 - i - 1)
			// fmt.Printf("mask %b\n", mask)
			gamma += mask
			// fmt.Printf("gamma %b\n", gamma)
		}
	}
	fmt.Printf("gamma %d \t%012b\n", gamma, gamma)

	//epsilon := (0x1 << 13) - gamma
	epsilon := ^gamma & 0x0fff

	fmt.Printf("epsilon %d \t%012b\n", epsilon, epsilon)

	fmt.Printf("answer: %d\n", gamma*epsilon)
}
