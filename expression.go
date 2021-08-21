package parser

type Expression struct {
	Parts []*ExpressionPart `@@+`
}

type ExpressionPart struct {
	Call        *MethodCall `@@`
	Parenthesis *Expression `|"(" @@ ")"`
	Operator    *Operator   `|@("+" | "-" | "*" | "/" | "//" | "%" | "==" | "!=" | ">" | "<" | ">=" | "<=" | "!" | "&" | "&&" | "|" | "||" | "|||" | "===")`
	ObjAccess   *ObjectName `|@@`
}

type MethodCall struct {
	Name   *ObjectName   `@@`
	Params []*Expression `"(" (@@ ",")* @@ ")"`
}
