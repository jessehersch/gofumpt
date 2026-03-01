// Copyright (c) 2021, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package format_test

import (
	"testing"

	"github.com/go-quicktest/qt"

	"github.com/jessehersch/gofumpt/format"
)

func TestSourceIncludesSimplify(t *testing.T) {
	t.Parallel()

	in := []byte(`
package p

var ()

func f() {
	for _ = range v {
	}
}
`[1:])
	want := []byte(`
package p

func f() {
	for range v {
	}
}
`[1:])
	got, err := format.Source(in, format.Options{})
	qt.Assert(t, qt.IsNil(err))
	qt.Assert(t, qt.Equals(string(got), string(want)))
}

func TestMultipleLineSignature(t *testing.T) {
	t.Parallel()

	in := []byte(`
package p

func ff(line string) int {
	sum := 0
	for i := 0; i < 68; i++ {
		switch line[i] {
		case '-':
			sum += 1
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sum += int(line[i] - '0')
		}
	}
	return sum
}

var requiredImageOrderFields = []string{
	"latitude_degrees",
	"longitude_degrees",
	"elevation_meters",
	"when_start",
	"when_end",
}

func f(x int,y float64,z string,w []int,u []float64)(int,float64,[]string,error){
	h(t,u,v,w,x,y,z)
	g(x,y,z,w,u,"asdf","xxxx","yyyy",[]string{"this is too long","this is too long","this is too long","this is too long",
		"this is too long","this is too long","this is too long","this is too long"},
		map[string]int{"aaa":10,"bbb":10,"ccc":10,"ddd":10})
}
`[1:])

	want := []byte(`
package p

func ff(line string) int {
	sum := 0
	for i := 0; i < 68; i++ {
		switch line[i] {
		case '-':
			sum += 1
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sum += int(line[i] - '0')
		}
	}
	return sum
}

var requiredImageOrderFields = []string{
	"latitude_degrees",
	"longitude_degrees",
	"elevation_meters",
	"when_start",
	"when_end",
}

func f(
	x int,
	y float64,
	z string,
	w []int,
	u []float64,
) (int, float64, []string, error) {
	h(t, u, v, w, x, y, z)
	g(
		x,
		y,
		z,
		w,
		u,
		"asdf",
		"xxxx",
		"yyyy",
		[]string{
			"this is too long",
			"this is too long",
			"this is too long",
			"this is too long",
			"this is too long",
			"this is too long",
			"this is too long",
			"this is too long",
		},
		map[string]int{"aaa": 10, "bbb": 10, "ccc": 10, "ddd": 10},
	)
}
`[1:])
	got, err := format.Source(in, format.Options{})
	qt.Assert(t, qt.IsNil(err))
	qt.Assert(t, qt.Equals(string(got), string(want)))
}
