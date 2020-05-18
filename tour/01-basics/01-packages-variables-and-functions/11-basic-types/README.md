# Basic types

Go`s basic types are:

1. `bool`
2. `string`
3. `int`, `int8`, `int16`, `int32`, `int64`
4. `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
5. `byte` alias for `uint8`
6. `rune` alias for `int32`, represents a Unicode code point
7. `float32`, `float64`
8. `complex64`, `complex128`

The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bits systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.
