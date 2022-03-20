package backtracks

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// {
		// 	"case1",
		// 	args{
		// 		board: [][]byte{
		// 			{'A', 'B', 'C', 'E'},
		// 			{'S', 'F', 'C', 'S'},
		// 			{'A', 'D', 'E', 'E'},
		// 		},
		// 		word: "ABCCED",
		// 	},
		// 	true,
		// },
		{
			"case2",
			args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
