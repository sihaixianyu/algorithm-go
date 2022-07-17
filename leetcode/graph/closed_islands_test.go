package graph

import "testing"

func Test_closedIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				grid: [][]int{
					{0, 0, 1, 0, 0},
					{0, 1, 0, 1, 0},
					{0, 1, 1, 1, 0},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closedIsland(tt.args.grid); got != tt.want {
				t.Errorf("closedIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
