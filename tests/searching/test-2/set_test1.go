
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

var rr []int = parseHTTPCodeRanges("400-499, 500-599")

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
func parseHTTPCodeRanges(r string) []int {
	re := regexp.MustCompile("\\d{3}(?:-\\d{3})*(?:,\\d{3}(?:-\\d{3})*)*")
	codes := []int{}
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
func appendOrdered(s []int, n int) []int {
	sort.Ints(s)
	i := sort.SearchInts(s, n)
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = n
	return s
}

// hasHTTPCodeError checks if the slice of codes contains a given HTTP error code.
func hasHTTPCodeError(b []int, c string) bool {
	code, _ := strconv.Atoi(c)
	for _, v := range b {
		if v == code {
			return true
		}
	}
	return false
}