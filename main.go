package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/racsoraul/tocsv/csv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal(`
		Path to input file must be provided!

		Usage: tocsv <path_to_file>
		`)
	}
	filePath := args[1:][0]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Can't open file")
	}
	defer file.Close()

	totalLines, err := csv.LineCounter(file)
	if err != nil {
		log.Fatal("Error reading file")
	}

	bar := pb.Full.Start(totalLines)

	file.Seek(0, io.SeekStart)

	csvFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("Can't create output file")
	}
	defer csvFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(csvFile)

	fmt.Print("Transforming file...\n\n")
	lineCounter := 1
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		line, err := csv.ProcessLines(scanner.Text(), lineCounter == 1)
		if err != nil {
			log.Printf("Skipped line %d - Reason: %v", lineCounter, err)
		} else {
			writer.WriteString(line + "\n")
		}

		bar.Increment()
		lineCounter++
	}

	if err := writer.Flush(); err != nil {
		panic("Not all data could be written to the output file")
	}

	bar.Finish()

	fmt.Println("\nCSV file succesfully created!")
}
