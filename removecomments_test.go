package parser

import "testing"

func TestRemoveComments(t *testing.T) {
	raw := `stuff is here yeah ok~oop this shouldnt be here now
this will be visible
~~~
were now in a block comment
so none of this is shown
~ ~~~ ~ this wont be here though
line comments will not override closing block comments only opening`

	expected := `stuff is here yeah ok
this will be visible
 
line comments will not override closing block comments only opening`

	actual := RemoveComments(raw)

	if actual != expected {
		t.Fail()
	}
}
