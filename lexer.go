
package main

import "fmt"
import "unicode"

var line, checker, size int = 1, 0, 0
var src string

func main() {
	fmt.Println("hho")

	src = "_huehue"
	size = len(src) - 1

	fmt.Println( IsIdentifier() )

}

func CurrentChar() (byte) {
	if checker <= size {
		return src[checker]
	}
	return 0
}

func StringAt(start, end int) (string) {
	if checker <= size {
		return string(src[start:end])
	}
	return ""
}

func IsSingleComment() string {
	if CurrentChar() == '~'  {
		for {
			checker++
			if CurrentChar() == '\n' || checker == size {
				line++
				return "SINGLE-COMMENT"
			}
		}
	}
	return ""
}

func IsMultiComment() string {
	if StringAt(checker, checker + 2) == "::" {
		checker++

		for {
			if CurrentChar() == '\n' {
				line++
			}

			checker++

			if StringAt(checker, checker + 2) == "::" {
				checker++
				return "MULTI-COMMENT"
			}
		}
	}
	return ""
}

func IsIdentifier() string {
	if 	CurrentChar() == '_' || 
		unicode.IsLetter(rune(CurrentChar())) {

		token := string(CurrentChar())
		checker++

		for CurrentChar() == '_' ||
			unicode.IsLetter(rune(CurrentChar())) || 
			unicode.IsNumber(rune(CurrentChar())) {

			token += string(CurrentChar())
			checker++
		}
		checker--
		return string(token)
	}
	return ""
}

// func IsInteger() string {
// 	if 	unicode.IsLetter(rune(CurrentChar())) ||
// 		unicode.IsLetter()
// }

