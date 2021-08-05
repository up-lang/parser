package parser

type Up struct {
	WithDeclarations      []*WithDeclaration      `@@*`
	NamespaceDeclarations []*NamespaceDeclaration `@@*`
}

type WithDeclaration struct {
	Namespace *Namespace `"with" @@ ";"`
}

type Namespace struct {
	NamespaceParts []string `(@Ident ":"?)*`
}

type NamespaceDeclaration struct {
	Name    *Namespace         `"namespace" @@`
	Members []*NamespaceMember `"{" @@ "}"`
}

type NamespaceMember struct {
	IsClass     bool           `@"class"?`
	IsEnum      bool           `@"enum"?`
	Name        string         `@Ident`
	EnumOptions []string       `"{" (@Ident ";")* "}"`
	Members     []*ClassMember `| "{" @@* "}"`
}

type ClassMember struct {
	Accessibility *AccessibilityModifier `@("public" | "private" | "operator")`
	Name          string                 `@Ident`
	Parameters    []*Parameter           `("(" (@@ ",")* @@? ")")?`
	Type          *TypeName              `@@`
	MethodBody    *Code                  `("{" @@ "}")?`
}

type Parameter struct {
	Name string    `@Ident`
	Type *TypeName `@@`
}

type TypeName struct {
	Namespace *Namespace `(@@ ".")?`
	Name      string     `@Ident`
}

type Code struct {
	//TODO: help
}
