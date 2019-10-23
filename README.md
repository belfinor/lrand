# lrand

## Introdution

*lrand* based on *Mersenne Twister* algorithm. The *Mersenne Twister* is a strong pseudo-random number generator. Strong PRNG has a long period (how many values it generates before repeating itself) and a statistically uniform distribution of values (bits 0 and 1 are equally likely to appear regardless of previous values). A version of the Mersenne Twister available in many programming languages, *MT19937*, has an impressive period of `<i>2<sup>19937</sup>-1</i>`.

## Installation

```
go get "github.com/belfinor/lrand"
```

## Usage

There are two options for using the library: *global context* or *custom context*.

### Custom context

```go
package main

import (

  "fmt"
  "time"

  "github.com/belfinor/lrand"
)

func main() {

  gen := lrand.New()
  gen.Seed(time.Now().UnixNano())


  for i := 0 ; i < 1000 ; i++ {

    // get int64 value
    // or you can call gen.Uint64 to get uint64
    fmt.Println(gen.Int63())
  }
}
```

### Global context

```go
package main

import (

  "fmt"
  "time"

  "github.com/belfinor/lrand"
)

func main() {

  for i := 0 ; i < 1000 ; i++ {
    fmt.Println(lrand.Next()) // return int64
  }
}
```

In this case calls New/Seed are transparent to the developer. Moreover, in this case, the values ​​are precomputed and stored in the channel for optimization.
