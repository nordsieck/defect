# defect
--
    import "github.com/nordsieck/defect"


## Usage

#### func  DeepEqual

```go
func DeepEqual(t *testing.T, obtained, expected interface{}, message ...interface{})
```
Asserts that obtained and expected are equal using reflect.DeepEqual. If they
are not, sets the test as an error, prints the optional message, followed by the
values of obtained and expected.

#### func  Equal

```go
func Equal(t *testing.T, obtained, expected interface{}, message ...interface{})
```
Asserts that obtained and expected are equal. If they are not, sets the test as
an error, prints the optional message, followed by the values of obtained and
expected.

#### type Error

```go
type Error string
```

Similar to errors.New(), but can be a const

#### func (Error) Error

```go
func (e Error) Error() string
```
