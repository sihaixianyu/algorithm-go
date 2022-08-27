package dp

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
            name: "case1",
            args: args{
               s: "aa",
			   p: "a", 
            },
            want: false,
        },
		{
            name: "case2",
            args: args{
               s: "aa",
			   p: "a*", 
            },
            want: true,
        },
		{
            name: "case3",
            args: args{
               s: "ab",
			   p: ".*", 
            },
            want: true,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
