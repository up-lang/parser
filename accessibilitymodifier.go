package parser

import "errors"

type AccessibilityModifier int

const (
	AccessModPublic AccessibilityModifier = iota
	AccessModPrivate
	AccessModOperator
)

func (a *AccessibilityModifier) Capture(values []string) error {
	switch values[0] {
	case "public":
		*a = AccessModPublic
		return nil
	case "private":
		*a = AccessModPrivate
		return nil
	case "operator":
		*a = AccessModOperator
		return nil
	}
	return errors.New("accessibilitymodifier was an invalid value")
}
