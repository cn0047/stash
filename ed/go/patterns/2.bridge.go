package main

type BinaryTree struct {
	tree interface{}
}

func (bt BinaryTree) IsBalanced() bool {
	return false
}

func (bt BinaryTree) Balance() {
	if bt.IsBalanced() {
		return
	}

	//bt.PerformBalance()
}

func NewBinaryTreeBridge() BinaryTreeBridge {
	return BinaryTreeBridge{tree: BinaryTree{}}
}

type BinaryTreeBridge struct {
	tree BinaryTree
}

func (btb BinaryTreeBridge) Maintaine() {
	btb.tree.Balance()
}

func main() {
	NewBinaryTreeBridge().Maintaine()
}
