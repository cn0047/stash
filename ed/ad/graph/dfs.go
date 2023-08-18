package main

import (
	"fmt"
)

func main() {
	var r int32

	r = dfs([][]int32{{1, 2}, {1, 3}, {3, 4}, {4, 5}}, 4, 1)
	fmt.Printf("case 1: %+v \n", r == 2)

	r = dfs([][]int32{{1, 2}, {1, 3}, {3, 4}, {4, 5}}, 1, 5)
	fmt.Printf("case 1: %+v \n", r == 3)

	r = dfs([][]int32{{1, 2}, {1, 3}, {3, 4}, {1, 5}, {5, 6}, {6, 7}}, 7, 2)
	fmt.Printf("case 3: %+v \n", r == 4)
}

func dfs(edges [][]int32, from int32, to int32) int32 {
	g := NewGraph()
	g.Init(edges)

	return DFS(g, from, to)
}

// DFS - Depth First Search.
//
// step 1: Put root node into stack.
// step 2: Get element from stack, process value and put into stack node's children.
// step 2: Goto step 2.
func DFS(g Graph, fromValue int32, toValue int32) int32 {
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

	stack := NewStack()
	seenEdges := make(map[int32]struct{})

	seenEdges[fromValue] = struct{}{}
	for edge, _ := range fromNode.Edges {
		stack.Push(edge, 1)
	}

	for stack.Len() > 0 {
		itemFromStack := stack.Pop()

		if itemFromStack.Value == toValue {
			return itemFromStack.Path
		}

		node := g.GetNode(itemFromStack.Value)
		if node == nil {
			continue
		}

		seenEdges[node.Value] = struct{}{}
		for edge, _ := range node.Edges {
			_, seen := seenEdges[edge]
			if seen {
				continue
			}
			stack.Push(edge, itemFromStack.Path+1)
		}
	}

	return -1
}

// ============================================================================

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

// ============================================================================

type itemFromStack struct {
	Value int32
	Path  int32
}

// Stack represents struct to work with stack data structure.
// Key info: stack tail is in data slice beginning, head in the ending of slice,
// so index in slice it's like sequential number in slice.
type Stack struct {
	Data []*itemFromStack
}

func NewStack() *Stack {
	return &Stack{Data: make([]*itemFromStack, 0)}
}

func (s *Stack) Len() int {
	return len(s.Data)
}

func (s *Stack) Push(value int32, path int32) {
	item := &itemFromStack{Value: value, Path: path}
	s.Data = append([]*itemFromStack{item}, s.Data...)
}

func (s *Stack) Pop() *itemFromStack {
	item := s.Data[0]
	s.Data = s.Data[1:]

	return item
}
