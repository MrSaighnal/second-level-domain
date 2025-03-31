// sld.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// Extract second-level domain + TLD using publicsuffix library
func extractSLD(domain string) string {
	effectiveTLDPlusOne, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return domain
	}
	return effectiveTLDPlusOne
}

func main() {
	listFlag := flag.String("l", "", "File with a list of domains")
	flag.Parse()

	var scanner *bufio.Scanner

	if *listFlag != "" {
		file, err := os.Open(*listFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	domainSet := make(map[string]bool)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}
		sld := extractSLD(domain)
		if !domainSet[sld] {
			domainSet[sld] = true
			fmt.Println(sld)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
