package parser

import "errors"

// sorry for the code soup that lives in this file, but it is what it is

func postProcessStatement(statement *Statement) error {
	if statement == nil {
		return errors.New("statement was nil")
	}

	if statement.Assignment == nil && statement.VarDef == nil && statement.For == nil && statement.ForEach == nil &&
		statement.If == nil && statement.Method == nil && statement.Return == nil && statement.While == nil {
		return errors.New("statement was nil")
	}

	if statement.Return != nil {
		err := postProcessExpression(statement.Return)
		if err != nil {
			return err
		}

	} else if statement.VarDef != nil {
		if statement.VarDef.Name == "" {
			return errors.New("defined variable with no name")
		}
		if statement.VarDef.Type == nil {
			return errors.New("defined variable type was nil")
		}
		if statement.VarDef.Type.Name == "" {
			return errors.New("defined variable type was empty")
		}
		if statement.VarDef.Type.IsVoid {
			return errors.New("defined variable type was void")
		}

		err := postProcessExpression(statement.VarDef.ValueToAssign)
		if err != nil {
			return err
		}

	} else if statement.Assignment != nil {
		if statement.Assignment.Target == nil {
			return errors.New("assignment target was nil")
		}
		err := postProcessExpression(statement.Assignment.ValueToAssign)
		if err != nil {
			return err
		}

	} else if statement.Method != nil {
		if statement.Method.Name == nil {
			return errors.New("called method name was nil")
		}
		if statement.Method.Name.Name == "" {
			return errors.New("called method name was empty")
		}
		for _, param := range statement.Method.Params {
			err := postProcessExpression(param)
			if err != nil {
				return err
			}
		}

	} else if statement.If != nil {
		err := postProcessExpression(statement.If.Condition)
		if err != nil {
			return err
		}

		// recursion here we go
		if statement.If.Body != nil {
			for _, statement := range statement.If.Body {
				err := postProcessStatement(statement)
				if err != nil {
					return err
				}
			}
		}

	} else if statement.For != nil {
		err := postProcessExpression(statement.For.Condition)
		if err != nil {
			return err
		}

		// recursion here we go
		err = postProcessStatement(statement.For.Increment)
		if err != nil {
			return err
		}

		err = postProcessStatement(statement.For.VarCreation)
		if err != nil {
			return err
		}

		if statement.For.Body != nil {
			for _, statement := range statement.For.Body {
				err := postProcessStatement(statement)
				if err != nil {
					return err
				}
			}
		}

	} else if statement.ForEach != nil {
		err := postProcessExpression(statement.ForEach.Collection)
		if err != nil {
			return err
		}

		if statement.ForEach.IndexName == "" {
			return errors.New("foreach index var had empty name")
		}
		if statement.ForEach.VarName == "" {
			return errors.New("foreach variable had empty name")
		}

		if statement.ForEach.VarType == nil {
			return errors.New("foreach variable type was nil")
		}
		if statement.ForEach.VarType.Name == "" {
			return errors.New("foreach variable type was empty")
		}
		if statement.ForEach.VarType.IsVoid {
			return errors.New("foreach variable type was void")
		}

		// recursion here we go
		if statement.ForEach.Body != nil {
			for _, statement := range statement.ForEach.Body {
				err := postProcessStatement(statement)
				if err != nil {
					return err
				}
			}
		}

	} else if statement.While != nil {
		err := postProcessExpression(statement.While.Condition)
		if err != nil {
			return err
		}

		// recursion here we go
		if statement.While.Body != nil {
			for _, statement := range statement.While.Body {
				err := postProcessStatement(statement)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
