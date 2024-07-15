package types

import (
	"fmt"
	"testing"
)

func TestListPoint_PopulateListPoint(t *testing.T) {
	newList := ListPoint{
		Points: []Point{},
	}
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				input: "1,1:2,0:2,3:3,4:4,3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newList
			l.PopulateListPoint(tt.args.input)
			fmt.Println(l)
		})
	}
}
