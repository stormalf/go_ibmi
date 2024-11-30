// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jpeg

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
)

func benchmarkDCT(b *testing.B, f func(*block)) {
	b.StopTimer()
	blocks := make([]block, 0, b.N*len(testBlocks))
	for i := 0; i < b.N; i++ {
		blocks = append(blocks, testBlocks[:]...)
	}
	b.StartTimer()
	for i := range blocks {
		f(&blocks[i])
	}
}

func BenchmarkFDCT(b *testing.B) {
	benchmarkDCT(b, fdct)
}

func BenchmarkIDCT(b *testing.B) {
	benchmarkDCT(b, idct)
}

func TestDCT(t *testing.T) {
	blocks := make([]block, len(testBlocks))
	copy(blocks, testBlocks[:])

	// Append some randomly generated blocks of varying sparseness.
	r := rand.New(rand.NewSource(123))
	for i := 0; i < 100; i++ {
		b := block{}
		n := r.Int() % 64
		for j := 0; j < n; j++ {
			b[r.Int()%len(b)] = r.Int31() % 256
		}
		blocks = append(blocks, b)
	}

	// Check that the FDCT and IDCT functions are inverses, after a scale and
	// level shift. Scaling reduces the rounding errors in the conversion from
	// floats to ints.
	for i, b := range blocks {
		got, want := b, b
		for j := range got {
			got[j] = (got[j] - 128) * 8
		}
		slowFDCT(&got)
		slowIDCT(&got)
		for j := range got {
			got[j] = got[j]/8 + 128
		}
		if differ(&got, &want) {
			t.Errorf("i=%d: IDCT(FDCT)\nsrc\n%s\ngot\n%s\nwant\n%s\n", i, &b, &got, &want)
		}
	}

	// Check that the optimized and slow FDCT implementations agree.
	// The fdct function already does a scale and level shift.
	for i, b := range blocks {
		got, want := b, b
		fdct(&got)
		for j := range want {
			want[j] = (want[j] - 128) * 8
		}
		slowFDCT(&want)
		if differ(&got, &want) {
			t.Errorf("i=%d: FDCT\nsrc\n%s\ngot\n%s\nwant\n%s\n", i, &b, &got, &want)
		}
	}

	// Check that the optimized and slow IDCT implementations agree.
	for i, b := range blocks {
		got, want := b, b
		idct(&got)
		slowIDCT(&want)
		if differ(&got, &want) {
			t.Errorf("i=%d: IDCT\nsrc\n%s\ngot\n%s\nwant\n%s\n", i, &b, &got, &want)
		}
	}
}

// differ reports whether any pair-wise elements in b0 and b1 differ by 2 or
// more. That tolerance is because there isn't a single definitive decoding of
// a given JPEG image, even before the YCbCr to RGB conversion; implementations
// can have different IDCT rounding errors.
func differ(b0, b1 *block) bool {
	for i := range b0 {
		delta := b0[i] - b1[i]
		if delta < -2 || +2 < delta {
			return true
		}
	}
	return false
}

// alpha returns 1 if i is 0 and returns √2 otherwise.
func alpha(i int) float64 {
	if i == 0 {
		return 1
	}
	return math.Sqrt2
}

var cosines = [32]float64{
	+1.0000000000000000000000000000000000000000000000000000000000000000, // cos(π/16 *  0)
	+0.9807852804032304491261822361342390369739337308933360950029160885, // cos(π/16 *  1)
	+0.9238795325112867561281831893967882868224166258636424861150977312, // cos(π/16 *  2)
	+0.8314696123025452370787883776179057567385608119872499634461245902, // cos(π/16 *  3)
	+0.7071067811865475244008443621048490392848359376884740365883398689, // cos(π/16 *  4)
	+0.5555702330196022247428308139485328743749371907548040459241535282, // cos(π/16 *  5)
	+0.3826834323650897717284599840303988667613445624856270414338006356, // cos(π/16 *  6)
	+0.1950903220161282678482848684770222409276916177519548077545020894, // cos(π/16 *  7)

	-0.0000000000000000000000000000000000000000000000000000000000000000, // cos(π/16 *  8)
	-0.1950903220161282678482848684770222409276916177519548077545020894, // cos(π/16 *  9)
	-0.3826834323650897717284599840303988667613445624856270414338006356, // cos(π/16 * 10)
	-0.5555702330196022247428308139485328743749371907548040459241535282, // cos(π/16 * 11)
	-0.7071067811865475244008443621048490392848359376884740365883398689, // cos(π/16 * 12)
	-0.8314696123025452370787883776179057567385608119872499634461245902, // cos(π/16 * 13)
	-0.9238795325112867561281831893967882868224166258636424861150977312, // cos(π/16 * 14)
	-0.9807852804032304491261822361342390369739337308933360950029160885, // cos(π/16 * 15)

	-1.0000000000000000000000000000000000000000000000000000000000000000, // cos(π/16 * 16)
	-0.9807852804032304491261822361342390369739337308933360950029160885, // cos(π/16 * 17)
	-0.9238795325112867561281831893967882868224166258636424861150977312, // cos(π/16 * 18)
	-0.8314696123025452370787883776179057567385608119872499634461245902, // cos(π/16 * 19)
	-0.7071067811865475244008443621048490392848359376884740365883398689, // cos(π/16 * 20)
	-0.5555702330196022247428308139485328743749371907548040459241535282, // cos(π/16 * 21)
	-0.3826834323650897717284599840303988667613445624856270414338006356, // cos(π/16 * 22)
	-0.1950903220161282678482848684770222409276916177519548077545020894, // cos(π/16 * 23)

	+0.0000000000000000000000000000000000000000000000000000000000000000, // cos(π/16 * 24)
	+0.1950903220161282678482848684770222409276916177519548077545020894, // cos(π/16 * 25)
	+0.3826834323650897717284599840303988667613445624856270414338006356, // cos(π/16 * 26)
	+0.5555702330196022247428308139485328743749371907548040459241535282, // cos(π/16 * 27)
	+0.7071067811865475244008443621048490392848359376884740365883398689, // cos(π/16 * 28)
	+0.8314696123025452370787883776179057567385608119872499634461245902, // cos(π/16 * 29)
	+0.9238795325112867561281831893967882868224166258636424861150977312, // cos(π/16 * 30)
	+0.9807852804032304491261822361342390369739337308933360950029160885, // cos(π/16 * 31)
}

// slowFDCT performs the 8*8 2-dimensional forward discrete cosine transform:
//
//	dst[u,v] = (1/8) * Σ_x Σ_y alpha(u) * alpha(v) * src[x,y] *
//		cos((π/2) * (2*x + 1) * u / 8) *
//		cos((π/2) * (2*y + 1) * v / 8)
//
// x and y are in pixel space, and u and v are in transform space.
//
// b acts as both dst and src.
func slowFDCT(b *block) {
	var dst [blockSize]float64
	for v := 0; v < 8; v++ {
		for u := 0; u < 8; u++ {
			sum := 0.0
			for y := 0; y < 8; y++ {
				for x := 0; x < 8; x++ {
					sum += alpha(u) * alpha(v) * float64(b[8*y+x]) *
						cosines[((2*x+1)*u)%32] *
						cosines[((2*y+1)*v)%32]
				}
			}
			dst[8*v+u] = sum / 8
		}
	}
	// Convert from float64 to int32.
	for i := range dst {
		b[i] = int32(dst[i] + 0.5)
	}
}

// slowIDCT performs the 8*8 2-dimensional inverse discrete cosine transform:
//
//	dst[x,y] = (1/8) * Σ_u Σ_v alpha(u) * alpha(v) * src[u,v] *
//		cos((π/2) * (2*x + 1) * u / 8) *
//		cos((π/2) * (2*y + 1) * v / 8)
//
// x and y are in pixel space, and u and v are in transform space.
//
// b acts as both dst and src.
func slowIDCT(b *block) {
	var dst [blockSize]float64
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			sum := 0.0
			for v := 0; v < 8; v++ {
				for u := 0; u < 8; u++ {
					sum += alpha(u) * alpha(v) * float64(b[8*v+u]) *
						cosines[((2*x+1)*u)%32] *
						cosines[((2*y+1)*v)%32]
				}
			}
			dst[8*y+x] = sum / 8
		}
	}
	// Convert from float64 to int32.
	for i := range dst {
		b[i] = int32(dst[i] + 0.5)
	}
}

func (b *block) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "{\n")
	for y := 0; y < 8; y++ {
		fmt.Fprintf(s, "\t")
		for x := 0; x < 8; x++ {
			fmt.Fprintf(s, "0x%04x, ", uint16(b[8*y+x]))
		}
		fmt.Fprintln(s)
	}
	fmt.Fprintf(s, "}")
	return s.String()
}

// testBlocks are the first 10 pre-IDCT blocks from ../testdata/video-001.jpeg.
var testBlocks = [10]block{
	{
		0x7f, 0xf6, 0x01, 0x07, 0xff, 0x00, 0x00, 0x00,
		0xf5, 0x01, 0xfa, 0x01, 0xfe, 0x00, 0x01, 0x00,
		0x05, 0x05, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0xff, 0xf8, 0x00, 0x01, 0xff, 0x00, 0x00,
		0x00, 0x01, 0x00, 0x01, 0x00, 0xff, 0xff, 0x00,
		0xff, 0x0c, 0x00, 0x00, 0x00, 0x00, 0xff, 0x01,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x01, 0xff, 0x01, 0x00, 0xfe,
	},
	{
		0x29, 0x07, 0x00, 0xfc, 0x01, 0x01, 0x00, 0x00,
		0x07, 0x00, 0x03, 0x00, 0x01, 0x00, 0xff, 0xff,
		0xff, 0xfd, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x04, 0x00, 0xff, 0x01, 0x00, 0x00,
		0x01, 0x00, 0x01, 0xff, 0x00, 0x00, 0x00, 0x00,
		0x01, 0xfa, 0x01, 0x00, 0x01, 0x00, 0x01, 0xff,
		0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0x02,
	},
	{
		0xc5, 0xfa, 0x01, 0x00, 0x00, 0x01, 0x00, 0xff,
		0x02, 0xff, 0x01, 0x00, 0x01, 0x00, 0xff, 0x00,
		0xff, 0xff, 0x00, 0xff, 0x01, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x01, 0x00, 0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
		0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	},
	{
		0x86, 0x05, 0x00, 0x02, 0x00, 0x00, 0x01, 0x00,
		0xf2, 0x06, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00,
		0xf6, 0xfa, 0xf9, 0x00, 0xff, 0x01, 0x00, 0x00,
		0xf9, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0xff, 0xff, 0xff, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x01, 0x00, 0xff, 0x01, 0x00,
		0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x01, 0xff, 0x01, 0x00, 0xff, 0x00, 0x00,
	},
	{
		0x24, 0xfe, 0x00, 0xff, 0x00, 0xff, 0xff, 0x00,
		0x08, 0xfd, 0x00, 0x01, 0x01, 0x00, 0x01, 0x00,
		0x06, 0x03, 0x03, 0xff, 0x00, 0x00, 0x00, 0x00,
		0x04, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
		0x01, 0x00, 0x01, 0xff, 0x00, 0x01, 0x00, 0x00,
		0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0xff, 0x01,
	},
	{
		0xcd, 0xff, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01,
		0x03, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0xff,
	},
	{
		0x81, 0xfe, 0x05, 0xff, 0x01, 0xff, 0x01, 0x00,
		0xef, 0xf9, 0x00, 0xf9, 0x00, 0xff, 0x00, 0xff,
		0x05, 0xf9, 0x00, 0xf8, 0x01, 0xff, 0x01, 0xff,
		0x00, 0xff, 0x07, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x01,
		0xff, 0x01, 0x01, 0x00, 0xff, 0x00, 0x00, 0x00,
		0x01, 0x01, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
	},
	{
		0x28, 0x00, 0xfe, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x0b, 0x02, 0x01, 0x03, 0x00, 0xff, 0x00, 0x01,
		0xfe, 0x02, 0x01, 0x03, 0xff, 0x00, 0x00, 0x00,
		0x01, 0x00, 0xfd, 0x00, 0x01, 0x00, 0xff, 0x00,
		0x01, 0xff, 0x00, 0xff, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0xff, 0x01, 0x01, 0x00, 0xff,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x01,
	},
	{
		0xdf, 0xf9, 0xfe, 0x00, 0x03, 0x01, 0xff, 0xff,
		0x04, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0xff, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0xfe, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x01, 0x00, 0x00, 0x00, 0x01,
		0xff, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0xff, 0x00, 0xff, 0x01, 0x00, 0x00, 0x01,
		0xff, 0xff, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	},
	{
		0x88, 0xfd, 0x00, 0x00, 0xff, 0x00, 0x01, 0xff,
		0xe1, 0x06, 0x06, 0x01, 0xff, 0x00, 0x01, 0x00,
		0x08, 0x00, 0xfa, 0x00, 0xff, 0xff, 0xff, 0xff,
		0x08, 0x01, 0x00, 0xff, 0x01, 0xff, 0x00, 0x00,
		0xf5, 0xff, 0x00, 0x01, 0xff, 0x01, 0x01, 0x00,
		0xff, 0xff, 0x01, 0xff, 0x01, 0x00, 0x01, 0x00,
		0x00, 0x01, 0x01, 0xff, 0x00, 0xff, 0x00, 0x01,
		0x02, 0x00, 0x00, 0xff, 0xff, 0x00, 0xff, 0x00,
	},
}