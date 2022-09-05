package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: input-file\n")
		os.Exit(1)
	}
	infile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %s", err.Error())
		os.Exit(1)
	}
	defer infile.Close()

	var chrom, readName, readString string
	var pos int
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Sscanf(scanner.Text(), "%s %s %s %d", &chrom, &readName, &readString, &pos)
		// Now we have the data; we just need to output it in Simple-SAM format
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
