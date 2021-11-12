# UhOh

Golang advanced error usage with stack tracing

### Usage

```go
// Original golang error
err := errors.New("Something Broke")

// When initiating new uhoh set original error
// Pick the type method to call and set the message you want to convey along with it
ctxErr := uhoh.New(err).General("Sorry something happened on our side")

fmt.Println(ctxErr)
// Output - "General Error: Sorry something happened on our side"
```

### Output

```go

```
