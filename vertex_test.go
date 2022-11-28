package main

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	type args struct {
		width         int
		height        int
		start         Point
		finish        Point
		obstacleCount int
		obstacles     [][2][2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "need to go backwards to gain speed",
			args: args{
				width:         4,
				height:        1,
				start:         Point{1, 0},
				finish:        Point{3, 0},
				obstacleCount: 1,
				obstacles: [][2][2]int{
					{
						{2, 2},
						{0, 0},
					},
				},
			},
			want: 4,
		},
		{
			name: "impossible to jump",
			args: args{
				width:         5,
				height:        1,
				start:         Point{1, 0},
				finish:        Point{4, 0},
				obstacleCount: 1,
				obstacles: [][2][2]int{
					{
						{2, 3},
						{0, 0},
					},
				},
			},
			want: 0,
		},
		{
			name: "C-like route",
			args: args{
				width:         4,
				height:        5,
				start:         Point{3, 0},
				finish:        Point{3, 4},
				obstacleCount: 5,
				obstacles: [][2][2]int{
					{
						{0, 0},
						{0, 0},
					},
					{
						{1, 2},
						{1, 1},
					},
					{
						{1, 3},
						{2, 3},
					},
					{
						{0, 0},
						{4, 4},
					},
					{
						{2, 2},
						{4, 4},
					},
				},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			vertex := NewVertex(0, tt.args.start, Velocity{DX: 0, DY: 0}, nil)
			obstacles := CreateObstacles(tt.args.obstacleCount, tt.args.obstacles)
			result := vertex.BFS(tt.args.finish, tt.args.width, tt.args.height, obstacles)
			fmt.Println(result, result == nil)

			if result == nil && tt.want != 0 {
				t.Errorf("get 0, want %v", tt.want)
			} else if result != nil && result.value != tt.want {
				t.Errorf("get %v, want %v", result.value, tt.want)
			}
		})
	}
}
