package isMatch

import "testing"

func Test_isMatchDalao(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatchDalao(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatchDalao() = %v, want %v", got, tt.want)
			}
		})
	}
}
