# The empty interface

The interface type that specifies zero methods is know as the _empty interface_:

```go
interface{}
```

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interface are used by code that handles values of unknown type. For example, `fmt.Print` takes any argument of type of `interface{}`.
