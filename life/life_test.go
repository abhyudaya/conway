package life

import "testing"
import "reflect"

func makeEmptyGridOfSize(inp [][]uint8) [][]uint8 {
      grid := make([][]uint8, len(inp))
        for i := range inp {
                grid[i] = make([]uint8, len(inp[i]))
        }
        return grid

}

func TestRules(t *testing.T) {
	cases := []struct {
		inp [][]uint8
		want [][]uint8
	} {
		{
			[][]uint8 {
            	[]uint8{0, 0, 0, 0, 0},
            	[]uint8{0, 0, 1, 0, 0},
            	[]uint8{0, 0, 1, 0, 0},
            	[]uint8{0, 0, 1, 0, 0},
            	[]uint8{0, 0, 0, 0, 0},
        	},
			[][]uint8 {
            	[]uint8{0, 0, 0, 0, 0},
            	[]uint8{0, 0, 0, 0, 0},
            	[]uint8{0, 1, 1, 1, 0},
            	[]uint8{0, 0, 0, 0, 0},
            	[]uint8{0, 0, 0, 0, 0},
        	},
		},
		{
			[][]uint8 {
            	[]uint8{0, 0, 0, 0, 0},
            	[]uint8{0, 0, 0, 0, 0},
            	[]uint8{0, 1, 1, 1, 0},
            	[]uint8{0, 0, 0, 0, 0},
            	[]uint8{0, 0, 0, 0, 0},
        	},
			[][]uint8 {
            	[]uint8{0, 0, 0, 0, 0},
            	[]uint8{0, 0, 1, 0, 0},
            	[]uint8{0, 0, 1, 0, 0},
            	[]uint8{0, 0, 1, 0, 0},
            	[]uint8{0, 0, 0, 0, 0},
        	},
		},
		{
			[][]uint8 {
            	[]uint8{0, 0, 0, 0},
            	[]uint8{0, 1, 1, 0},
            	[]uint8{0, 1, 1, 0},
            	[]uint8{0, 0, 0, 0},
        	},
			[][]uint8 {
            	[]uint8{0, 0, 0, 0},
            	[]uint8{0, 1, 1, 0},
            	[]uint8{0, 1, 1, 0},
            	[]uint8{0, 0, 0, 0},
        	},
		},
	}

	for _, c := range cases {
		got := makeEmptyGridOfSize(c.inp)
		if got = NextLife(c.inp, got); !reflect.DeepEqual(got, c.want) {
			t.Errorf("NextLife failed: input = %v, got = %v, expected = %v", c.inp, got, c.want)
		}
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		grid [][]uint8
		in_x int
		in_y int
		want int
	}{
		{[][]uint8{
			[]uint8{0, 0, 0},
			[]uint8{0, 0, 0},
			[]uint8{0, 0, 0},
		}, 0, 0, 0},
		{[][]uint8{
			[]uint8{1, 1, 1, 1},
			[]uint8{0, 1, 1, 1},
			[]uint8{0, 0, 1, 1},
			[]uint8{0, 0, 0, 1},
		}, 0, 0, 2},
		{[][]uint8{
			[]uint8{1, 1, 1, 1},
			[]uint8{0, 1, 1, 1},
			[]uint8{0, 0, 1, 1},
			[]uint8{0, 0, 0, 1},
		}, 0, 1, 4},
		{[][]uint8{
			[]uint8{1, 1, 1, 1},
			[]uint8{0, 1, 1, 1},
			[]uint8{0, 0, 1, 1},
			[]uint8{0, 0, 0, 1},
		}, 1, 0, 3},
		{[][]uint8{
			[]uint8{1, 1, 1, 1},
			[]uint8{0, 1, 1, 1},
			[]uint8{0, 0, 1, 1},
			[]uint8{0, 0, 0, 1},
		}, 1, 1, 5},
		{[][]uint8{
			[]uint8{1, 1, 1, 1},
			[]uint8{0, 1, 1, 1},
			[]uint8{0, 0, 1, 1},
			[]uint8{0, 0, 0, 1},
		}, 2, 2, 5},
		{[][]uint8{
			[]uint8{1, 1, 1, 1},
			[]uint8{0, 1, 1, 1},
			[]uint8{0, 0, 1, 1},
			[]uint8{0, 0, 0, 1},
		}, 3, 0, 0},
		{[][]uint8{
			[]uint8{1, 1, 1, 1},
			[]uint8{0, 1, 1, 1},
			[]uint8{0, 0, 1, 1},
			[]uint8{0, 0, 0, 1},
		}, 3, 3, 2},
	}

	for _, c := range cases {
		if got := getCount(c.grid, c.in_x, c.in_y); got != c.want {
			t.Errorf("getCount failed: grid = %v, x = %v, y = %v, got = %v, want = %v", c.grid, c.in_x, c.in_y, got, c.want)
		}
	}
}
