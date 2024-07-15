package board

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want Board
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(tt.args.n, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
