package QuickUnion

import "testing"

func TestUnionFind(t *testing.T) {
	type test struct {
		size      int
		unions    [][2]int
		connected [][]int
	}

	var tests = []test{
		{
			8,
			[][2]int{
				{1, 2},
				{5, 4},
				{4, 6},
			},
			[][]int{
				{0, 0},
				{1, 2},
				{3, 3},
				{5, 4, 6},
			},
		},
	}

	for _, tt := range tests {
		unionFind := NewUnionFind(tt.size)
		for _, e := range tt.unions {
			err := unionFind.UnionElements(e[0], e[1])
			if err != nil {
				t.Errorf("check UnionFind failed: excute union elements has err: %s. q %d p %d", err.Error(), e[0], e[1])
			}
		}

		for _, e := range tt.connected {
			for i := 0; i < len(e)-1; i++ {
				for j := i; j < len(e)-1; j++ {
					connected, err := unionFind.IsConnected(e[i], e[j])
					if err != nil {
						t.Errorf("check UnionFind failed: excute is connected has err: %s. q %d p %d", err.Error(), e[i], e[j])
					}

					if !connected {
						t.Errorf("check UnionFind failed: expect true actually false. q %d p %d", e[i], e[j])
					}
				}
			}

			for i := len(e) - 1; i >= 0; i-- {
				for j := i; j >= 0; j-- {
					connected, err := unionFind.IsConnected(e[i], e[j])
					if err != nil {
						t.Errorf("check UnionFind failed: excute is connected has err: %s. q %d p %d", err.Error(), e[i], e[j])
					}

					if !connected {
						t.Errorf("check UnionFind failed: expect true actually false. q %d p %d", e[i], e[j])
					}
				}
			}
		}
	}
}
