
package main

import (
	"testing"
	"strings"
	"regexp"
	"fmt"
	"strconv"
)

func BenchmarkParsingClient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseHTTPCodeRanges("400-499")
	}
}

func BenchmarkParsingServer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseHTTPCodeRanges("500-599")
	}
}

func BenchmarkParsingClientAndServer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseHTTPCodeRanges("400-499")
		parseHTTPCodeRanges("500-599")
	}
}

func BenchmarkParsingFlipperooClient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseHTTPCodeRanges("499-400")
	}
}

func BenchmarkParsingMergeRangesClient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseHTTPCodeRanges("400-499, 450-499, 400-449")
	}
}

// heavyloading
func BenchmarkParsingMergeRangesClientAndServer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseHTTPCodeRanges("400-499, 450-499, 400-449")
		parseHTTPCodeRanges("500-599, 550-599, 500-549")
	}
}

func BenchmarkParsingFlipperooAndMergeClient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseHTTPCodeRanges("499-400,400-450")
	}
}

func BenchmarkParsingClientAndInvalid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseHTTPCodeRanges("400-499,300,1000")
	}
}

type codeRange struct {
	min int
	max int
}

// parseHTTPCodeRanges parses range pairs and returns a valid slice of HTTP status codes.
func parseHTTPCodeRanges(r string) codeRange {
	re := regexp.MustCompile("\\d{3}(?:-\\d{3})*(?:,\\d{3}(?:-\\d{3})*)*")
	codes := codeRange{}
	for _, code := range strings.Split(r, ",") {
		code = strings.TrimSpace(code)
		if code == "" {
			continue
		}
		if !re.MatchString(code) {
			fmt.Println("Invalid range for: ", code)
			continue
		}
		rg := strings.Split(code, "-")
		if len(rg) == 1 {
			val, _ := strconv.Atoi(rg[0])
			codes = appendOrdered(codes, val)
		} else {
			if rg[0] > rg[1] {
				rg[0], rg[1] = rg[1], rg[0]
			}
			min, _ := strconv.Atoi(rg[0])
			max, _ := strconv.Atoi(rg[1])
			for i := min; i <= max; i++ {
				codes = appendOrdered(codes, i)
			}
		}
	}
	return codes
}

// appendOrdered appends n to the slice s at the necessary index such that the resulting slice is kept ordered.
func appendOrdered(codes codeRange, c int) codeRange {
	if c < codes.min || codes.min == 0 {
		codes.min = c
	} 
	
	if c > codes.max || codes.max == 0 {
		codes.max = c
	}
	return codes
}