package ds

import "testing"

func TestBFS(t *testing.T) {
	res := make(map[int][]int)
	expectedOutput, output := []int{12, 45, 33, 3, 4, 0, 20, 1, 7}, []int{}
	BFS(buildTestTree(), 0, res)
	for i := 0; i < len(res); i++ {
		output = append(output, res[i]...)
	}

	if !compareSlice(output, expectedOutput) {
		t.Errorf("Incorrect order, want: %+v, received: %+v", expectedOutput, output)
	}
}

func TestInOrder(t *testing.T) {
	expectedOutput := []int{0, 3, 20, 45, 12, 33, 1, 4, 7}
	var output []int
	InOrder(buildTestTree(), &output)

	if !compareSlice(output, expectedOutput) {
		t.Errorf("Incorrect order, want: %+v, received: %+v", expectedOutput, output)
	}
}

func TestPreOrder(t *testing.T) {
	expectedOutput := []int{12, 45, 3, 0, 20, 33, 4, 1, 7}
	var output []int
	PreOrder(buildTestTree(), &output)

	if !compareSlice(output, expectedOutput) {
		t.Errorf("Incorrect order, want: %+v, received: %+v", expectedOutput, output)
	}
}

func TestPostOrder(t *testing.T) {
	expectedOutput := []int{0, 20, 3, 45, 1, 7, 4, 33, 12}
	var output []int
	PostOrder(buildTestTree(), &output)

	if !compareSlice(output, expectedOutput) {
		t.Errorf("Incorrect order, want: %+v, received: %+v", expectedOutput, output)
	}
}

func compareSlice[T comparable](s, target []T) bool {
	if len(s) != len(target) {
		return false
	}
	for i := range s {
		if s[i] != target[i] {
			return false
		}
	}
	return true
}

func buildTestTree() *TreeNode[int] {
	/*
		             12
		         45      33
			   3            4
			 0  20        1   7
	*/
	return &TreeNode[int]{
		Value: 12,
		Left: &TreeNode[int]{
			Value: 45,
			Left: &TreeNode[int]{
				Value: 3,
				Left:  &TreeNode[int]{},
				Right: &TreeNode[int]{
					Value: 20,
				},
			},
		},
		Right: &TreeNode[int]{
			Value: 33,
			Right: &TreeNode[int]{
				Value: 4,
				Left: &TreeNode[int]{
					Value: 1,
				},
				Right: &TreeNode[int]{
					Value: 7,
				},
			},
		},
	}
}
