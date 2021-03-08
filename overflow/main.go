package main

import (
	"fmt"
	"math"
)

func main() {
	// Overflow is when there is more data than there are bits available to be stored.
	// Go ONLY catches overflow when it's caught at compile time - it does not catch at runtime.
	var x uint8 = 255 // Maximum uint8 value
	x++               // Overflow trigger.
	fmt.Println(x)    // REturns zero because it just moves to zero when overflowing at runtime.

	// a :=int8(255 + 1) would get caught at compile time because this is compiled rather than done at runtime.

	// Floats overflow to infinity.
	f := float32(math.MaxFloat32)
	fmt.Println(f * 1.2) // Returns +Inf

}
