package main

//Obstacle is a struct that represents an obstacle
type Obstacle struct {
	point Point
}

//NewObstacle creates a new obstacle
func NewObstacle(p Point) *Obstacle {
	return &Obstacle{p}
}

//CreateOBstacles creates the obstacles for the graph
func CreateObstacles(count int, obstaclesCoor [][2][2]int) []Obstacle {
	obstacles := make([]Obstacle, 0)
	for k := 0; k < count; k++ {
		for i := obstaclesCoor[k][0][0]; i <= obstaclesCoor[k][0][1]; i++ {
			for j := obstaclesCoor[k][1][0]; j <= obstaclesCoor[k][1][1]; j++ {
				obstacles = append(obstacles, *NewObstacle(Point{i, j}))
			}
		}
	}
	return obstacles
}

// Visit represents a state that can have unique consequences
// In our task, marking only points in grid as visited is not helpful
// Because, from the same point, different hoppers can navigate to different places
// based on their velocity (e.g. one can jump over an obstacle, the other can't)
type Visit struct {
	Point
	Velocity
}
