package VariadicFunctions

import (
	"testing"
)

func Test_max(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max-1",
			args: args{
				values: []int{1, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "max-2",
			args: args{
				values: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.values...); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "min-1",
			args: args{
				values: []int{1, 2, 3, 4},
			},
			want: 1,
		}, {
			name: "min-2",
			args: args{
				values: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.values...); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
