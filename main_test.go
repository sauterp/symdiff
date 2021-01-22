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
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
		},
	}
	expr := "123"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
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
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
		},
		Term{
			Element{
				Type:     Number,
				Value:    45,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
		},
	}
	expr := "123 + 45"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
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
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
		},
		Term{
			Element{
				Type:     Number,
				Value:    45,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: false,
				Power:    -1,
			},
		},
	}
	expr := "123 - 45"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestMult1(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    2,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
		},
	}
	expr := "1 * 2"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
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

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestDiv2(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    3,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    5,
				Positive: true,
				Power:    -1,
			},
		},
	}
	expr := "3 / 5"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestDiv3(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    3,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    5,
				Positive: true,
				Power:    -1,
			},
		},
	}
	p := SimplifyTerms(should)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestVar1(t *testing.T) {
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
				Value:    1,
				Positive: true,
				Power:    -1,
			},
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

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestVar2(t *testing.T) {
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
				Value:    1,
				Positive: true,
				Power:    -1,
			},
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

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestVar3(t *testing.T) {
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
				Value:    1,
				Positive: true,
				Power:    -1,
			},
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

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestVar4(t *testing.T) {
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
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Value:    0,
				Positive: true,
				Power:    3,
				Name:     "x",
			},
		},
	}
	expr := "x * x * x"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestNumVar1(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    8,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Value:    0,
				Positive: true,
				Power:    1,
				Name:     "x",
			},
		},
	}
	expr := "8*x"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestNumVar2(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    8,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Value:    0,
				Positive: true,
				Power:    1,
				Name:     "x",
			},
		},
	}
	expr := "x*8"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestNumVar3(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    8,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Value:    0,
				Positive: true,
				Power:    1,
				Name:     "x",
			},
		},
	}
	expr := "2*x*2 *  2"
	p := ParseExpr(expr)

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
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

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
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

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestDiffVar1(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    2,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Name:     "x",
				Positive: true,
				Power:    1,
			},
		},
	}
	expr := "x*x"
	parsed := ParseExpr(expr)
	t.Log("parsed: ", parsed)
	p := Differentiate(parsed, "x")

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestDiffVar2(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    2,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Name:     "x",
				Positive: true,
				Power:    1,
			},
		},
		Term{
			Element{
				Type:     Number,
				Value:    8,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Name:     "x",
				Positive: true,
				Power:    3,
			},
		},
	}
	expr := "x*x + 2 * x * x * x * x"
	parsed := ParseExpr(expr)
	t.Log("parsed: ", parsed)
	p := Differentiate(parsed, "x")

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestDiffVar3(t *testing.T) {
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
				Value:    8,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Name:     "x",
				Positive: true,
				Power:    3,
			},
		},
	}
	expr := "3* y * y + 2 * x * x * x * x"
	parsed := ParseExpr(expr)
	t.Log("parsed: ", parsed)
	p := Differentiate(parsed, "x")

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestDiffVar4(t *testing.T) {
	is := is.New(t)

	should := []Term{
		Term{
			Element{
				Type:     Number,
				Value:    6,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Name:     "y",
				Positive: true,
				Power:    2,
			},
			Element{
				Type:     Variable,
				Name:     "x",
				Positive: true,
				Power:    1,
			},
		},
		Term{
			Element{
				Type:     Number,
				Value:    8,
				Positive: true,
				Power:    1,
			},
			Element{
				Type:     Number,
				Value:    1,
				Positive: true,
				Power:    -1,
			},
			Element{
				Type:     Variable,
				Name:     "x",
				Positive: true,
				Power:    3,
			},
		},
	}
	expr := "3* y * y * x * x + 2 * x * x * x * x"
	parsed := ParseExpr(expr)
	t.Log("parsed: ", parsed)
	p := Differentiate(parsed, "x")

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}

func TestDiffVar5(t *testing.T) {
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
	expr := "3* y * y"
	parsed := ParseExpr(expr)
	t.Log("parsed: ", parsed)
	p := Differentiate(parsed, "x")

	t.Log("is:     ", RenderExpr(p))
	t.Log("should: ", RenderExpr(should))
	is.Equal(should, p)
}
