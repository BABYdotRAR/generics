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
