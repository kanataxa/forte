package tree

type SV int

const def = 0

type Segment struct {
	n    int
	data []SV
	fn   func(SV, SV) SV
	def  SV
}

func NewSegment(n int, fn func(SV, SV) SV) *Segment {
	return new(Segment).init(n, fn)
}

func (s *Segment) init(n int, fn func(SV, SV) SV) *Segment {
	s.n = 1
	for s.n < n {
		s.n *= 2
	}
	s.data = make([]SV, 2*s.n-1)
	s.fn = fn
	s.def = def
	for i := 0; i < 2*s.n-1; i++ {
		s.data[i] = def
	}
	return s
}

func (s *Segment) Update(idx int, value SV) {
	idx = idx + s.n - 1
	s.data[idx] = value
	for idx > 0 {
		idx = (idx - 1) / 2
		s.data[idx] = s.fn(s.data[idx*2+1], s.data[idx*2+2])
	}
}

func (s *Segment) Query(p, q int) SV {
	return s.query(p, q, 0, 0, s.n)
}

func (s *Segment) query(p, q, k, l, r int) SV {
	if r <= p || q <= l {
		return s.def
	}

	if p <= l && r <= q {
		return s.data[k]
	}

	m := (l + r) / 2
	lv := s.query(p, q, 2*k+1, l, m)
	rv := s.query(p, q, 2*k+2, m, r)
	return s.fn(lv, rv)
}

func (s *Segment) Get(idx int) SV {
	return s.data[idx+(s.n-1)]
}
