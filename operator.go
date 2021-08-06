package parser

import "errors"

type Operator int

const (
	OpAdd Operator = iota
	OpSub
	OpMul
	OpDiv
	OpIntDiv
	OpMod
	OpEq
	OpNeq
	OpGt
	OpLt
	OpGtEq
	OpLtEq
	OpNot
	OpAnd
	OpScAnd
	OpOr
	OpScOr
	OpXor
	OpBinEq
)

func (a *Operator) Capture(values []string) error {
	switch values[0] {
	case "+":
		*a = OpAdd
		return nil
	case "-":
		*a = OpSub
		return nil
	case "*":
		*a = OpMul
		return nil
	case "/":
		*a = OpDiv
		return nil
	case "//":
		*a = OpIntDiv
		return nil
	case "%":
		*a = OpMod
		return nil

	case "==":
		*a = OpEq
		return nil
	case "!=":
		*a = OpNeq
		return nil
	case ">":
		*a = OpGt
		return nil
	case "<":
		*a = OpLt
		return nil
	case ">=":
		*a = OpGtEq
		return nil
	case "<=":
		*a = OpLtEq
		return nil

	case "!":
		*a = OpNot
		return nil
	case "&":
		*a = OpAnd
		return nil
	case "&&":
		*a = OpScAnd
		return nil
	case "|":
		*a = OpOr
		return nil
	case "||":
		*a = OpScOr
		return nil
	case "|||":
		*a = OpXor
		return nil
	case "===":
		*a = OpBinEq
		return nil
	}
	return errors.New("operator was an invalid value")
}
