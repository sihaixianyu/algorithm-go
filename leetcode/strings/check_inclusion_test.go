package str

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// : Add test cases.
		{
			"case1",
			args{
				s1: "ab",
				s2: "eidbaooo",
			},
			true,
		},
		{
			"case2",
			args{
				s1: "abcdxabcde",
				s2: "abcdeabcdx",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
