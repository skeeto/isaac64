package isaac64_test

import (
	"testing"

	"nullprogram.com/x/isaac64"
)

func TestLcg128(t *testing.T) {
	// Output from official reference implementation
	want := []uint64{
		0x3f55d36a1b609a5b, 0x6601ce80f1cf7a35, 0x92334584ef40e08b,
		0x0617023b18a2a93e, 0xaae254dc9a92559b, 0x6af1360e8bcbb24a,
		0xb337d3cd3ccf0d2f, 0x35b816475fc147a1, 0x03f664ed80a50776,
		0xd31ea35f9ff8e218, 0xf11183fe6ffee9a2, 0xe3673b3bcf274227,
		0xad84ed35e1bb51b3, 0xcda8fb07c32cd77f, 0xf5d90c644f88b639,
		0xfacca7ebbe19e0a7, 0x9e910c9abcdf788b, 0x6094f6eac27a706a,
		0x4658f1f887095621, 0x2964a037d9c1b7cd, 0x084a81250d0654b1,
		0xe8d38dc859e47007, 0x9b146f60ba0a77fc, 0x2d14d572e62098b9,
		0x43a272e9b2e37de7, 0x45452ff03fa24e1c, 0xe718c93392107227,
	}

	r := isaac64.New()
	r.Seed(0)
	// Shuffle state 64 times
	for i := 0; i < 64*256; i++ {
		r.Uint64()
	}

	for i, w := range want {
		got := r.Uint64()
		if got != w {
			t.Errorf("Rand.Uint64(%d), got %#016x, want %#016x", i, got, w)
		}
	}
}

func BenchmarkIsaac64(b *testing.B) {
	r := isaac64.New()
	r.Seed(int64(b.N))
	for i := 0; i < b.N; i++ {
		r.Uint64()
	}
}
