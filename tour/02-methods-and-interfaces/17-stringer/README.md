# Stringers

One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package.

```go
type Stringer interface {
    String() string
}
```

A `Stringer` is a type that can describe itself as a string, the `fmt` package (and many others) look for this interface to print values.
