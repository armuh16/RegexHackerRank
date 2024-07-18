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
	inputFile, err := os.Open("inputDetectTheDomain.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("outputDetectTheDomain.txt")
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

	// Read the HTML fragment lines
	var textLines []string
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		textLines = append(textLines, line)
	}

	// Join all text lines into a single string
	textContent := strings.Join(textLines, " ")

	// Regular expression to find URLs
	re := regexp.MustCompile(`https?://(?:www\.)?([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)
	matches := re.FindAllStringSubmatch(textContent, -1)

	// Use a map to store unique domains
	domainSet := make(map[string]struct{})
	for _, match := range matches {
		domain := match[1]
		domainSet[domain] = struct{}{}
	}

	// Collect unique domains into a slice
	var domains []string
	for domain := range domainSet {
		domains = append(domains, domain)
	}

	// Sort the domains lexicographically
	sort.Strings(domains)

	// Print the sorted domains
	fmt.Fprintln(writer, strings.Join(domains, ";"))
}
