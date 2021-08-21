package parser

type Expression struct {
	Parts []*ExpressionPart `@@+`
}

type ExpressionPart struct {
	Null         bool          `@"null"`
	Literal      *Literal      `|@@`
	Call         *MethodCall   `|@@`
	Construction *Construction `|@@`
	Parenthesis  *Expression   `|"(" @@ ")"`
	Operator     *Operator     `|@("+" | "-" | "*" | "/" | "//" | "%" | "==" | "!=" | ">" | "<" | ">=" | "<=" | "!" | "&" | "&&" | "|" | "||" | "|||" | "===")`
	ObjAccess    *ObjectName   `|@@`
}

type MethodCall struct {
	Name   *ObjectName   `@@`
	Params []*Expression `"(" (@@ ",")* @@ ")"`
}

type Construction struct {
	Type   *TypeName     `"new" @@`
	Params []*Expression `"(" (@@ ",")* @@ ")"`
}
