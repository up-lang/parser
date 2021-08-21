package parser

import (
	"errors"
)

func postProcess(parsed *Up) (*Up, error) {
	if parsed == nil {
		return nil, errors.New("root node was nil")
	}

	working := *parsed

	if working.NamespaceDeclarations != nil {
		for i, declaration := range working.NamespaceDeclarations {
			decl, err := postProcessNamespaceDecl(declaration)
			if err != nil {
				return nil, err
			}
			working.NamespaceDeclarations[i] = decl
		}
	}

	return &working, nil
}

func postProcessNamespaceDecl(decl *NamespaceDeclaration) (*NamespaceDeclaration, error) {
	if decl == nil {
		return nil, errors.New("namespace declaration was nil")
	}

	working := *decl

	if working.Name == nil {
		return nil, errors.New("namespace name was nil")
	}
	if len(working.Name.NamespaceParts) == 0 {
		return nil, errors.New("namespace name was empty")
	}

	if working.Members != nil {
		for i, member := range working.Members {
			if member == nil || (member.Class == nil && member.Enum == nil) {
				return nil, errors.New("member was nil")
			}

			if member.Class == nil { // must be enum
				enum, err := postProcessEnum(member.Enum)
				if err != nil {
					return nil, err
				}
				working.Members[i].Enum = enum
			} else { // must be class
				class, err := postProcessClass(member.Class)
				if err != nil {
					return nil, err
				}
				working.Members[i].Class = class
			}
		}
	}

	return &working, nil
}

func postProcessEnum(enum *Enum) (*Enum, error) {
	if enum == nil {
		return nil, errors.New("enum was nil")
	}

	working := *enum

	if working.Name == "" {
		return nil, errors.New("enum name was empty")
	}
	for _, opt := range working.Options {
		if opt == "" {
			return nil, errors.New("enum option was empty")
		}
	}

	return &working, nil
}

func postProcessClass(class *Class) (*Class, error) {
	if class == nil {
		return nil, errors.New("class was nil")
	}

	working := *class

	if working.Name == "" {
		return nil, errors.New("class name was empty")
	}

	for i, member := range working.Members {
		if member.Name == "" {
			return nil, errors.New("class member name was empty")
		}

		if member.Type == nil {
			return nil, errors.New("class member type was nil")
		}
		if member.Type.Name == "" {
			return nil, errors.New("class member type was empty")
		}
		if member.MethodBody == nil && member.Type.IsVoid {
			return nil, errors.New("non-method class member had void as its type")
		}

		if member.MethodBody == nil && member.Parameters != nil {
			return nil, errors.New("method \"" + member.Name + "\" was missing a body")
		}

		if member.Parameters != nil {
			for _, param := range member.Parameters {
				if param.Name == "" {
					return nil, errors.New("method parameter had empty name")
				}
				if param.Type == nil {
					return nil, errors.New("method parameter type was nil")
				}
				if param.Type.Name == "" {
					return nil, errors.New("method parameter type was empty")
				}
				if param.Type.IsVoid {
					return nil, errors.New("method parameter had void as its type")
				}
			}
		}

		if member.MethodBody != nil {
			for i2, statement := range member.MethodBody {
				s, err := postProcessStatement(statement)
				if err != nil {
					return nil, err
				}
				working.Members[i].MethodBody[i2] = s
			}
		}
	}

	return &working, nil
}

func postProcessExpression(expr *Expression) (*Expression, error) {
	if expr == nil {
		return nil, errors.New("expression was nil")
	}

	working := *expr

	if working.Parts == nil {
		return nil, errors.New("expression was nil")
	}

	for i, part := range working.Parts {
		p, err := postProcessExpressionPart(part)
		if err != nil {
			return nil, err
		}
		working.Parts[i] = p
	}

	return &working, nil
}

func postProcessExpressionPart(exprPart *ExpressionPart) (*ExpressionPart, error) {
	if exprPart == nil {
		return nil, errors.New("expression part was nil")
	}

	working := *exprPart

	if working.Literal == nil && working.ObjAccess == nil && working.Parenthesis == nil && working.Operator == nil &&
		working.Call == nil || working.Construction == nil {
		return nil, errors.New("expression part was nil")
	}

	if working.Literal != nil {
		if working.Literal.Array != nil {
			if working.Literal.Array.Type == nil {
				return nil, errors.New("array literal type was nil")
			}
			if working.Literal.Array.Type.Name == "" {
				return nil, errors.New("array literal type was empty")
			}
			if working.Literal.Array.Type.IsVoid {
				return nil, errors.New("array literal had void as its type")
			}
		}

		if working.Literal.String != nil {
			trimmed := (*working.Literal.String)[1 : len(*working.Literal.String)-2]
			working.Literal.String = &trimmed
		}

		if working.Literal.Char != nil {
			trimmed := (*working.Literal.String)[1:1]
			working.Literal.Char = &trimmed
		}

	} else if working.Call != nil {
		if working.Call.Name == nil {
			return nil, errors.New("called method name was nil")
		}
		if working.Call.Name.Name == "" {
			return nil, errors.New("called method name was empty")
		}

		if working.Call.Params != nil {
			for i, param := range working.Call.Params {
				expr, err := postProcessExpression(param)
				if err != nil {
					return nil, err
				}
				working.Call.Params[i] = expr
			}
		}

	} else if working.Construction != nil {
		if working.Construction.Type == nil {
			return nil, errors.New("constructor call type was nil")
		}
		if working.Construction.Type.Name == "" {
			return nil, errors.New("constructor call type was empty")
		}
		if working.Construction.Type.IsVoid {
			return nil, errors.New("constructor call had void as its type")
		}

		if working.Construction.Params != nil {
			for i, param := range working.Construction.Params {
				expr, err := postProcessExpression(param)
				if err != nil {
					return nil, err
				}
				working.Construction.Params[i] = expr
			}
		}

	} else if working.Parenthesis != nil {
		// recursion!
		expr, err := postProcessExpression(working.Parenthesis)
		if err != nil {
			return nil, err
		}
		working.Parenthesis = expr

	} else if working.Operator != nil {
		// meh
	} else if working.ObjAccess != nil {
		if working.ObjAccess.Name == "" {
			return nil, errors.New("accessed object name was empty")
		}
	}

	return &working, nil
}
