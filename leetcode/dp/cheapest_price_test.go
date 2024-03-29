package dp

import (
	"testing"
)

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// : Add test cases.
		{
			name: "case1",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				src:     0,
				dst:     2,
				k:       1,
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k); got != tt.want {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCheapestPrice2(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// : Add test cases.
		{
			name: "case1",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				src:     0,
				dst:     2,
				k:       1,
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCheapestPrice2(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k); got != tt.want {
				t.Errorf("findCheapestPrice2() = %v, want %v", got, tt.want)
			}
		})
	}
}
