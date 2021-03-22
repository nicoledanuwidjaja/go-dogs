
package main

import (
	"testing"
	"strings"
	"regexp"
	"fmt"
	"strconv"
	"github.com/willf/bitset"
	"sort"
)

var rr bitset.BitSet = parseHTTPCodeRanges("400-499, 500-599")

func BenchmarkFind599(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hasHTTPCodeError(rr, "599")
	}
}

func BenchmarkFind499(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hasHTTPCodeError(rr, "499")
	}
}

func BenchmarkFind400(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hasHTTPCodeError(rr, "400")
	}
}

func BenchmarkFind500(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hasHTTPCodeError(rr, "500")
	}
}

func BenchmarkFind600(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hasHTTPCodeError(rr, "600")
	}
}

func BenchmarkFind450(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hasHTTPCodeError(rr, "450")
	}
}

// parseHTTPCodeRanges parses range pairs and returns a valid slice of HTTP status codes.
func parseHTTPCodeRanges(r string) bitset.BitSet {
	re := regexp.MustCompile("\\d{3}(?:-\\d{3})*(?:,\\d{3}(?:-\\d{3})*)*")
	codes := bitset.BitSet{}
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
			codes.Set(uint(val))
		} else {
			if rg[0] > rg[1] {
				rg[0], rg[1] = rg[1], rg[0]
			}
			min, _ := strconv.Atoi(rg[0])
			max, _ := strconv.Atoi(rg[1])
			for i := min; i <= max; i++ {
				codes.Set(uint(i))
			}
		}
	}
	return codes
}

// hasHTTPCodeError checks if the bitset of HTTP codes contains a given HTTP error code.
func hasHTTPCodeError(b bitset.BitSet, c string) bool {
	code, _ := strconv.Atoi(c)
	return b.Test(uint(code))
}