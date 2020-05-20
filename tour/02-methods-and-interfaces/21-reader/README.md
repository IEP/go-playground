# Readers

The `io` package specifies the `io.Reader` interface, which represent the read end of a stream data.

The Go standard library contains many implementations of these interfaces, including files, network connections, compressors, chiphers, and others.

The `io.Reader` interface has a `Read` method:

```go
func (T) Read(b []byte) (n int, err error)
```

`Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It return an `io.EOF` error when the stream ends.

The example code creates a `strings.Reader` and consume its output 8 bytes at a time.
