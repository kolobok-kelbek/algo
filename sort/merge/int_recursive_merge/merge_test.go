package int_recursive_merge

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRecursiveMerge(t *testing.T) {
	cases := []struct {
		data []int
		want []int
	}{
		{
			data: nil,
			want: []int{},
		},
		{
			data: []int{},
			want: []int{},
		},
		{
			data: []int{5, 4, 3, 7, 8, 6, 1, 9, 2},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			data: []int{5, -3, 4, 3, 0, 7, -2, 8, 6, 1, 9, 2, -1},
			want: []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			data: []int{5, -3, 4, -3, 3, 0, 7, 7, -2, 8, 6, 1, 5, 9, 7, 2, -1, 2, 5},
			want: []int{-3, -3, -2, -1, 0, 1, 2, 2, 3, 4, 5, 5, 5, 6, 7, 7, 7, 8, 9},
		},
	}

	for i, c := range cases {
		num := strconv.Itoa(i)
		t.Run("case №"+num, func(t *testing.T) {
			got := RecursiveMerge(c.data)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestMerge(t *testing.T) {

	cases := []struct {
		data1 []int
		data2 []int
		want  []int
	}{
		{
			data1: []int{1, 2},
			data2: []int{3, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			data1: []int{1, 3, 5, 7, 8, 9, 10},
			data2: []int{2, 4, 6},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for i, c := range cases {
		num := strconv.Itoa(i)
		t.Run("case №"+num, func(t *testing.T) {
			got := merge(c.data1, c.data2)
			assert.Equal(t, c.want, got)
		})
	}
}
