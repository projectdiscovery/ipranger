package ipranger

import (
	"reflect"
	"testing"
)

func TestNewBlackRock(t *testing.T) {
	type args struct {
		rangez int64
		seed   int64
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 *BlackRock
	}{
		{
			name: "must solve the square root and increment B while the result is less than rangez",
			args: func(*testing.T) args {
				return args{
					rangez: 4,
					seed:   1,
				}
			},
			want1: &BlackRock{
				Rounds: 3,
				Seed:   1,
				Range:  4,
				A:      1,
				B:      5,
			},
		},
		{
			name: "if split is zero the value of A must be 1",
			args: func(*testing.T) args {
				return args{
					rangez: 0,
					seed:   1,
				}
			},
			want1: &BlackRock{
				Rounds: 3,
				Seed:   1,
				Range:  0,
				A:      1,
				B:      1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := NewBlackRock(tArgs.rangez, tArgs.seed)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewBlackRock got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
