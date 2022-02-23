package slides

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				s: "abcabcbb",
			},
			3,
		},
		{
			"case2",
			args{
				s: "bbbbb",
			},
			1,
		},
		{
			"case3",
			args{
				s: "pwwkew",
			},
			3,
		},
		{
			"case4",
			args{
				s: "dvdf",
			},
			3,
		},
		{
			"case5",
			args{
				s: "abba",
			},
			2,
		},
	}
	for _, tt := range tests[len(tests)-1:] {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}