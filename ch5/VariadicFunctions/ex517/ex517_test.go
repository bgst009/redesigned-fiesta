package ex517

import (
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	type args struct {
		doc  *html.Node
		name []string
	}
	tests := []struct {
		name string
		args args
		want []*html.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementsByTagName(tt.args.doc, tt.args.name...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementsByTagName() = %v, want %v", got, tt.want)
			}
		})
	}
}
