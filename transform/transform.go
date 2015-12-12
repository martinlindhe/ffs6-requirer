package transform

import (
	"fmt"
	"regexp"
	"strings"
)

func ReplaceRequiresWithEs6Imports(s string) string {

	r1 := regexp.MustCompile(`([a-zA-Z]+) = require\x28'([a-zA-Z]+)'\x29;`)

	r2 := regexp.MustCompile(`([a-zA-Z]+): require\x28'([a-zA-Z.\/]+)'\x29,`)

	header := ""
	res := ""

	arr := strings.Split(s, "\n")
	for _, line := range arr {

		fmt.Printf("processing %s\n", line)

		matches := r1.FindStringSubmatch(line)
		if len(matches) > 1 {
			fmt.Println(matches)
			line = "import " + matches[1] + " from '" + matches[2] + "';"
		} else {

			matches := r2.FindStringSubmatch(line)
			if len(matches) > 1 {
				fmt.Println(matches)

				// get intendation from line
				padding := strings.IndexAny(line, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
				fmt.Println(padding)

				header += "import " + matches[1] + " from '" + matches[2] + "';\n"

				if padding > 0 {
					line = strings.Repeat(" ", padding) + matches[1] + ","
				} else {
					line = matches[1] + ","
				}
			}

		}

		res += line + "\n"
	}

	return header + res
}
