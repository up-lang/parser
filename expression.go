package parser

type Expression struct {
	Parts []*ExpressionPart `@@+`
}

type ExpressionPart struct {
	ObjAccess *ObjectName    `@@`
	Call      *MethodCall    `| @@`
	OpChain   []*OperatorUse `| @@+`
}

type OperatorUse struct {
	FirstObj  *ObjectName `@@`
	Operator  *Operator   `@@?`
	SecondObj *ObjectName `@@?`
}

type MethodCall struct {
	Name   *ObjectName   `@@`
	Params []*Expression `"(" (@@ ",")* ")"`
}
