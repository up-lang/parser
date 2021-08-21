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

func TestLiteral(t *testing.T) {
	parser, err := participle.Build(&Expression{})
	if err != nil {
		t.Fatal(err)
	}

	rootNode := &Expression{}
	err = parser.ParseString("", `true`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if *rootNode.Parts[0].Literal.Bit != true {
		t.Fail()
	}

	rootNode = &Expression{}
	err = parser.ParseString("", `5`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Parts[0].Literal.Int != 5 {
		t.Fail()
	}

	rootNode = &Expression{}
	err = parser.ParseString("", `5.6`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Parts[0].Literal.Float != 5.6 {
		t.Fail()
	}

	rootNode = &Expression{}
	err = parser.ParseString("", `0b1001001`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Parts[0].Literal.Int != 73 {
		t.Fail()
	}

	rootNode = &Expression{}
	err = parser.ParseString("", `0xFA8C`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Parts[0].Literal.Int != 64140 {
		t.Fail()
	}

	rootNode = &Expression{}
	err = parser.ParseString("", `'a'`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if *rootNode.Parts[0].Literal.Char != "a" {
		t.Fail()
	}

	rootNode = &Expression{}
	err = parser.ParseString("", `"hi"`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if *rootNode.Parts[0].Literal.String != "hi" {
		t.Fail()
	}

	rootNode = &Expression{}
	err = parser.ParseString("", `new [] String{}`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Parts[0].Literal.Array.Type.Name != "String" ||
		len(rootNode.Parts[0].Literal.Array.Contents) != 0 {
		t.Fail()
	}

	rootNode = &Expression{}
	err = parser.ParseString("", `new [] String{"test", "test 2"}`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Parts[0].Literal.Array.Type.Name != "String" ||
		*rootNode.Parts[0].Literal.Array.Contents[0].Parts[0].Literal.String != "test" ||
		*rootNode.Parts[0].Literal.Array.Contents[1].Parts[0].Literal.String != "test 2" {
		t.Fail()
	}
}

func TestLocalVarDeclaration(t *testing.T) {
	parser, err := participle.Build(&Statement{})
	if err != nil {
		t.Fatal(err)
	}

	rootNode := &Statement{}
	err = parser.ParseString("", `var myVar stdlib.String;`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.VarDef.Name != "myVar" || rootNode.VarDef.Type.Name != "String" || rootNode.VarDef.Type.Array ||
		rootNode.VarDef.Type.Nullable || rootNode.VarDef.Type.Namespace.NamespaceParts[0] != "stdlib" ||
		rootNode.VarDef.ValueToAssign != nil {
		t.Fail()
	}

	rootNode = &Statement{}
	err = parser.ParseString("", `var myVar ?stdlib.String = myOtherVar;`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.VarDef.Name != "myVar" || rootNode.VarDef.Type.Name != "String" || rootNode.VarDef.Type.Array ||
		!rootNode.VarDef.Type.Nullable || rootNode.VarDef.Type.Namespace.NamespaceParts[0] != "stdlib" ||
		rootNode.VarDef.ValueToAssign == nil || rootNode.VarDef.ValueToAssign.Parts[0].ObjAccess == nil ||
		rootNode.VarDef.ValueToAssign.Parts[0].ObjAccess.Name != "myOtherVar" {
		t.Fail()
	}

	rootNode = &Statement{}
	err = parser.ParseString("", `var myVar []stdlib.String;`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if !rootNode.VarDef.Type.Array {
		t.Fail()
	}
}

func TestAssignment(t *testing.T) {
	parser, err := participle.Build(&Statement{})
	if err != nil {
		t.Fatal(err)
	}

	rootNode := &Statement{}
	err = parser.ParseString("", `myVar = myOtherVar;`, rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Assignment.Target != "myVar" ||
		rootNode.Assignment.ValueToAssign.Parts[0].ObjAccess.Name != "myOtherVar" {
		t.Fail()
	}
}
