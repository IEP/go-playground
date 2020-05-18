# Type conversions

The expression `T(v)` converts the value `v` to the type `T`. This syntax is similar to Python.

Unlike in C, in Go assignment between items of different type requires an explicit conversion. So, for example `var i int = 32.1` won't works.
