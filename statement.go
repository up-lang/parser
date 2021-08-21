package parser

type LocalVarDefinition struct {
	Name          string      `"var" @Ident`
	Type          *TypeName   `@@`
	ValueToAssign *Expression `("=" @@)?`
}

type Assignment struct {
	Target        *ObjectName `@@ "="`
	ValueToAssign *Expression `@@`
}

type IfStatement struct {
	Condition *Expression  `"if" "(" @@ ")"`
	Body      []*Statement `"{" @@* "}"`
}

type ForLoop struct {
	VarCreation *Statement   `"for" "(" @@ ";"`
	Condition   *Expression  `@@ ";"`
	Increment   *Statement   `@@ ")"`
	Body        []*Statement `"{" @@* "}"`
}

type ForEachLoop struct {
	VarName    string       `"foreach" "(" "var" @Ident`
	VarType    *TypeName    `@@`
	IndexName  string       `"index" @Ident`
	Collection *Expression  `"in" @@ ")"`
	Body       []*Statement `"{" @@* "}"`
}

type WhileLoop struct {
	Condition *Expression  `"while" "(" @@ ")"`
	Body      []*Statement `"{" @@* "}"`
}

type Statement struct {
	Return     *Expression         `"return" @@ ";"`
	VarDef     *LocalVarDefinition `|@@ ";"`
	Assignment *Assignment         `|@@ ";"`
	Method     *MethodCall         `|@@ ";"`
	If         *IfStatement        `|@@ ";"`
	For        *ForLoop            `|@@ ";"`
	ForEach    *ForEachLoop        `|@@ ";"`
	While      *WhileLoop          `|@@ ";"`
}
