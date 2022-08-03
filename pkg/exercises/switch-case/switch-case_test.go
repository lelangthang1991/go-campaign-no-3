package main

import "testing"

func Test_test(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{1},
			want: "one",
		},
		{
			name: "2",
			args: args{2},
			want: "two",
		},
		{
			name: "3",
			args: args{3},
			want: "three",
		},
		{
			name: "4",
			args: args{4},
			want: "nothing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(tt.args.i); got != tt.want {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}
