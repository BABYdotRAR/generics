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
