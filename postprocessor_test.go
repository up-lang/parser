package parser

import "testing"

func TestCharStringTrimming(t *testing.T) {
	a := "'a'"
	bc := "\"bc\""
	expr := &Expression{Parts: []*ExpressionPart{
		{Literal: &Literal{Char: &a}},
		{Literal: &Literal{String: &bc}},
	}}

	err := postProcessExpression(expr)
	if err != nil {
		t.Fatal(err)
	}

	if *expr.Parts[0].Literal.Char != "a" ||
		*expr.Parts[1].Literal.String != "bc" {
		t.Fail()
	}
}
