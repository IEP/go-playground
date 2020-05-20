# Exercise: Equivalent Binary Trees

There can be many different binary trees with the same sequence of values store in it. Here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

```txt
     3                    8
    / \                  / \
   /   \                3  13
  1     8              / \
 / \   / \            1   5
1   2 5  13          / \
                    1   2
```

A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

This example uses the `tree` package, which defined the type:

```go
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```
