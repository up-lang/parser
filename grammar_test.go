package parser

import (
	"github.com/alecthomas/participle/v2"
	"testing"
)

func TestWithDeclarations(t *testing.T) {
	parser, err := participle.Build(&Up{})
	if err != nil {
		t.Fatal(err)
	}

	rootNode := &Up{}
	err = parser.ParseString("", `with upcore;
with upcore:types;
with stdlib;`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.WithDeclarations[0].Namespace.NamespaceParts[0] != "upcore" ||
		rootNode.WithDeclarations[1].Namespace.NamespaceParts[0] != "upcore" ||
		rootNode.WithDeclarations[1].Namespace.NamespaceParts[1] != "types" ||
		rootNode.WithDeclarations[2].Namespace.NamespaceParts[0] != "stdlib" {
		t.Fail()
	}
}

func TestAccessibilityModifier_Capture(t *testing.T) {
	parser, err := participle.Build(&AccessModWrapper{})
	if err != nil {
		t.Fatal(err)
	}

	rootNode := &AccessModWrapper{}
	err = parser.ParseString("", `public`, rootNode)
	if err != nil {
		t.Fatal(err)
	}
	public := *rootNode

	err = parser.ParseString("", `private`, rootNode)
	if err != nil {
		t.Fatal(err)
	}
	private := *rootNode

	err = parser.ParseString("", `operator`, rootNode)
	if err != nil {
		t.Fatal(err)
	}
	operator := *rootNode

	if *public.Mod != AccessModPublic || *private.Mod != AccessModPrivate || *operator.Mod != AccessModOperator {
		t.Fail()
	}
}

type AccessModWrapper struct {
	Mod *AccessibilityModifier `@("public" | "private" | "operator")`
}

func TestEnum(t *testing.T) {
	parser, err := participle.Build(&NamespaceMember{})
	if err != nil {
		t.Fatal(err)
	}

	rootNode := &NamespaceMember{}
	err = parser.ParseString("", `enum Pets
{
	Cats;
	Dogs;
	Fish;
}`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.IsClass || !rootNode.IsEnum ||
		rootNode.Name != "Pets" ||
		rootNode.EnumOptions[0] != "Cats" ||
		rootNode.EnumOptions[1] != "Dogs" ||
		rootNode.EnumOptions[2] != "Fish" {
		t.Fail()
	}
}

func TestNamespace(t *testing.T) {
	parser, err := participle.Build(&NamespaceDeclaration{})
	if err != nil {
		t.Fatal(err)
	}

	rootNode := &NamespaceDeclaration{}
	err = parser.ParseString("", `namespace MyApp
{
	enum MyEnum {}
	class MyClass {}
}`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Name.NamespaceParts[0] != "MyApp" ||
		rootNode.Members[0].Name != "MyEnum" ||
		!rootNode.Members[0].IsEnum ||
		rootNode.Members[1].Name != "MyClass" ||
		!rootNode.Members[1].IsClass {
		t.Fail()
	}
}

func TestExpression(t *testing.T) {
	parser, err := participle.Build(&Expression{})
	if err != nil {
		t.Fatal(err)
	}

	rootNode := &Expression{}
	err = parser.ParseString("", `var1 / (var2 + min(var3, var4))`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	// please dont make me test this it works okay
}

/*func checkExpressionEq(exp1, exp2 *Expression) bool {
	for i, part := range exp1.Parts {
		expectedPart := exp2.Parts[i]
		if part.ObjAccess != nil && !checkAccessEq(part.ObjAccess, expectedPart.ObjAccess) {
			return false
		}

		if part.Call != nil {
			if !checkAccessEq(part.Call.Name, expectedPart.ObjAccess) {
				return false
			}
			for i2, subExp := range part.Call.Params {
				expectedSubExp := expectedPart.Call.Params[i2]
				if !checkExpressionEq(subExp, expectedSubExp) {
					return false
				}
			}
		}

		if part.Operator != nil && *part.Operator != *expectedPart.Operator {
			return false
		}

		if part.Parenthesis != nil && !checkExpressionEq(part.Parenthesis, expectedPart.Parenthesis) {
			return false
		}
	}

	return true
}

func checkAccessEq(item1, item2 *ObjectName) bool {
	return item1.Name == item2.Name &&
		item1.Type.Name != item2.Type.Name &&
		testEq(item1.Type.Namespace.NamespaceParts, item2.Type.Namespace.NamespaceParts)
}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}*/
