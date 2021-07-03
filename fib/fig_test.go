package fib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sink int

const N = 20


func TestFibWithCache(t *testing.T) {
	table := [] struct {
		arg int
		want int
	}{
		{1, 1},
		{2, 2},
		{10, 55},
	}
	for _, entry := range table {
		got := GetFibWithCache(entry.arg)
		want := entry.want
		assert.Equal(t, want, got)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
}

func BenchmarkGetFibNoCache(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = GetFibNoCache(N)
	}
	sink = res
}

func BenchmarkGetFibWithCache(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = GetFibWithCache(N)
	}
	sink = res
}