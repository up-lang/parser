package parser

type Up struct {
	WithDeclarations      []*WithDeclaration      `@@*`
	NamespaceDeclarations []*NamespaceDeclaration `@@*`
}

type WithDeclaration struct {
	Namespace *Namespace `"with" @@ ";"`
}

type Namespace struct {
	NamespaceParts []string `(@Ident ":")* @Ident`
}

type NamespaceDeclaration struct {
	Name    *Namespace         `"namespace" @@`
	Members []*NamespaceMember `("{" @@* "}"`
	Class   *Class             `|@@`
	Enum    *Enum              `|@@)`
}

type NamespaceMember struct {
	Class *Class `@@`
	Enum  *Enum  `|@@`
}

type Enum struct {
	Name    string   `"enum" @Ident`
	Options []string `"{" ((@Ident ";")* @Ident)? "}"`
}

type Class struct {
	Name    string         `"class" @Ident`
	Members []*ClassMember `"{" (@@ ";")* "}"`
}

type ClassMember struct {
	Accessibility *AccessibilityModifier `@("public" | "private" | "operator")`
	Name          string                 `@Ident`
	Parameters    []*Parameter           `("(" (@@ ","?)* ")")?`
	Type          *TypeName              `@@`
	MethodBody    []*Statement           `("{" @@* "}")?`
}

type Parameter struct {
	Name string    `@Ident`
	Type *TypeName `@@`
}

type TypeName struct {
	IsVoid    bool       `@"void"`
	Nullable  bool       `|(@"?"?`
	Array     bool       `@("[""]")?`
	Namespace *Namespace `(@@ ".")?`
	Name      string     `@Ident)`
}

type ObjectName struct {
	Type *TypeName `(@@ ".")?`
	Name string    `@Ident`
}
