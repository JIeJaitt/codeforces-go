package copypasta

// See 915E for example
// https://codeforces.com/contest/915/problem/E
// todo https://www.luogu.com.cn/problem/P5350
//      https://www.luogu.com.cn/problem/P5586

type odtNode struct {
	tpNode
	l, r int
}

func (t *treap) put1(l, r int, val tpValueType) {}
func (t *treap) floor(int) (_ *odtNode)       { return }
func (t *treap) next(int) (_ *odtNode)        { return }

func (t *treap) split(mid int) {
	if o := t.floor(mid); o.l < mid && mid <= o.r {
		r, val := o.r, o.val
		o.r = mid - 1
		t.put1(mid, r, val)
	}
}

func (t *treap) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func (t *treap) merge(l, r int, value tpValueType) {
	t.prepare(l, r)
	for o := t.next(l); o != nil && o.l <= r; o = t.next(o.l) {
		t.delete(tpKeyType(o.l))
	}
	o := t.floor(l)
	o.r = r
	o.val = value
}
