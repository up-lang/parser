package parser

type ArrayLiteral struct {
	Type     *TypeName     `"new" "[""]" @@`
	Contents []*Expression `"{" (@@ ",")* "}"`
}

type Literal struct {
	Array  *ArrayLiteral `@@`
	Bit    *Bit          `|@("1" | "0" | "true" | "false")`
	Float  float64       `|@Float`
	Int    int           `|@Int`
	Char   *string       `|"'" @Char "'"`
	String *string       `|"\"" @String "\""`
}
