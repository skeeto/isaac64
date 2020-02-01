package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"nullprogram.com/x/isaac64"
)

func main() {
	const n = 1 << 12
	var buf [8 * n]byte
	r := isaac64.New()
	r.Seed(0)
	for {
		for i := 0; i < n; i++ {
			binary.LittleEndian.PutUint64(buf[i*8:], r.Uint64())
		}
		if _, err := os.Stdout.Write(buf[:]); err != nil {
			fmt.Fprintf(os.Stderr, "fatal: %v\n", err)
			os.Exit(1)
		}
	}
}
