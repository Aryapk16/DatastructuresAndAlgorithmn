package main

import (
	"fmt"
)

type graph struct {
	vertices []*vertex
}
type vertex struct {
	data     int
	adjacent []*vertex
}

func main() {
	g := graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(5, 3)
	g.addEdge(2, 4)
	g.print()
	fmt.Println("\n BFS :")
	g.BFS(4)
}
func (g *graph) addVertex(data int) {
	if !contains(g.vertices, data) {
		g.vertices = append(g.vertices, &vertex{data: data})
	}
}
func contains(vertex []*vertex, data int) bool {
	for _, v := range vertex {
		if v.data == data {
			return true
		}
	}
	return false
}
func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Print("\n vertex", v.data, ":")
		for _, v := range v.adjacent {
			fmt.Print(" ", v.data)
		}

	}
}
func (g *graph) getVertex(data int) *vertex {
	for _, v := range g.vertices {
		if v.data == data {
			return v
		}
	}
	return nil
}
func (g *graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Error: Vertex not found.")
		return
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}

type Queue struct {
	arr []int
}

func (g *graph) BFS(vertex int) {
	queue := Queue{}
	visited := make(map[int]bool)
	queue.arr = append(queue.arr, vertex)
	visited[vertex] = true
	for len(queue.arr) != 0 {
		vertex = queue.arr[0]
		queue.arr = queue.arr[1:]
		fmt.Print(vertex, " ")
		for _, v := range g.getVertex(vertex).adjacent {
			if !visited[v.data] {
				visited[v.data] = true
				queue.arr = append(queue.arr, v.data)
			}
		}
	}
}
