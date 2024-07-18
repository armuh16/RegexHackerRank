package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Open the input file
	inputFile, err := os.Open("inputTheEmailAdresses.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("outputTheEmailAdresses.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	// Read the number of lines
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	// Read the text fragment lines
	var textLines []string
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		textLines = append(textLines, line)
	}

	// Join all text lines into a single string
	textContent := strings.Join(textLines, " ")

	// Regular expression to find email addresses
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find all email addresses
	matches := re.FindAllString(textContent, -1)

	// Use a map to store unique email addresses
	emailSet := make(map[string]struct{})
	for _, email := range matches {
		emailSet[email] = struct{}{}
	}

	// Collect unique email addresses into a slice
	var emails []string
	for email := range emailSet {
		emails = append(emails, email)
	}

	// Sort the email addresses lexicographically
	sort.Strings(emails)

	// Print the sorted email addresses
	fmt.Fprintln(writer, strings.Join(emails, ";"))
}
