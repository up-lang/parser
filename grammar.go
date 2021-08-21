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
	Members []*NamespaceMember `"{" @@* "}"`
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
	MethodBody    []*Statement           `("{" @@* "}")?`
}

type Parameter struct {
	Name string    `@Ident`
	Type *TypeName `@@`
}

type TypeName struct {
	Nullable  bool       `@"?"?`
	Array     bool       `@"[]"?`
	Namespace *Namespace `(@@ ".")?`
	Name      string     `@Ident`
}

type ObjectName struct {
	Type *TypeName `(@@ ".")?`
	Name string    `@Ident`
}
