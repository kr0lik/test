package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// var r = regexp.MustCompile(`\(([a-z)]*)\)\[(\d+)\]`)
var r = regexp.MustCompile(`\(([a-z)]*)\)\[(\d*)\]`)

func bracketGrammar(s string) string {
	for {
		subMatches := r.FindStringSubmatch(s)
		if len(subMatches) == 0 {
			break
		}

		mainString := subMatches[0]
		repeatableString := subMatches[1]
		countStr := subMatches[2]

		if countStr != "" {
			count, _ := strconv.Atoi(countStr)
			repeatableString = strings.Repeat(repeatableString, count)
		}

		s = strings.Replace(s, mainString, repeatableString, 1)
	}

	return s
}

func main() {
	testsCases := [][2]string{
		[2]string{"", ""},
		[2]string{"ab", "ab"},
		[2]string{"(ab)[3]", "ababab"},
		[2]string{"((ab)[2])[2]", "abababab"},
		[2]string{"(()[1])[2]", ""},
		// ==========
		[2]string{"(a)[0]bc", "bc"},
		[2]string{"(a)[2](b)[2]", "aabb"},
		[2]string{"((a)[2]b)[3]", "aabaabaab"},
		[2]string{"abc(d)[2]", "abcdd"},
		//========== Cases added by me
		[2]string{"()[]", ""},
		[2]string{"(ab)[]", "ab"},
		[2]string{"(ab)[0]", ""},
	}
	for _, tc := range testsCases {
		in, out := tc[0], tc[1]
		fmt.Println(in, "->", bracketGrammar(in), "|", out, "|", reflect.DeepEqual(bracketGrammar(in), out))
	}
}
