package main

import "fmt"

var (
	pub = []uint64{3418282, 8719412}
	sn  = uint64(7)
	mod = uint64(20201227)
)

func main() {

	ls := 0

	r := uint64(1)

	for r != pub[0] && r != pub[1] {
		r = (r * sn) % mod
		ls++
	}

	loopSize := ls
	fmt.Println(loopSize)

	nsn := pub[0]
	if r == nsn {
		nsn = pub[1]
	}

	r = uint64(1)

	for ls > 0 {
		r = (r * nsn) % mod
		ls--
	}

	fmt.Println(r)
}
