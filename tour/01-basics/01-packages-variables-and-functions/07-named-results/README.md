# Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values and known as a "naked" return.

Naked return statements should be used only in short functions. They can harm readability in longer functions.
