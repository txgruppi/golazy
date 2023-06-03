# Go Lazy - WIP WIP WIP WIP WIP WIP

A lazy way to process data in Go. No runtime reflection, no interface{}. Generics!

```go
package main

import (
	"fmt"
	"golazy"
	"math"
)

func main() {
	it := golazy.FromRange(0, math.MaxInt, 1)
	it = golazy.Map(it, func(x int) int { return x * x })
	it = golazy.Filter(it, func(x int) bool { return x%2 == 0 })
	it = golazy.Take(it, 10)
	fmt.Printf("%v\n", golazy.ToArray(it))
}
```
