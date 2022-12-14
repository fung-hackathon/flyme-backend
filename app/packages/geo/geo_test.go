package geo_test

import (
	"flyme-backend/app/packages/geo"
	"math"
	"testing"
)

func TestGetDistance(t *testing.T) {

	ignorableKm := 1e-5

	type args struct {
		coords []geo.Coordinate
	}

	tests := []struct {
		name   string
		arg    args
		want   float64
		hasErr bool
	}{
		{
			name: "はこだて未来大と函館市役所の距離(Km)",
			arg: args{
				coords: []geo.Coordinate{
					{
						Longitude: 140.766944,
						Latitude:  41.841806,
					},
					{
						Longitude: 140.72892,
						Latitude:  41.76867,
					},
				},
			},
			want:   8.716124,
			hasErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distance, err := geo.GetDistanceKm(tt.arg.coords)

			if (err != nil) != tt.hasErr {
				t.Errorf("GetDistance() error = %v, hasErr %v", err, tt.hasErr)
			}
			if math.Abs(float64(distance-tt.want)) > float64(ignorableKm) {
				t.Errorf("GetDistance() = %f, want %f", distance, float64(tt.want))
			}
		})
	}
}
