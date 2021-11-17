# UhOh

Golang advanced error usage with stack tracing

uhoh consists of 3 parts:

- Original error
- Description of error
- Stack trace
  - File
  - Function
  - Line

## Usage

```go
import "github.com/brianvoe/uhoh"

// Errors
originalErr := errors.New("original error")
describeErr := errors.New("describe error")

// Create error
err := uhoh.New(originalErr, describeErr)

// Output info
fmt.Println(err.Error()) // Will prioritize describe error
fmt.Println(err.Original())
fmt.Println(err.Describe())
fmt.Println(err.File())
fmt.Println(err.Function())
fmt.Println(err.Line())

// Output:
// describe error
// original error
// describe error
// uhoh_test.go
// Example
// 16
```

## Real world usage

To give you a better understanding of how uhoh works, let's look at a real world example

```go
_, err := os.Open("/test.txt")
if err != nil {
    err = uhoh.New(err, errors.New("Failed to open file. Please check settings."))

    fmt.Printf("%s", err.ToJson())
}

```

```json
{
  "original": "open /test.txt: no such file or directory",
  "describe": "Failed to open file. Please check settings.",
  "file": "file.go",
  "function": "OpenFile",
  "line": 13,
  "date": "2021-09-12T01:10:30Z"
}
```

## Output

```go
originalErr := errors.New("original error")
describeErr := errors.New("describe error")

err := New(originalErr, describeErr)
fmt.Printf("%s", err.ToJson())

// Output:
// {"date":"2021-09-12T01:10:30Z","describe":"describe error","file":"file.go","function":"ExampleErr_ToJson","line":13,"original":"original error"}
```

## No Describe Error

You may not want to include a description error for the reason that the original error is the best description.

```go
// Errors
originalErr := errors.New("original error")

// Create error
err := uhoh.New(originalErr, nil)

// Output info
fmt.Println(err.Error())
fmt.Println(err.File())
fmt.Println(err.Function())
fmt.Println(err.Line())

// Output:
// original error
// uhoh_test.go
// Example
// 16
```

## Set Date

The date is set to the current date and time. If you want to set the date to a specific date, you can use the `SetDate` function.

```go
// Errors
originalErr := errors.New("original error")

err := New(originalErr, nil)
err.SetDate(time.Date(2021, time.Month(9), 12, 1, 20, 30, 0, time.UTC))
fmt.Println(err.Date())

// Output:
// 2021-09-12 01:20:30 +0000 UTC
```
