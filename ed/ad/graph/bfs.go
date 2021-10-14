package main

import (
	"fmt"
)

func main() {
	var r int32

	r = bfs([][]int32{{1, 2}, {1, 3}, {3, 4}, {4, 5}}, 4, 1)
	fmt.Printf("case 1: %+v \n", r == 2)

	r = bfs([][]int32{{1, 2}, {1, 3}, {3, 4}, {4, 5}}, 1, 5)
	fmt.Printf("case 2: %+v \n", r == 3)

	r = bfs([][]int32{{1, 2}, {1, 3}, {3, 4}, {1, 5}, {5, 6}, {6, 7}}, 7, 2)
	fmt.Printf("case 3: %+v \n", r == 4)

	r = bfs([][]int32{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}, {3, 7}, {6, 8}, {8, 9}, {9, 10}}, 5, 10)
	fmt.Printf("case 4: %+v \n", r == 7)
}

func bfs(edges [][]int32, from int32, to int32) int32 {
	g := NewGraph()
	g.Init(edges)

	return g.BFS(from, to)
}

type Node struct {
	Value int32
	Edges map[int32]struct{}
}

type Graph map[int32]*Node

func NewGraph() Graph {
	return make(Graph)
}

func (g Graph) GetNode(value int32) *Node {
	node, ok := g[value]
	if !ok {
		return nil
	}

	return node
}

func (g Graph) SetNode(node *Node) {
	g[node.Value] = node
}

func (g Graph) AddNode(value int32) {
	node := &Node{Value: value, Edges: make(map[int32]struct{}, 0)}
	g.SetNode(node)
}

func (g Graph) AddEdge(fromValue int32, toValue int32) {
	fromNode := g.GetNode(fromValue)
	if fromNode == nil {
		fromNode = &Node{Value: fromValue, Edges: make(map[int32]struct{}, 0)}
		g.SetNode(fromNode)
	}

	toNode := g.GetNode(toValue)
	if toNode == nil {
		toNode = &Node{Value: toValue, Edges: make(map[int32]struct{}, 0)}
		g.SetNode(toNode)
	}

	fromNode.Edges[toValue] = struct{}{}
	toNode.Edges[fromValue] = struct{}{}
}

func (g Graph) Init(edges [][]int32) {
	for _, edge := range edges {
		if len(edge) == 0 {
			continue
		}

		fromValue := edge[0]

		// Not edge but only 1 node.
		if len(edge) == 1 {
			g.AddNode(fromValue)
			continue
		}

		toValue := edge[1]

		// Edge between 2 nodes.
		g.AddEdge(fromValue, toValue)
	}
}

func (g Graph) BFS(fromValue int32, toValue int32) int32 {
	if fromValue == toValue {
		return 0
	}

	fromNode := g.GetNode(fromValue)
	if fromNode == nil {
		return -1
	}
	toNode := g.GetNode(toValue)
	if toNode == nil {
		return -1
	}

	queue := NewQueue()
	seenEdges := make(map[int32]struct{})

	seenEdges[fromValue] = struct{}{}
	for edge, _ := range fromNode.Edges {
		queue.Push(edge, 1)
	}

	for queue.Len() > 0 {
		queueElement := queue.Pop()

		if queueElement.Value == toValue {
			return queueElement.Path
		}

		node := g.GetNode(queueElement.Value)
		if node == nil {
			continue
		}

		seenEdges[node.Value] = struct{}{}
		for edge, _ := range node.Edges {
			_, seen := seenEdges[edge]
			if seen {
				continue
			}
			queue.Push(edge, queueElement.Path+1)
		}
	}

	return -1
}

type QueueElement struct {
	Value int32
	Path  int32
}

type Queue struct {
	Data []*QueueElement
}

func NewQueue() *Queue {
	return &Queue{Data: make([]*QueueElement, 0)}
}

func (q *Queue) Len() int {
	return len(q.Data)
}

func (q *Queue) Push(value int32, path int32) {
	q.Data = append(q.Data, &QueueElement{Value: value, Path: path})
}

func (q *Queue) Pop() *QueueElement {
	if len(q.Data) == 0 {
		return nil
	}

	value := q.Data[0]
	q.Data = q.Data[1:]

	return value
}
