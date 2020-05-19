# Slice defaults

When slicing, you may omit the high or low bounds to use their defauls instead.

The default is zero for the low bound and the length of the slice for the high bound, like Python.

For the array

```go
var a [10]int
```

these slice expressions are equivalent:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```
