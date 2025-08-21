package ds

type TreeNode[T any] struct {
	Value       T
	Left, Right *TreeNode[T]
}

type Comparator[T any] interface {
	Equal(a, b T) bool
	Less(a, b T) bool
}

func BFS[T any](node *TreeNode[T], level int, res map[int][]T) {
	if node == nil {
		return
	}
	if res == nil {
		res = make(map[int][]T)
	}

	res[level] = append(res[level], node.Value)

	BFS(node.Left, level+1, res)
	BFS(node.Right, level+1, res)
}

func InOrder[T any](tree *TreeNode[T], res *[]T) {
	if tree == nil {
		return
	}
	InOrder(tree.Left, res)
	*res = append(*res, tree.Value)
	InOrder(tree.Right, res)
}

func PreOrder[T any](tree *TreeNode[T], res *[]T) {
	if tree == nil {
		return
	}

	*res = append(*res, tree.Value)
	PreOrder(tree.Left, res)
	PreOrder(tree.Right, res)
}

func PostOrder[T any](tree *TreeNode[T], res *[]T) {
	if tree == nil {
		return
	}

	PostOrder(tree.Left, res)
	PostOrder(tree.Right, res)
	*res = append(*res, tree.Value)
}