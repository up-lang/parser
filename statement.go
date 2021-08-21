package parser

type LocalVarDefinition struct {
	Name          string      `"var" @Ident`
	Type          *TypeName   `@@`
	ValueToAssign *Expression `("=" @@)?`
}

type Assignment struct {
	Target        string      `@Ident "="`
	ValueToAssign *Expression `@@`
}

type Statement struct {
}
