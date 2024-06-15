package apng

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestDecodeAll(t *testing.T) {
	// Open our animated PNG file
	f, err := os.Open("tests/blendover.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	a, err := DecodeAll(f, func(f *FrameHookArgs) error {
		fmt.Printf("Frame %d | Delay:%v | xOffset:%d\n", f.Num, f.Delay, f.OffsetX)
		fmt.Printf("BlendOp: %d | DisposeOp: %d\n", f.BlendOp, f.DisposeOp)
		return nil
	})

	if err != nil {
		panic(err)
	}

	// Print some information on the APNG
	log.Printf("Found %d frames\n", len(a.Frames))
}
