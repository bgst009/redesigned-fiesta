package ex516

import "testing"

func TestStringJoin(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name     string
		args     args
		wantJoin string
	}{
		{
			name: "StringJoin",
			args: args{
				ss: []string{"The", "Go", "Programming"},
			},
			wantJoin: "TheGoProgramming",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJoin := StringJoin(tt.args.ss...); gotJoin != tt.wantJoin {
				t.Errorf("StringJoin() = %v, want %v", gotJoin, tt.wantJoin)
			}
		})
	}
}
