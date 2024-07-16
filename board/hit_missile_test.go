package board

import (
	"testing"

	"github.com/alvinwilta/go-simple-battleship-sim/types"
)

func Test_board_HitMissile(t *testing.T) {
	type fields struct {
		Ships  map[int]map[int]string
		Scores int
	}
	type args struct {
		input types.ListPoint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{
				Ships:  tt.fields.Ships,
				Scores: tt.fields.Scores,
			}
			b.HitMissile(tt.args.input)
		})
	}
}
