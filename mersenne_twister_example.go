package mt

import "fmt"

func ExampleMersenneTwisterInt32() {
	key := []uint32{0x123, 0x234, 0x345, 0x456}
	Init_by_array(key)
	fmt.Printf("%d", Genrand_int32())
	// Output:
	// 1067595299
}

func ExampleMersenneTwisterFloat64() {
	key := []uint32{0x123, 0x234, 0x345, 0x456}
	Init_by_array(key)
	fmt.Printf("%.8f", Genrand_real2())
	// Output:
	// 0.76275443
}
