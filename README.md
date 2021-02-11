# wfd

A variety of utility functions in go from various ProjectEuler programs.

`$ go get github.com/wfdoran/wfd`

Sample Program

```go
package main

import "github.com/wfdoran/wfd"

func main() {
     wfd_util.Hello()
}
```

## Generate Integer Partitions

```go
package main

import "github.com/wfdoran/wfd"
import "fmt"

func main() {
        for x := range wfd_util.GenPartitions(8) {
                fmt.Println(x)
        }
}
```

## Other


