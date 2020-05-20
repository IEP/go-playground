# Images

`Package image` defines the `Image` interface:

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

**Note:** the `Rectangle` return value of the `Bounds` method is actually an `image.Rectangle`, as the declaration is inside package `image`.

The `color.Color` and `color.Model` types are also interfaces, but we'll ignore thath by using the predefined implementations `color.RGBA` and `color.RGBAModel`. These interfaces and types are specified by the `image/color` package.
