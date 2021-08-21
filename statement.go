package parser

type LocalVarDefinition struct {
	Name          string      `"var" @Ident`
	Type          *TypeName   `@@`
	ValueToAssign *Expression `("=" @@)?`
}

type Assignment struct {
	Name          string      `@Ident "="`
	ValueToAssign *Expression `@@`
}

type Statement struct {
}
