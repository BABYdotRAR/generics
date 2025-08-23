package ds

import (
	"testing"

	"github.com/BABYdotRAR/generics"
)

func TestBFS(t *testing.T) {
	res := make(map[int][]int)
	expectedOutput, output := []int{12, 45, 33, 3, 4, 0, 20, 1, 7}, []int{}
	buildTestTree().BFS(0, res)
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
	buildTestTree().InOrder(&output)

	if !compareSlice(output, expectedOutput) {
		t.Errorf("Incorrect order, want: %+v, received: %+v", expectedOutput, output)
	}
}

func TestPreOrder(t *testing.T) {
	expectedOutput := []int{12, 45, 3, 0, 20, 33, 4, 1, 7}
	var output []int
	buildTestTree().PreOrder(&output)

	if !compareSlice(output, expectedOutput) {
		t.Errorf("Incorrect order, want: %+v, received: %+v", expectedOutput, output)
	}
}

func TestPostOrder(t *testing.T) {
	expectedOutput := []int{0, 20, 3, 45, 1, 7, 4, 33, 12}
	var output []int
	buildTestTree().PostOrder(&output)

	if !compareSlice(output, expectedOutput) {
		t.Errorf("Incorrect order, want: %+v, received: %+v", expectedOutput, output)
	}
}

func TestHeight(t *testing.T) {
	if buildTestTree().Height() != 3 {
		t.Errorf("Incorrect height, want: 3, received: %+v", buildTestTree().Height())
	}
}

func compareSlice[T comparable](s, target []T) bool {
	return generics.CompareSlice(s, target)
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
