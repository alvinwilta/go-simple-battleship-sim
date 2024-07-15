package board

import (
	"testing"

	"github.com/alvinwilta/gojek-coding-test/types"
)

func Test_board_FillBoard(t *testing.T) {
	boardShip := NewBoard(5, 5)
	testResult := NewBoard(5, 5)
	type args struct {
		input types.ListPoint
	}
	tests := []struct {
		name string
		args args
		want Board
	}{
		{
			name: "success 1 point",
			args: args{
				input: types.ListPoint{
					Points: []types.Point{
						{
							X: 1,
							Y: 1,
						},
						{
							X: 5,
							Y: 6,
						},
					},
				},
			},
			want: testResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := boardShip
			b.FillBoard(tt.args.input)
		})
	}
}
