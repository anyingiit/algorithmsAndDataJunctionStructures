package UnionFind

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/UnionFind/QuickFind"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/UnionFind/QuickUnionV1"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/UnionFind/QuickUnionV2"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/UnionFind/QuickUnionV3"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/UnionFind/QuickUnionV4"
	"math/rand"
	"testing"
)

func BenchmarkQuickFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unionFind := QuickFind.NewUnionFind(b.N)
		for j := 0; j < b.N; j++ {
			err := unionFind.UnionElements(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute UnionElements has err %s", err.Error())
			}
		}
		for j := 0; j < b.N; j++ {
			_, err := unionFind.IsConnected(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute IsConnected has err %s", err.Error())
			}
		}
	}
}

func BenchmarkQuickUnionV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unionFind := QuickUnionV1.NewUnionFind(b.N)
		for j := 0; j < b.N; j++ {
			err := unionFind.UnionElements(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute UnionElements has err %s", err.Error())
			}
		}
		for j := 0; j < b.N; j++ {
			_, err := unionFind.IsConnected(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute IsConnected has err %s", err.Error())
			}
		}
	}
}

func BenchmarkQuickUnionV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unionFind := QuickUnionV2.NewUnionFind(b.N)
		for j := 0; j < b.N; j++ {
			err := unionFind.UnionElements(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute UnionElements has err %s", err.Error())
			}
		}
		for j := 0; j < b.N; j++ {
			_, err := unionFind.IsConnected(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute IsConnected has err %s", err.Error())
			}
		}
	}
}

func BenchmarkQuickUnionV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unionFind := QuickUnionV3.NewUnionFind(b.N)
		for j := 0; j < b.N; j++ {
			err := unionFind.UnionElements(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute UnionElements has err %s", err.Error())
			}
		}
		for j := 0; j < b.N; j++ {
			_, err := unionFind.IsConnected(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute IsConnected has err %s", err.Error())
			}
		}
	}
}

func BenchmarkQuickUnionV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unionFind := QuickUnionV4.NewUnionFind(b.N)
		for j := 0; j < b.N; j++ {
			err := unionFind.UnionElements(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute UnionElements has err %s", err.Error())
			}
		}
		for j := 0; j < b.N; j++ {
			_, err := unionFind.IsConnected(rand.Intn(b.N), rand.Intn(b.N))
			if err != nil {
				b.Errorf("failed: excute IsConnected has err %s", err.Error())
			}
		}
	}
}
