// Package isaac64 implements the ISAAC64 fast cryptographic random
// number generator.
//
// Ref: https://www.burtleburtle.net/bob/rand/isaacafa.html
package isaac64

import (
	"encoding/binary"
	"io"
	"math/rand"
)

// A Rand is an ISAAC64 random number generator state.
type Rand struct {
	out     [256]uint64
	buf     [256]uint64
	a, b, c uint64
	i       int
}

var _ rand.Source64 = (*Rand)(nil)

// New returns a new, unseeded ISAAC64 random number generator.
func New() *Rand {
	return &Rand{i: 256}
}

// Seed fully initializes the generator state from a seed.
func (r *Rand) Seed(seed int64) {
	for i := 0; i < 256; i++ {
		z := uint64(i) + uint64(seed)
		z ^= z >> 30
		z *= 0xbf58476d1ce4e5b9
		z ^= z >> 27
		z *= 0x94d049bb133111eb
		z ^= z >> 31
		r.buf[i] = z
	}
	r.a = 0
	r.b = 0
	r.c = 0
	r.shuffle()
}

// SeedFrom fills the internal state by reading 2,048 bytes. Typically
// this would be crypto/rand.Reader, but any reader that outputs bytes
// with reasonable statistical quality is suitable.
func (r *Rand) SeedFrom(src io.Reader) error {
	var buf [8 * 256]byte
	if _, err := io.ReadFull(src, buf[:]); err != nil {
		return err
	}
	for i := 0; i < 256; i++ {
		r.buf[i] = binary.LittleEndian.Uint64(buf[i*8:])
	}
	r.a = 0
	r.b = 0
	r.c = 0
	r.shuffle()
	return nil
}

// Uint64 returns the next 64-bit integer from the generator.
func (r *Rand) Uint64() uint64 {
	if r.i == 256 {
		r.shuffle()
	}
	next := r.out[r.i]
	r.i++
	return next
}

// Int63 returns the next 63-bit integer from the generator.
func (r *Rand) Int63() int64 {
	return int64(r.Uint64() >> 1)
}

func (r *Rand) shuffle() {
	r.c++
	r.b += r.c
	for i := 0; i < 256; i += 4 {
		x := r.buf[i+0]
		r.a = ^r.a ^ (r.a << 21)
		r.a += r.buf[(i+0+128)&0xff]
		r.buf[i+0] = r.buf[x>>3&0xff] + r.a + r.b
		r.b = r.buf[r.buf[i+0]>>11&0xff] + x
		r.out[i+0] = r.b

		x = r.buf[i+1]
		r.a ^= r.a >> 5
		r.a += r.buf[(i+1+128)&0xff]
		r.buf[i+1] = r.buf[x>>3&0xff] + r.a + r.b
		r.b = r.buf[r.buf[i+1]>>11&0xff] + x
		r.out[i+1] = r.b

		x = r.buf[i+2]
		r.a ^= r.a << 12
		r.a += r.buf[(i+2+128)&0xff]
		r.buf[i+2] = r.buf[x>>3&0xff] + r.a + r.b
		r.b = r.buf[r.buf[i+2]>>11&0xff] + x
		r.out[i+2] = r.b

		x = r.buf[i+3]
		r.a ^= r.a >> 33
		r.a += r.buf[(i+3+128)&0xff]
		r.buf[i+3] = r.buf[x>>3&0xff] + r.a + r.b
		r.b = r.buf[r.buf[i+3]>>11&0xff] + x
		r.out[i+3] = r.b
	}
	r.i = 0
}
