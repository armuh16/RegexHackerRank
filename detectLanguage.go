package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func detectLanguage(code string) string {
	// Java Detection
	javaPatterns := []string{
		`import java\.\S+;`,         // import java.io.*;
		`public class \S+ {`,        // public class SquareNum {
		`System\.out\.println\(`,    // System.out.println(
		`public static void main\(`, // public static void main(String args[])
	}
	for _, pattern := range javaPatterns {
		matched, _ := regexp.MatchString(pattern, code)
		if matched {
			return "Java"
		}
	}

	// C Detection
	cPatterns := []string{
		`#include <\S+>`, // #include <stdio.h>
		`#include "\S+"`, // #include "header.h"
		`printf\(`,       // printf(
		`scanf\(`,        // scanf(
		`int main\(`,     // int main()
		`return 0;`,      // return 0;
	}
	for _, pattern := range cPatterns {
		matched, _ := regexp.MatchString(pattern, code)
		if matched {
			return "C"
		}
	}

	// Python Detection
	pythonPatterns := []string{
		`\bimport \S+`,                         // import sys
		`def \S+\(`,                            // def function_name(
		`print\(`,                              // print(
		`^\s*#`,                                // # Comment
		`\bclass \S+:\b`,                       // class ClassName:
		`\bif __name__ == ['"]__main__['"]:\b`, // if __name__ == "__main__":
	}
	for _, pattern := range pythonPatterns {
		matched, _ := regexp.MatchString(pattern, code)
		if matched {
			return "Python"
		}
	}

	// If no specific keywords were found, guess based on syntax
	if matched, _ := regexp.MatchString(`;`, code); matched {
		return "C"
	}
	if matched, _ := regexp.MatchString(`:`, code); matched {
		return "Python"
	}

	return "Unknown"
}

func main() {
	// Open the inputTheEmailAdresses.txt file
	inputFile, err := os.Open("inputdetectLanguage.txt")
	if err != nil {
		fmt.Println("Error opening inputTheEmailAdresses.txt file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	var codeLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		codeLines = append(codeLines, line)
	}
	code := strings.Join(codeLines, "\n")
	result := detectLanguage(code)

	fmt.Fprintln(writer, result)
}
