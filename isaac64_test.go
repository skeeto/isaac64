package isaac64_test

import (
	"testing"

	"nullprogram.com/x/isaac64"
)

func TestLcg128(t *testing.T) {
	// Output from official reference implementation
	want := []uint64{
		0x76da1d9489e950e0, 0xc1c5482e8b47ba48,
		0xb734c4e94c6e03fe, 0x281c6b8319ee67f9,
		0x3f58339754ab4c23, 0x030b5b3e26049fe1,
		0xd582f73a4a41be9b, 0x488bdc4908a24b8f,
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
