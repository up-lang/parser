package parser

import "errors"

// sorry for the code soup that lives in this file, but it is what it is

func postProcessStatement(statement *Statement) (*Statement, error) {
	if statement == nil {
		return nil, errors.New("statement was nil")
	}

	working := *statement

	if working.Assignment == nil && working.VarDef == nil && working.For == nil && working.ForEach == nil &&
		working.If == nil && working.Method == nil && working.Return == nil && working.While == nil {
		return nil, errors.New("statement was nil")
	}

	if working.Return != nil {
		expr, err := postProcessExpression(working.Return)
		if err != nil {
			return nil, err
		}
		working.Return = expr

	} else if working.VarDef != nil {
		if working.VarDef.Name == "" {
			return nil, errors.New("defined variable with no name")
		}
		if working.VarDef.Type == nil {
			return nil, errors.New("defined variable type was nil")
		}
		if working.VarDef.Type.Name == "" {
			return nil, errors.New("defined variable type was empty")
		}
		if working.VarDef.Type.IsVoid {
			return nil, errors.New("defined variable type was void")
		}

		expr, err := postProcessExpression(working.VarDef.ValueToAssign)
		if err != nil {
			return nil, err
		}
		working.VarDef.ValueToAssign = expr

	} else if working.Assignment != nil {
		if working.Assignment.Target == nil {
			return nil, errors.New("assignment target was nil")
		}
		expr, err := postProcessExpression(working.Assignment.ValueToAssign)
		if err != nil {
			return nil, err
		}
		working.Assignment.ValueToAssign = expr

	} else if working.Method != nil {
		if working.Method.Name == nil {
			return nil, errors.New("called method name was nil")
		}
		if working.Method.Name.Name == "" {
			return nil, errors.New("called method name was empty")
		}
		for i, param := range working.Method.Params {
			expr, err := postProcessExpression(param)
			if err != nil {
				return nil, err
			}
			working.Method.Params[i] = expr
		}

	} else if working.If != nil {
		expr, err := postProcessExpression(working.If.Condition)
		if err != nil {
			return nil, err
		}
		working.If.Condition = expr

		// recursion here we go
		if working.If.Body != nil {
			for i, statement := range working.If.Body {
				s, err := postProcessStatement(statement)
				if err != nil {
					return nil, err
				}
				working.If.Body[i] = s
			}
		}

	} else if working.For != nil {
		expr, err := postProcessExpression(working.For.Condition)
		if err != nil {
			return nil, err
		}
		working.If.Condition = expr

		// recursion here we go
		s, err := postProcessStatement(working.For.Increment)
		if err != nil {
			return nil, err
		}
		working.For.Increment = s

		s, err = postProcessStatement(working.For.VarCreation)
		if err != nil {
			return nil, err
		}
		working.For.VarCreation = s

		if working.For.Body != nil {
			for i, statement := range working.For.Body {
				s, err := postProcessStatement(statement)
				if err != nil {
					return nil, err
				}
				working.For.Body[i] = s
			}
		}

	} else if working.ForEach != nil {
		expr, err := postProcessExpression(working.ForEach.Collection)
		if err != nil {
			return nil, err
		}
		working.ForEach.Collection = expr

		if working.ForEach.IndexName == "" {
			return nil, errors.New("foreach index var had empty name")
		}
		if working.ForEach.VarName == "" {
			return nil, errors.New("foreach variable had empty name")
		}

		if working.ForEach.VarType == nil {
			return nil, errors.New("foreach variable type was nil")
		}
		if working.ForEach.VarType.Name == "" {
			return nil, errors.New("foreach variable type was empty")
		}
		if working.ForEach.VarType.IsVoid {
			return nil, errors.New("foreach variable type was void")
		}

		// recursion here we go
		if working.ForEach.Body != nil {
			for i, statement := range working.ForEach.Body {
				s, err := postProcessStatement(statement)
				if err != nil {
					return nil, err
				}
				working.ForEach.Body[i] = s
			}
		}

	} else if working.While != nil {
		expr, err := postProcessExpression(working.While.Condition)
		if err != nil {
			return nil, err
		}
		working.While.Condition = expr

		// recursion here we go
		if working.While.Body != nil {
			for i, statement := range working.While.Body {
				s, err := postProcessStatement(statement)
				if err != nil {
					return nil, err
				}
				working.While.Body[i] = s
			}
		}
	}

	return &working, nil
}
