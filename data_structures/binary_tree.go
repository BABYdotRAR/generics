package ds

import "github.com/BABYdotRAR/generics"

// NewBinaryTree retrieves a new generic binary tree
func NewBinaryTree[T any]() *TreeNode[T] {
	return &TreeNode[T]{}
}

// TreeNode is the base struct for a generic binary tree
type TreeNode[T any] struct {
	Value       T
	Left, Right *TreeNode[T]
}

// Height returns the maximum depth of the tree
func (t *TreeNode[T]) Height() int {
	if t == nil {
		return -1
	}
	lHeight := t.Left.Height()
	rHeight := t.Right.Height()
	return generics.Max(lHeight, rHeight) + 1
}

type Comparator[T any] interface {
	Equal(a, b T) bool
	Less(a, b T) bool
}

// BFS returns res which is a map where the key represents the level of the tree
// and the values are the elements in that level from left to right
func (node *TreeNode[T]) BFS(level int, res map[int][]T) {
	if node == nil {
		return
	}
	if res == nil {
		res = make(map[int][]T)
	}

	res[level] = append(res[level], node.Value)

	node.Left.BFS(level+1, res)
	node.Right.BFS(level+1, res)
}

// InOrder visits all nodes in the sequence: Left → Root → Right
func (tree *TreeNode[T]) InOrder(res *[]T) {
	if tree == nil {
		return
	}
	tree.Left.InOrder(res)
	*res = append(*res, tree.Value)
	tree.Right.InOrder(res)
}

// PreOrder visits all nodes in the sequence: Root → Left → Right
func (tree *TreeNode[T]) PreOrder(res *[]T) {
	if tree == nil {
		return
	}

	*res = append(*res, tree.Value)
	tree.Left.PreOrder(res)
	tree.Right.PreOrder(res)
}

// PostOrder visits all nodes in the sequence: Left → Right → Root
func (tree *TreeNode[T]) PostOrder(res *[]T) {
	if tree == nil {
		return
	}

	tree.Left.PostOrder(res)
	tree.Right.PostOrder(res)
	*res = append(*res, tree.Value)
}
