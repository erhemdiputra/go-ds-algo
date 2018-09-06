package main

import (
	"container/list"
	"fmt"
)

const maxVertex = 100

var (
	graph          [maxVertex][maxVertex]bool
	visited        [maxVertex]bool
	nVertex, nEdge int
)

type Data struct {
	vertex int
	step   int
}

func NewData(vertex int, step int) Data {
	return Data{
		vertex: vertex,
		step:   step,
	}
}

func bfs(start int, finish int) int {
	data := NewData(start, 1)

	queue := list.New()
	queue.PushBack(data)

	visited[start] = true

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		current, _ := element.Value.(Data)

		if current.vertex == finish {
			return current.step
		}

		for i := 1; i <= nVertex; i++ {
			if graph[current.vertex][i] && !visited[i] {
				next := NewData(i, current.step+1)
				queue.PushBack(next)
				visited[i] = true
			}
		}
	}

	return -1 // no path
}

func main() {
	fmt.Scanf("%d %d\n", &nVertex, &nEdge)

	for i := 1; i <= nEdge; i++ {
		var v1, v2 int
		fmt.Scanf("%d %d\n", &v1, &v2)
		graph[v1][v2] = true
		graph[v2][v1] = true
	}

	var start, finish int
	fmt.Scanf("%d %d\n", &start, &finish)

	steps := bfs(start, finish)
	if steps != -1 {
		fmt.Printf("There is a path from %d to %d with %d steps.\n", start, finish, steps)
	} else {
		fmt.Printf("There is no path from %d to %d\n", start, finish)
	}
}

// Sample Input :
// 10 7
// 1 3
// 2 5
// 6 7
// 1 8
// 9 10
// 5 9
// 1 2
// 1 10

// Sample Output :
// There is a path from 1 to 10 with 5 steps
