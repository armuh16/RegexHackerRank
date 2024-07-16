package main

import (
	"bufio"
	"fmt"
	"html"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Open the input file
	inputFile, err := os.Open("InputDetectHTMLlinks.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("OutputDetectHTMLlinks.txt")
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

	// Read the HTML lines
	var htmlLines []string
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		htmlLines = append(htmlLines, line)
	}

	// Join all HTML lines into a single string
	htmlContent := strings.Join(htmlLines, " ")

	// Regular expression to capture <a> tags with href and inner text content
	re := regexp.MustCompile(`<a\s+href="([^"]*)".*?>(.*?)</a>`)

	// Find all matches
	matches := re.FindAllStringSubmatch(htmlContent, -1)

	// Process each match
	for _, match := range matches {
		href := match[1]
		text := stripTags(match[2])
		text = html.UnescapeString(text)
		fmt.Fprintf(writer, "%s,%s\n", href, strings.TrimSpace(text))
	}
}

// stripTags removes HTML tags from a string
func stripTags(html string) string {
	re := regexp.MustCompile(`<[^>]+>`)
	return re.ReplaceAllString(html, "")
}
