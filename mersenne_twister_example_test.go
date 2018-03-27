package mt_test

import (
	"fmt"
	"github.com/nasa9084/go-mersenne-twister"
)

// ExampleMersenneTwisterInt32 is an example of
// 32bit unsigned int random value generation
func ExampleGenrandInt32() {
	key := []uint32{0x123, 0x234, 0x345, 0x456}
	mt.InitByArray(key)
	fmt.Printf("%d", mt.GenrandInt32())
	// Output:
	// 1067595299
}

// ExampleMersenneTwisterFloat64 is an example of
// 32bit real random value generation
func ExampleGenrandReal2() {
	key := []uint32{0x123, 0x234, 0x345, 0x456}
	mt.InitByArray(key)
	fmt.Printf("%.8f", mt.GenrandReal2())
	// Output:
	// 0.24856890
}
