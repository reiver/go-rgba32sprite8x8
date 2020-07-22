# go-rgba32sprite8x8

Package **rgba32sprite8x8** provides a type that represents an **8×8 sprite** with RGBA color pixels,
that seamlessly works with Go's built-in `"image"`, `"image/color"`, and `"image/draw"` packages.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-rgba32sprite8x8

[![GoDoc](https://godoc.org/github.com/reiver/go-rgba32sprite8x8?status.svg)](https://godoc.org/github.com/reiver/go-rgba32sprite8x8)


## Example
```go
import (
	"github.com/reiver/go-rgba32sprite8x8"
)

// ...

var memory [rgba32sprite8x8.ByteSize]uint8

sprite := rgba32sprite8x8.Paletted(memory[:])

// ...

draw.Draw(receiver, img.Bounds(), img, image.Bounds().Min, draw.Src)
```

## See Also

Also see:
* https://github.com/reiver/go-dast8x8
* https://github.com/reiver/go-dynaimg
* https://github.com/reiver/go-font8x8
* https://github.com/reiver/go-frame256x288
* https://github.com/reiver/go-imagerelocate
* https://github.com/reiver/go-imgstr
* https://github.com/reiver/go-palette2048
* https://github.com/reiver/go-pel
* https://github.com/reiver/go-rgba32
* https://github.com/reiver/go-sprite8x8
* https://github.com/reiver/go-sprite32x32
* https://github.com/reiver/go-spritesheet8x8x256
* https://github.com/reiver/go-spritesheet32x32x256
* https://github.com/reiver/go-text32x36
