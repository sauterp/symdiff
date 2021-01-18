package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("symdiff accepts exactly one argument for the expression you would like to differentiate.")
	}
	expr := os.Args[1]

	ParseExpr(expr)
}

type Type int

const (
	Variable  Type = iota
	Number    Type = iota
	Undefined Type = iota
)

type Element struct {
	// Elements can only be of type Number or Variable
	Type     Type
	Value    int    // Only used with Numbers.
	Name     string // Only used with Variables.
	Power    int    // Only used with Variables.
	Positive bool
}

type Term []Element

func SimplifyTerms(terms []Term) []Term {
	// move all Numbers to the front and multiply them
	// group all Variables with the same name and express them as one power
	// remove 0 powers

	for i := 0; i < len(terms); i++ {
		for j := 0; j < len(terms[i]); j++ {
			e := &terms[i][j]
			// remove Undefined Elements
			if e.Type == Undefined {
				t := terms[i]
				terms[i] = append(t[:j], t[j+1:]...)
				j--
			}
		}
	}

	return terms
}

// ParseExpr parses polynomials expressed like this: "x*x + x + 1"
func ParseExpr(expr string) []Term {

	var terms []Term
	var t Term           // always points to the last Term in terms
	var e *Element = nil // always points to the last Element in t
	newTerm := func() {
		terms = append(terms, Term{
			Element{
				Type:     Undefined,
				Positive: true,
				Power:    1,
			},
		})
		t = terms[len(terms)-1]
		e = &t[len(t)-1]
	}
	newEl := func() {
		terms[len(terms)-1] = append(terms[len(terms)-1],
			Element{
				Type:     Undefined,
				Positive: true,
				Power:    1,
			})
		t = terms[len(terms)-1]
		e = &t[len(t)-1]
	}
	newTerm()
	for _, c := range expr {
		switch c {
		case ' ', '\t', '\n':
			// whitespace is skipped
		case '+':
			// start a new Term
			newTerm()
			e.Positive = true
		case '-':
			// start a new negative Term
			newTerm()
			e.Positive = false
		case '*':
			// extend Term
			newEl()
			e.Power = 1
		case '/':
			// extend Term
			newEl()
			e.Power = -1
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			// extend Term
			if e.Type == Variable {
				e.Name += string(c)
			} else if e.Type == Undefined {
				e.Type = Variable
				e.Name = string(c)
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			// extend Term
			v, _ := strconv.Atoi(string(c))
			if e.Type == Number {
				e.Value = e.Value*10 + v
			} else if e.Type == Variable {
				newEl()
				e.Type = Number
				e.Value = v
			} else if e.Type == Undefined {
				e.Type = Number
				e.Value = v
			}
		default:
			fmt.Printf("unrecognized character: %s\n", string(c))
			os.Exit(1)
		}
	}

	return SimplifyTerms(terms)
}

// Differentiate differentiates expr by v
func Differentiate(expr []Term, v string) []Term {
	for i := 0; i < len(expr); i++ {
		isPureNumeric := true
		for j := 0; j < len(expr[i]); j++ {
			e := &expr[i][j]
			if e.Type == Undefined {
				fmt.Println("undefined elements in expression")
				os.Exit(1)
			} else if e.Type != Number {
				isPureNumeric = false
			}
		}
		if isPureNumeric {
			// We assume, that the numeric Elements are always at [0] in a Term
			expr[i][0].Value = 0
		}
	}
	return expr
}
