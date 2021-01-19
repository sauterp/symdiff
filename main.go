package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("symdiff accepts exactly two arguments, the expression you would like to differentiate followed by the variable you would like to differentiate by.")
	}
	expr := os.Args[1]

	fmt.Println(RenderExpr(Differentiate(ParseExpr(expr), os.Args[2])))
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

	// we need one number Number and one V
	var r []Term
	for i := 0; i < len(terms); i++ {
		// collect all numbers
		numberEl := Element{
			Type:     Undefined,
			Positive: true,
			Power:    1,
		}
		// collect all variables with the same name
		vars := make(map[string]Element)
		for j := 0; j < len(terms[i]); j++ {
			e := &terms[i][j]
			switch e.Type {
			case Undefined:
				// remove Undefined Elements
				t := terms[i]
				terms[i] = append(t[:j], t[j+1:]...)
				j--
			case Variable:
				vEl, ok := vars[e.Name]
				if ok {
					vEl.Power += e.Power
					vars[e.Name] = vEl
				} else {
					vars[e.Name] = Element{
						Type:     Variable,
						Name:     e.Name,
						Power:    e.Power,
						Positive: true,
					}
				}
			case Number:
				if numberEl.Type == Undefined {
					numberEl.Type = Number
					numberEl.Value = e.Value
					numberEl.Positive = e.Positive
					numberEl.Power = e.Power
				} else {
					// TODO support fractions
					numberEl.Value *= e.Value
					if numberEl.Positive == e.Positive {
						numberEl.Positive = true
					} else {
						numberEl.Positive = false
					}
				}
			}
		}

		var newTerm Term
		if numberEl.Type == Undefined {
			newTerm = append(newTerm, Element{
				Type:     Number,
				Positive: true,
				Power:    1,
				Value:    1,
			})
		} else {
			newTerm = append(newTerm, numberEl)
		}
		for _, v := range vars {
			newTerm = append(newTerm, v)
		}
		r = append(r, newTerm)
	}

	return r
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
		noDiffVar := true
		for j := 0; j < len(expr[i]); j++ {
			e := &expr[i][j]
			if e.Type == Undefined {
				fmt.Println("undefined elements in expression")
				os.Exit(1)
			} else if e.Type == Variable {
				if e.Name == v {
					// We assume, there is always a Number Element at expr[i][0], this is ensured by SimplifyTerms
					expr[i][0].Value *= e.Power
					expr[i][j].Power--
					noDiffVar = false
				}
			}
		}
		if noDiffVar {
			// We assume, that the numeric Elements are always at [0] in a Term
			rhTerms := expr[i+1:]
			expr = append(expr[:i], Term{
				Element{
					Type:     Number,
					Value:    0,
					Power:    1,
					Positive: true,
				},
			})
			expr = append(expr[:i+1], rhTerms...)
		}
	}
	return expr
}

func RenderExpr(terms []Term) string {
	r := ""
	for i := 0; i < len(terms); i++ {
		for j := 0; j < len(terms[i]); j++ {
			e := &terms[i][j]
			switch e.Type {
			case Variable:
				r += fmt.Sprintf(" (%s^%d)", e.Name, e.Power)
			case Number:
				var op string
				if e.Positive {
					op = "+"
				} else {
					op = "-"
				}
				r += fmt.Sprintf(" %s%d", op, e.Value)
			}
		}
	}
	return r
}
