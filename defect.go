package defect

import (
	"reflect"
	"testing"
)

type Error string

var _ error = Error("")

func (e Error) Error() string { return string(e) }

// Asserts that obtained and expected are equal.  If they are not, sets the test as an error,
// prints the optional message, followed by the values of obtained and expected.
func Equal(t *testing.T, obtained, expected interface{}, message ...interface{}) {
	if obtained != expected {
		t.Error(append(message, []interface{}{
			"\n",
			"obtained:", obtained, "\n",
			"expected:", expected, "\n"}...))
	}
}

// Asserts that obtained and expected are equal using reflect.DeepEqual.  If they are not,
// sets the test as an error, prints the optional message, followed by the values of obtained
// and expected.
func DeepEqual(t *testing.T, obtained, expected interface{}, message ...interface{}) {
	if !reflect.DeepEqual(obtained, expected) {
		t.Error(append(message, []interface{}{
			"\n",
			"obtained:", obtained, "\n",
			"expected:", expected, "\n"}...))
	}
}
