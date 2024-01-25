package apng

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestDecodeAll(t *testing.T) {
	// Open our animated PNG file
	f, err := os.Open("tests/MultipleIDATs.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	a, err := DecodeAll(f, func(f *FrameHookArgs) error {
		fmt.Printf("Frame %d | Delay:%v\n", f.Num, f.Delay)
		return nil
	})

	if err != nil {
		panic(err)
	}

	// Print some information on the APNG
	log.Printf("Found %d frames\n", len(a.Frames))
}
