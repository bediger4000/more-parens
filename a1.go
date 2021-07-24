package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Args[1]
	var stack []rune

	inputs := []rune(str)

	tryit(inputs, stack)
}

var possibilities []rune = []rune{')', '('}

// tryit recurses on inputs slice: it's one less in length each
// level of recursion, so validate when length is zero.
// stack either increases by input[0] (if it's not '*'),
// or by a candidate '(' or ')' if it input[0] holds '*'.
// stack may stay same length for input[0] == '*', too:
// empty string case.
func tryit(inputs, stack []rune) {
	if len(inputs) == 0 {
		validate(stack)
		return
	}

	if inputs[0] == '*' {
		// try parentheses possibilities
		for _, p := range possibilities {
			stack = append(stack, p)
			tryit(inputs[1:], stack)
			stack = stack[:len(stack)-1]
		}
		// try empty string
		tryit(inputs[1:], stack)
		return
	}
	stack = append(stack, inputs[0])
	tryit(inputs[1:], stack)
	stack = stack[:len(stack)-1]
}

// validate the substituted-in stack of parens:
// Parity of '(' as 1, ')' as -1. If parity ever goes below 0
// it's an unbalanced set of parens. If parity is not zero after
// checking the whole stack, it's unbalanced.
func validate(stack []rune) {

	parity := 0
	for _, r := range stack {
		switch r {
		case '(':
			parity++
		case ')':
			parity--
		default:
			// stack should only contain '(' and ')'
			fmt.Printf("Something wrong: %q\n", string(stack))
			os.Exit(1)
		}
		if parity < 0 {
			return
		}
	}
	if parity == 0 {
		fmt.Printf("Balanced expression: %q\n", string(stack))
	}
}
