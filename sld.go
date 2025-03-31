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

func printBanner() {
	fmt.Println("   _____                           ____                  ______                        _     ")
	fmt.Println("  / ___/___  _________  ____  ____/ / /  ___ _   _____  / / __ \\____  ____ ___  ____ _(_)___ ")
	fmt.Println("  \\__ \\/ _ \\/ ___/ __ \\/ __ \\/ __  / /  / _ \\ | / / _ \\/ / / / / __ \\/ __ `__ \\/ __ `/ / __ \\")
	fmt.Println(" ___/ /  __/ /__/ /_/ / / / / /_/ / /__/  __/ |/ /  __/ / /_/ / /_/ / / / / / / /_/ / / / / /")
	fmt.Println("/____/\\___/\\___/\\____/_/ /_/\\__,_/_____|___/|___/\\___/_/_____/_/____/_/ /_/ /_/\\__,_/_/_/ /_/  ")
	fmt.Println("")
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
