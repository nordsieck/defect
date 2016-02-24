package defect

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

// Similar to errors.New(), but can be a const
type Error string

var _ error = Error("")

func (e Error) Error() string { return string(e) }

// Asserts that obtained and expected are equal.  If they are not, sets the test as an error,
// prints the optional message, followed by the values of obtained and expected.
func Equal(t *testing.T, obtained, expected interface{}, message ...interface{}) {
	if obtained != expected {
		t.Error(append(message, []interface{}{
			"\r\t",
			lineAndFile(2),
			"obtained:", obtained, "\n",
			"expected:", expected, "\n"}...)...)
	}
}

// Asserts that obtained and expected are equal using reflect.DeepEqual.  If they are not,
// sets the test as an error, prints the optional message, followed by the values of obtained
// and expected.
func DeepEqual(t *testing.T, obtained, expected interface{}, message ...interface{}) {
	if !reflect.DeepEqual(obtained, expected) {
		t.Error(append(message, []interface{}{
			"\r\t",
			lineAndFile(2),
			"obtained:", obtained, "\n",
			"expected:", expected, "\n"}...)...)
	}
}

func lineAndFile(skip int) string {
	if _, file, line, ok := runtime.Caller(skip); ok {
		return filepath.Base(file) + ":" + strconv.Itoa(line) + "\n"
	}
	return ""
}
