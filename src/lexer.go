
package main

import (
	"fmt"
	"regexp"
)

var line, checker, size int = 1, 0, 0
var src string = 
`
'_s = "str"
#ien87 = -87
##dbl_89 = -8.9
?blbltrue = /
?blblfalse = X
^{
}
`

func lex() {

	fmt.Println("Lexer start!")
	size = len(src) - 1

	// identifiers
	re := regexp.MustCompile("\\s+(##|#|'|\\?)[_a-zA-Z][_a-zA-Z0-9]*\\s+")
	fmt.Printf("identifiers: %q\n", re.FindAllString(src, size) )

	// integers
	re = regexp.MustCompile("\\s+-*[0-9]+.[0-9]+\\s+")
	fmt.Printf("integers: %q\n" , re.FindAllString(src, size) )

	// doubles
	re = regexp.MustCompile("\\s+-*[0-9]*\\s+")
	fmt.Printf("doubles: %q\n" , re.FindAllString(src, size) )
	
	// strings
	re = regexp.MustCompile("\\s+\".*\"")
	fmt.Printf("strings: %q\n" , re.FindAllString(src, size) )

	// bools
	re = regexp.MustCompile("\\s+[/X]\\s+")
	fmt.Printf("bools: %q\n" , re.FindAllString(src, size) )

	// first level delimiters: commas, arrays, parenthesis
	re = regexp.MustCompile("(\\[|\\]|\\(|\\)|,)")
	fmt.Printf("first level delimiters: %q\n" , re.FindAllString(src, size) )

	// second level delimiters: logic levels
	re = regexp.MustCompile("(}|\\^{|\\?\\?\\?{|\\?\\?{|\\?{)")
	fmt.Printf("second level delimiters: %q\n" , re.FindAllString(src, size) )

	// assignment operators
	re = regexp.MustCompile("\\s+(\\+|\\+=|-=|\\*=)\\s+")
	fmt.Printf("assignment operators: %q\n" , re.FindAllString(src, size) )

	// increment operators
	re = regexp.MustCompile("\\s+(\\+\\+|--|\\*\\*)\\s+")
	fmt.Printf("increment operators: %q\n" , re.FindAllString(src, size) )

	// arithmetic operators
	re = regexp.MustCompile("\\s+(\\+|-|\\*|/|%|\\^)\\s+")
	fmt.Printf("arithmetic operators: %q\n" , re.FindAllString(src, size) )

	// relational operators
	re = regexp.MustCompile("\\s+(>|<|<=|>=|==|!=)\\s+")
	fmt.Printf("relational operators: %q\n" , re.FindAllString(src, size) )

	// logical operators
	re = regexp.MustCompile("\\s+(\\|\\||&&|!)\\s+")
	fmt.Printf("relational operators: %q\n" , re.FindAllString(src, size) )

	fmt.Println("Lexer ended.")
}
