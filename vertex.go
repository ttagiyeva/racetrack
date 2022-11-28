package main

import (
	"math"
)

// Point represents position of graph vertex
type Point struct {
	x int
	y int
}

// Velocity represents delta X/Y per move
// velocity with dx=1,dy=1 would move from top-left to bottom-right
type Velocity struct {
	DX int
	DY int
}

// Vertex represents a graph vertex
type Vertex struct {
	value int
	coor  Point
	speed Velocity
	edges []*Vertex
	past  *Vertex //for animated representetion
}

//NewVertex creates a new vertex
func NewVertex(value int, coor Point, speed Velocity, past *Vertex) *Vertex {
	return &Vertex{
		value: value,
		coor:  coor,
		speed: speed,
		past:  past,
	}
}

// keep track of visited (considering velocity) places
var visitedEdges = make(map[Visit]bool)

//createGraph creates the graph for the current vertex
func (v *Vertex) createGraph(width, height int, obstacles []Obstacle) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			dx := v.speed.DX + i
			dy := v.speed.DY + j
			x := v.coor.x + dx
			y := v.coor.y + dy

			if x >= 0 && x < width && y >= 0 && y < height && math.Abs(float64(dx)) <= 3 && math.Abs(float64(dy)) <= 3 {
				edge := NewVertex(v.value+1, Point{x, y}, Velocity{dx, dy}, v)
				v.addEdge(edge, obstacles)
			}

		}
	}
}

//isObstacle checks if the current vertex is an obstacle
func (v *Vertex) isObstacle(obstacles []Obstacle) bool {
	for _, obstacle := range obstacles {
		if v.coor.x == obstacle.point.x && v.coor.y == obstacle.point.y {
			return true
		}
	}
	return false
}

//addEdge adds an edge to the current vertex
func (v *Vertex) addEdge(vertex *Vertex, obstacles []Obstacle) {
	if !vertex.isObstacle(obstacles) {

		visit := Visit{Point: vertex.coor, Velocity: vertex.speed}
		if visitedEdges[visit] {
			return
		}

		visitedEdges[visit] = true
		v.edges = append(v.edges, vertex)
	}
}

//finished checks if the current vertex is the finish
func (v *Vertex) finished(finish Point) bool {
	return v.coor.x == finish.x && v.coor.y == finish.y
}

//BFS performs a breadth first search on the graph
func (v *Vertex) BFS(finish Point, width, height int, obstacles []Obstacle) *Vertex {
	var queue []*Vertex
	queue = append(queue, v)

	for len(queue) > 0 {

		current := queue[0]
		if current.finished(finish) {
			return current
		}

		queue = queue[1:]
		current.createGraph(width, height, obstacles)
		queue = append(queue, current.edges...)
	}

	return nil
}

// History returns list of vertexes from the initial vertex to this version of vertex
// This is used for animated representation
func (v *Vertex) History() []*Vertex {
	vertexes := make([]*Vertex, v.value+1)
	i := len(vertexes) - 1
	current := v
	for current != nil {
		vertexes[i] = current
		current = current.past
		i--
	}
	return vertexes
}
