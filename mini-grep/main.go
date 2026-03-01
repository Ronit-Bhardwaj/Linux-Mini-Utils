package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func grepFile(pattern *regexp.Regexp, scanner *bufio.Scanner, source string, quiet bool) {
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if pattern.MatchString(line) {
			if quiet {
				fmt.Println(line)
			} else {
				if source != "" {
					fmt.Printf("%s:%d: %s\n", source, lineNum, line)
				} else {
					fmt.Printf("%d: %s\n", lineNum, line)
				}
			}
		}
		lineNum++
	}
}

func main() {
	quiet := flag.Bool("q", false, "Omit line numbers from output")
	pattern := flag.String("e", "", "Regex pattern to search for (required)")
	flag.Parse()

	if *pattern == "" {
		fmt.Fprintln(os.Stderr, "Error: -e PATTERN is required")
		flag.Usage()
		os.Exit(1)
	}

	re, err := regexp.Compile(*pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex pattern: %v\n", err)
		os.Exit(1)
	}

	files := flag.Args()

	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		grepFile(re, scanner, "", *quiet)
		return
	}

	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mini-grep: %s: %v\n", path, err)
			continue
		}
		scanner := bufio.NewScanner(f)
		source := ""
		if len(files) > 1 {
			source = path
		}
		grepFile(re, scanner, source, *quiet)
		f.Close()
	}
}
