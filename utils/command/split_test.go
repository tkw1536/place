package command

import (
	"fmt"
	"reflect"
	"testing"
)

func arrayToString(strings []string) string {
	joined := "[ "

	for _, s := range strings {
		joined += fmt.Sprintf("%q ", s)
	}

	joined += "]"

	return joined
}

// assert that the splitting of input succeeds
func assertSplitSuccess(input string, expected []string, t *testing.T) {
	actual, err := Split(input)
	if err != nil {
		t.Errorf("Split(%q) failed: %s", input, err.Error())
		return
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Split(%q) failed: Expected %s, got %s", input, arrayToString(expected), arrayToString(actual))
		return
	}
}

// assert that the splitting of input fails
func assertSplitFailure(input string, t *testing.T) {
	_, err := Split(input)
	if err == nil {
		t.Errorf("Splitting command %q succeeded, but was expecting failure. ", input)
	}
}

func TestSplitNormal(t *testing.T) {
	assertSplitSuccess("a b c", []string{"a", "b", "c"}, t)
}

func TestSplitDouble(t *testing.T) {
	assertSplitSuccess("\"a b\" c", []string{"a b", "c"}, t)
}

func TestSplitSingle(t *testing.T) {
	assertSplitSuccess("'a b' c", []string{"a b", "c"}, t)
}

func TestSplitEscape(t *testing.T) {
	assertSplitSuccess("a\\ b c", []string{"a b", "c"}, t)
}

func TestSplitEscapeDouble(t *testing.T) {
	assertSplitSuccess("\"a\\\"b\" c", []string{"a\"b", "c"}, t)
}

func TestSplitEscapeSingle(t *testing.T) {
	assertSplitSuccess("'a\\'b' c", []string{"a'b", "c"}, t)
}

func TestSplitEscapeFail(t *testing.T) {
	assertSplitFailure("a\\", t)
}

func TestSplitQuoteFailSingle(t *testing.T) {
	assertSplitFailure("' a", t)
}

func TestSplitQuoteFailDouble(t *testing.T) {
	assertSplitFailure("\" a", t)
}
