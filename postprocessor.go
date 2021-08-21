package parser

import (
	"errors"
)

func postProcess(parsed *Up) error {
	if parsed == nil {
		return errors.New("root node was nil")
	}

	if parsed.NamespaceDeclarations != nil {
		for _, declaration := range parsed.NamespaceDeclarations {
			err := postProcessNamespaceDecl(declaration)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func postProcessNamespaceDecl(decl *NamespaceDeclaration) error {
	if decl == nil {
		return errors.New("namespace declaration was nil")
	}

	if decl.Name == nil {
		return errors.New("namespace name was nil")
	}
	if len(decl.Name.NamespaceParts) == 0 {
		return errors.New("namespace name was empty")
	}

	if decl.Members != nil {
		for _, member := range decl.Members {
			if member == nil || (member.Class == nil && member.Enum == nil) {
				return errors.New("member was nil")
			}

			if member.Class == nil { // must be enum
				err := postProcessEnum(member.Enum)
				if err != nil {
					return err
				}
			} else { // must be class
				err := postProcessClass(member.Class)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func postProcessEnum(enum *Enum) error {
	if enum == nil {
		return errors.New("enum was nil")
	}

	if enum.Name == "" {
		return errors.New("enum name was empty")
	}
	for _, opt := range enum.Options {
		if opt == "" {
			return errors.New("enum option was empty")
		}
	}

	return nil
}

func postProcessClass(class *Class) error {
	if class == nil {
		return errors.New("class was nil")
	}

	if class.Name == "" {
		return errors.New("class name was empty")
	}

	for _, member := range class.Members {
		if member.Name == "" {
			return errors.New("class member name was empty")
		}

		if member.Type == nil {
			return errors.New("class member type was nil")
		}
		if member.Type.Name == "" {
			return errors.New("class member type was empty")
		}
		if member.MethodBody == nil && member.Type.IsVoid {
			return errors.New("non-method class member had void as its type")
		}

		if member.MethodBody == nil && member.Parameters != nil {
			return errors.New("method \"" + member.Name + "\" was missing a body")
		}

		if member.Parameters != nil {
			for _, param := range member.Parameters {
				if param.Name == "" {
					return errors.New("method parameter had empty name")
				}
				if param.Type == nil {
					return errors.New("method parameter type was nil")
				}
				if param.Type.Name == "" {
					return errors.New("method parameter type was empty")
				}
				if param.Type.IsVoid {
					return errors.New("method parameter had void as its type")
				}
			}
		}

		if member.MethodBody != nil {
			for _, statement := range member.MethodBody {
				err := postProcessStatement(statement)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func postProcessExpression(expr *Expression) error {
	if expr == nil {
		return errors.New("expression was nil")
	}

	if expr.Parts == nil {
		return errors.New("expression was nil")
	}

	for _, part := range expr.Parts {
		err := postProcessExpressionPart(part)
		if err != nil {
			return err
		}
	}

	return nil
}

func postProcessExpressionPart(exprPart *ExpressionPart) error {
	if exprPart == nil {
		return errors.New("expression part was nil")
	}

	if exprPart.Literal == nil && exprPart.ObjAccess == nil && exprPart.Parenthesis == nil && exprPart.Operator == nil &&
		exprPart.Call == nil && exprPart.Construction == nil {
		return errors.New("expression part was nil")
	}

	if exprPart.Literal != nil {
		if exprPart.Literal.Array != nil {
			if exprPart.Literal.Array.Type == nil {
				return errors.New("array literal type was nil")
			}
			if exprPart.Literal.Array.Type.Name == "" {
				return errors.New("array literal type was empty")
			}
			if exprPart.Literal.Array.Type.IsVoid {
				return errors.New("array literal had void as its type")
			}
		}

		if exprPart.Literal.String != nil {
			trimmed := (*exprPart.Literal.String)[1 : len(*exprPart.Literal.String)-1]
			exprPart.Literal.String = &trimmed
		}

		if exprPart.Literal.Char != nil {
			trimmed := (*exprPart.Literal.Char)[1:2]
			exprPart.Literal.Char = &trimmed
		}

	} else if exprPart.Call != nil {
		if exprPart.Call.Name == nil {
			return errors.New("called method name was nil")
		}
		if exprPart.Call.Name.Name == "" {
			return errors.New("called method name was empty")
		}

		if exprPart.Call.Params != nil {
			for _, param := range exprPart.Call.Params {
				err := postProcessExpression(param)
				if err != nil {
					return err
				}
			}
		}

	} else if exprPart.Construction != nil {
		if exprPart.Construction.Type == nil {
			return errors.New("constructor call type was nil")
		}
		if exprPart.Construction.Type.Name == "" {
			return errors.New("constructor call type was empty")
		}
		if exprPart.Construction.Type.IsVoid {
			return errors.New("constructor call had void as its type")
		}

		if exprPart.Construction.Params != nil {
			for _, param := range exprPart.Construction.Params {
				err := postProcessExpression(param)
				if err != nil {
					return err
				}
			}
		}

	} else if exprPart.Parenthesis != nil {
		// recursion!
		err := postProcessExpression(exprPart.Parenthesis)
		if err != nil {
			return err
		}

	} else if exprPart.Operator != nil {
		// meh
	} else if exprPart.ObjAccess != nil {
		if exprPart.ObjAccess.Name == "" {
			return errors.New("accessed object name was empty")
		}
	}

	return nil
}
