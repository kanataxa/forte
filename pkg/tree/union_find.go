package tree

type UnionFind struct {
	par  []int
	size []int
}

func NewUnionFind(size int) *UnionFind {
	return new(UnionFind).init(size)
}

func (uf *UnionFind) init(size int) *UnionFind {
	uf.par = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.par[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Unite(p int, q int) {
	rp := uf.Root(p)
	rq := uf.Root(q)
	if rp == rq {
		return
	}
	if uf.size[rp] < uf.size[rq] {
		uf.par[rp] = uf.par[rq]
		uf.size[rq] += uf.size[rp]
	} else {
		uf.par[rq] = uf.par[rp]
		uf.size[rp] += uf.size[rq]
	}
}

func (uf *UnionFind) Root(p int) int {
	for uf.par[p] == p {
		return p
	}
	uf.par[p] = uf.Root(uf.par[p])
	return uf.par[p]
}

func (uf *UnionFind) Size(p int) int {
	return uf.size[uf.Root(p)]
}

func (uf *UnionFind) Same(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
}
