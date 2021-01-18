package main

import (
	"github.com/matryer/is"
	"testing"
)

func TestNumber1(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    123,
				Positive: true,
				Power:    1,
			},
		},
	}
	expr := "123"
	p := ParseExpr(expr)

	is.Equal(should, p)
}

func TestNumber2(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    123,
				Positive: true,
				Power:    1,
			},
		},
		Term{
			Element{
				Type:     Number,
				Value:    45,
				Positive: true,
				Power:    1,
			},
		},
	}
	expr := "123 + 45"
	p := ParseExpr(expr)

	is.Equal(should, p)
}

func TestNumber3(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    123,
				Positive: true,
				Power:    1,
			},
		},
		Term{
			Element{
				Type:     Number,
				Value:    45,
				Positive: false,
				Power:    1,
			},
		},
	}
	expr := "123 - 45"
	p := ParseExpr(expr)

	is.Equal(should, p)
}

func TestMult1(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    2,
				Positive: true,
				Power:    1,
			},
		},
	}
	expr := "1 * 2"
	p := ParseExpr(expr)

	is.Equal(should, p)
}

func TestDiv1(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    2,
				Positive: true,
				Power:    -1,
			},
		},
	}
	expr := "1 / 2"
	p := ParseExpr(expr)

	is.Equal(should, p)
}

func TestVar1(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Variable,
				Value:    0,
				Positive: true,
				Power:    1,
				Name:     "x",
			},
		},
	}
	expr := "x"
	p := ParseExpr(expr)

	is.Equal(should, p)
}

func TestVar2(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Variable,
				Value:    0,
				Positive: true,
				Power:    1,
				Name:     "xy",
			},
		},
	}
	expr := "xy"
	p := ParseExpr(expr)

	is.Equal(should, p)
}

func TestVar3(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Variable,
				Value:    0,
				Positive: true,
				Power:    1,
				Name:     "xy",
			},
			Element{
				Type:     Variable,
				Value:    0,
				Positive: true,
				Power:    -1,
				Name:     "abc",
			},
		},
	}
	expr := "xy / abc"
	p := ParseExpr(expr)

	is.Equal(should, p)
}

func TestDiff1(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    0,
				Positive: true,
				Power:    1,
			},
		},
	}
	expr := "1"
	p := Differentiate(ParseExpr(expr), "x")

	is.Equal(should, p)
}

func TestDiff2(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    0,
				Positive: true,
				Power:    1,
			},
		},
		Term{
			Element{
				Type:     Number,
				Value:    0,
				Positive: true,
				Power:    1,
			},
		},
	}
	expr := "1 + 2"
	p := Differentiate(ParseExpr(expr), "x")

	is.Equal(should, p)
}
