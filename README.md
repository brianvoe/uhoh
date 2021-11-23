# UhOh

Golang advanced error usage with stack tracing

uhoh consists of a few parts:

- Original error
- Description of error
- Type error
- Stack trace
  - File
  - Function
  - Line
- Date

## Simple Usage

Uhoh can be as descriptive or as simple as you want. The only thing it requires is an initial error.

```go
import "github.com/brianvoe/uhoh"

// Create new uhoh error
err := uhoh.New(errors.New("original error"))
fmt.Println(err.Error())

// Output:
// 2021-09-12T01:20:30Z | original error
```

## Advanced Usage

Lets see what uhoh can do with a more data and output values.

```go
// Create error
err := New(errors.New("original error"))
err.SetDescribe(errors.New("describe error"))
err.SetType(uhoh.ErrGeneral)

// Output info
fmt.Println(err.Error())
fmt.Println(err.Stack)

// Output:
// 2021-09-12T01:20:30Z | general error | original error | describe error
// [{uhoh_test.go Example 16} {run_example.go runExample 64} {example.go runExamples 44}]
```

## Using With Str

In order to simplify the use of uhoh, uhoh has Str equivalents for set methods so you dont have to always initiate an errors.New().

```go
// Create error
err := uhoh.NewStr("original error")
err.SetDescribeStr("describe error")
err.SetTypeStr("general error")

// Output info
fmt.Println(err.Error())
fmt.Println(err.Stack)

// Output:
// 2021-09-12T01:20:30Z | general error | original error | describe error
// [{uhoh_test.go Example 16} {run_example.go runExample 64} {example.go runExamples 44}]
```

## Real World Usage

To give you a better understanding of how uhoh works, let's look at a real world example

```go
var fileErrType = errors.New("file error")

_, err := os.Open("/test.txt")
if err != nil {
    uhohErr := uhoh.New(err)
    uhohErr.SetDescribe(errors.New("Failed to open file. Please check settings."))
    uhohErr.SetType(fileErrType)

    // Lets log out all the uhoh fields
    log.Printf("%s", uhohErr.ToJson())
}
```

```json
{
  "original": "open /test.txt: no such file or directory",
  "describe": "Failed to open file. Please check settings.",
  "type": "file error",
  "stack": [
    { "file": "uhoh_test.go", "function": "Example_realWorld", "line": 40 },
    { "file": "run_example.go", "function": "runExample", "line": 64 },
    { "file": "example.go", "function": "runExamples", "line": 44 }
  ],
  "date": "2021-09-12T01:20:30Z"
}
```

## Type for Is checking

You may want to check if an error is of a certain type.
Uhoh has a list of types that you can use to check if an error is of a certain type.
Otherwise you can set your own internal types.

```go
var (
    ErrValidation = errors.New("validation error")
    ErrFile       = errors.New("file error")
    ErrNetwork    = errors.New("network error")
)

func main() {
  err := Validate()
  if err != nil {
      if uhoh.Is(err, ErrValidation) {
          // do something
      }
  }
}

func Validate() error {
  return uhoh.New(errors.New("field is invalid")).SetType(ErrValidation)
}

// Output json of everything
log.Println(err.ToJson())
```

## Custom Error

You can create your own Error return string.
By default, uhoh will return with format: date | type error | original error | describe error

```go
// Create your own error return value
customError := func(err *Err) string {
  return err.Original.Error() + ": " + err.Describe.Error()
}

// Set it to your function
SetDefaultErrorFormatter(customError)

originalErr := errors.New("original error")
describeErr := errors.New("describe error")

err := New(originalErr).SetDescribe(describeErr).SetType(uhoh.ErrGeneral)
fmt.Printf("%s", err.Error())

// Output:
// original error: describe error
```

## Set Date

The date is set to the current date and time. If you want to set the date to a specific date, you can use the `SetDate` function.

```go
// Errors
originalErr := errors.New("original error")

err := New(originalErr)
err.SetDate(time.Date(2021, time.Month(9), 12, 1, 20, 30, 0, time.UTC))
fmt.Println(err.Date)

// Output:
// 2021-09-12 01:20:30 +0000 UTC
```
