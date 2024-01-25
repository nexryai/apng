## APNG golang library
APNG decoder with frame-hook for golang  
Instead of storing every frame of the apng in memory, you can run any function you like on every frame. Forked from [github.com/kettek/apng](https://github.com/kettek/apng)


### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/nexryai/apng"
)

func main() {
	// Open our animated PNG file
	f, err := os.Open("input.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = apng.DecodeAll(f, func(f *apng.FrameHookArgs) error {
		fmt.Printf("Frame %d | Delay:%v\n", f.Num, f.Delay)
		return nil
	})

	if err != nil {
		panic(err)
	}
}

```