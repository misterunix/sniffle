// Copyright 2021 by William Jones. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of functions to simplify string operations
package stringy

import (
	"strings"
)

// Center : Center a string with a width, fill, and if the string is padded to the end.
func Center(s string, width int, fill string, fillend bool) string {
	div := (width / 2) - (len(s) / 2)
	var returnString string
	if fillend {
		returnString = strings.Repeat(fill, div) + s + strings.Repeat(fill, div)
	} else {
		returnString = strings.Repeat(fill, div) + s
	}
	return returnString
}

// Right : Right justify a string of width with a fill padding
func Right(s string, width int, fill string) string {
	return strings.Repeat(fill, width-len(s)) + s
}

// Left : Left justify a string of width padding to the end with fill
func Left(s string, width int, fill string) string {
	return s + strings.Repeat(fill, width-len(s))
}

// SplitLines : Split a line with a maximum width splitting lines either at width or whitespace
func SplitLines(line string, width int) []string {
	var retstring []string
	var wrapped string
	words := strings.Fields(line)
	if len(words) == 0 {
		return retstring
	}
	wrapped = words[0]
	spaceleft := width - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceleft {
			retstring = append(retstring, wrapped)
			wrapped = word
			spaceleft = width - len(word)
		} else {
			wrapped += " " + word
			spaceleft -= 1 + len(word)
		}
	}
	retstring = append(retstring, wrapped)
	return retstring
}
