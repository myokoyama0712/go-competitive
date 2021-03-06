/*
URL:
https://atcoder.jp/contests/abc140/tasks/abc140_e
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

/********** FAU standard libraries **********/

//fmt.Sprintf("%b\n", 255) 	// binary expression

/********** I/O usage **********/

//str := ReadString()
//i := ReadInt()
//X := ReadIntSlice(n)
//S := ReadRuneSlice()
//a := ReadFloat64()
//A := ReadFloat64Slice(n)

//str := ZeroPaddingRuneSlice(num, 32)
//str := PrintIntsLine(X...)

/*******************************************************************/

const (
	// General purpose
	MOD = 1000000000 + 7
	// MOD          = 998244353
	ALPHABET_NUM = 26
	INF_INT64    = math.MaxInt64
	INF_BIT60    = 1 << 60
	INF_INT32    = math.MaxInt32
	INF_BIT30    = 1 << 30
	NIL          = -1

	// for dijkstra, prim, and so on
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

var (
	n int
	P []int
)

type Item struct {
	key      int
	idx, val int
}
type ItemList []*Item
type byKey struct {
	ItemList
}

func (l ItemList) Len() int {
	return len(l)
}
func (l ItemList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	return l.ItemList[i].key < l.ItemList[j].key
}

// how to use
// L := make(ItemList, 0, 200000+5)
// L = append(L, &Item{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

func main() {
	n = ReadInt()
	P = ReadIntSlice(n)

	L := make(ItemList, 0)
	for i := 0; i < n; i++ {
		L = append(L, &Item{key: P[i], idx: i, val: P[i]})
	}
	sort.Stable(sort.Reverse(byKey{L}))

	tr := NewTreap(func(l, r T) bool { return l < r })
	tr.Insert(T(L[0].idx))
	ans := 0
	for i := 1; i < n; i++ {
		idx, val := L[i].idx, L[i].val
		r1 := tr.MinGreater(T(idx))
		l1 := tr.MaxLess(T(idx))

		if r1 != nil {
			// 右側に1つだけ大きいものを含む区間
			r2 := tr.MinGreater(T(r1.key))
			var c, d int
			c = int(r1.key)
			if r2 != nil {
				d = int(r2.key) - 1
			} else {
				d = n - 1
			}

			var a, b int
			b = idx
			if l1 != nil {
				a = int(l1.key) + 1
			} else {
				a = 0
			}

			num := (d - c + 1) * (b - a + 1)
			ans += val * num
		}

		if l1 != nil {
			// 左側に1つだけ大きいものを含む区間
			l2 := tr.MaxLess(T(l1.key))
			var a, b int
			b = int(l1.key)
			if l2 != nil {
				a = int(l2.key) + 1
			} else {
				a = 0
			}

			var c, d int
			c = idx
			if r1 != nil {
				d = int(r1.key) - 1
			} else {
				d = n - 1
			}

			num := (d - c + 1) * (b - a + 1)
			ans += val * num
		}

		tr.Insert(T(idx))
	}

	fmt.Println(ans)
}

// Treap usage
// tr := NewTreap() 				// constructor
// tr.Insert(key) 					// insert one key node
// cnt := tr.Count(key) 		// return a number of key nodes
// node := tr.Find(key) 		// return a pointer
// min := tr.FindMinimum() 	// return a pointer
// max := tr.FindMaximum() 	// return a pointer
// tr.Delete(key) 					// delete one key node
// node := tr.MinGeq(x) 		// return a pointer
// node := tr.MinGreater(x) // return a pointer
// node := tr.MaxLeq(x) 		// return a pointer
// node := tr.MaxLess(x) 		// return a pointer

// fmt.Println(PrintIntsLine(tr.Inorder()...))
// fmt.Println(PrintIntsLine(tr.Preorder()...))
// tr.InsertBySettingPri(key, p)

// type of key
// type T struct {
// 	s, t int
// }

type T int

type Node struct {
	key         T
	priority    int
	right, left *Node
}

type Treap struct {
	root    *Node
	cnts    map[T]int
	randInt func() int
	less    func(l, r T) bool // *strictly less*
}

/*************************************/
// Public method
/*************************************/

// NewTreap returns a pointer of a Treap instance.
func NewTreap(less func(l, r T) bool) *Treap {
	tr := new(Treap)

	tr.root = nil
	tr.cnts = make(map[T]int)
	tr.less = less

	// XorShiftによる乱数生成
	// 下記URLを参考
	// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
	_gtx, _gty, _gtz, _gtw := 123456789, 362436069, 521288629, 88675123
	tr.randInt = func() int {
		tt := (_gtx ^ (_gtx << 11))
		_gtx = _gty
		_gty = _gtz
		_gtz = _gtw
		_gtw = (_gtw ^ (_gtw >> 19)) ^ (tt ^ (tt >> 8))
		return _gtw
	}

	return tr
}

// Count method returns the number of the key.
// If there hasn't been the key in the treap, this returns 0.
func (tr *Treap) Count(key T) int {
	return tr.cnts[key]
}

// InsertBySettingPri method inserts a new node consisting of new key and priority.
// A duplicate key is ignored and nothing happens.
// func (tr *Treap) InsertBySettingPri(key, priority int) {
// 	tr.root = tr.insert(tr.root, key, priority)
// }

// Insert method inserts a new node consisting o new key.
// The priority is automatically set by random value.
// A duplicate key is ignored and nothing happens.
func (tr *Treap) Insert(key T) {
	preCnt := tr.Count(key)
	tr.increase(key, 1)
	if preCnt > 0 {
		return
	}

	tr.root = tr.insert(tr.root, key, tr.randInt())
}

// Find returns a node that has an argument key value.
// Find returns nil when there is no node that has an argument key value.
func (tr *Treap) Find(key T) *Node {
	cnt := tr.cnts[key]
	if cnt == 0 {
		return nil
	}

	u := tr.root
	for u != nil && key != u.key {
		if tr.less(key, u.key) {
			u = u.left
		} else {
			u = u.right
		}
	}
	return u
}

// FindMinimum returns a node that has the minimum key in the treap.
// FindMinimum returns nil when there is no nodes.
func (tr *Treap) FindMinimum() *Node {
	u := tr.root
	for u != nil && u.left != nil {
		u = u.left
	}
	return u
}

// FindMaximum returns a node that has the maximum key in the treap.
// FindMaximum returns nil when there is no nodes.
func (tr *Treap) FindMaximum() *Node {
	u := tr.root
	for u != nil && u.right != nil {
		u = u.right
	}
	return u
}

// Delete method deletes a node that has an argument key value.
// A duplicate key is ignored and nothing happens.
func (tr *Treap) Delete(key T) {
	tr.decrease(key, 1)
	curCnt := tr.Count(key)
	if curCnt > 0 {
		return
	}

	tr.root = tr.delete(tr.root, key)
}

// Inorder returns a slice consisting of treap nodes in order of INORDER.
// The nodes are sorted by key values.
func (tr *Treap) Inorder() []T {
	res := make([]T, 0, 200000+5)
	tr.inorder(tr.root, &res)
	return res
}

// Preorder returns a slice consisting of treap nodes in order of PREORDER.
func (tr *Treap) Preorder() []T {
	res := make([]T, 0, 200000+5)
	tr.preorder(tr.root, &res)
	return res
}

// MinGeq returns a node that has MINIMUM KEY MEETING key >= x.
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func (tr *Treap) MinGeq(x T) *Node {
	return tr.biggerLowerBound(tr.root, x)
}

// MinGreater returns a node that has MINIMUM KEY MEETING key > x.
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func (tr *Treap) MinGreater(x T) *Node {
	return tr.biggerUpperBound(tr.root, x)
}

// MaxLeq returns a node that has MAXIMUM KEY MEETING key <= x.
// for AGC005-B
func (tr *Treap) MaxLeq(x T) *Node {
	return tr.smallerUpperBound(tr.root, x)
}

// MaxLess returns a node that has MAXIMUM KEY MEETING key < x.
// for AGC005-B
func (tr *Treap) MaxLess(x T) *Node {
	return tr.smallerLowerBound(tr.root, x)
}

/*************************************/
// Private method
/*************************************/

func (tr *Treap) increase(key T, num int) {
	tr.cnts[key] += num
}

func (tr *Treap) decrease(key T, num int) {
	curCnt := tr.cnts[key]
	if curCnt-num < 0 {
		panic("too many elements is deleted!")
	}

	tr.cnts[key] -= num
}

func (tr *Treap) insert(t *Node, key T, priority int) *Node {
	// 葉に到達したら新しい節点を生成して返す
	if t == nil {
		node := new(Node)
		node.key, node.priority = key, priority
		return node
	}

	// 重複したkeyは無視
	if key == t.key {
		return t
	}

	if tr.less(key, t.key) {
		// 左の子へ移動
		t.left = tr.insert(t.left, key, priority) // 左の子へのポインタを更新
		// 左の子の方が優先度が高い場合右回転
		if t.priority < t.left.priority {
			t = tr.rightRotate(t)
		}
	} else {
		// 右の子へ移動
		t.right = tr.insert(t.right, key, priority) // 右の子へのポインタを更新
		if t.priority < t.right.priority {
			// 右の子の方が優先度が高い場合左回転
			t = tr.leftRotate(t)
		}
	}

	return t
}

// 削除対象の節点を回転によって葉まで移動させた後に削除する
func (tr *Treap) delete(t *Node, key T) *Node {
	if t == nil {
		return nil
	}

	// 削除対象を検索
	if key == t.key {
		// 削除対象を発見、葉ノードとなるように回転を繰り返す
		return tr._delete(t, key)
	} else if tr.less(key, t.key) {
		t.left = tr.delete(t.left, key)
	} else {
		t.right = tr.delete(t.right, key)
	}

	return t
}

// 削除対象の節点の場合
func (tr *Treap) _delete(t *Node, key T) *Node {
	if t.left == nil && t.right == nil {
		// 葉の場合
		return nil
	} else if t.left == nil {
		// 右の子のみを持つ場合は左回転
		t = tr.leftRotate(t)
	} else if t.right == nil {
		// 左の子のみを持つ場合は右回転
		t = tr.rightRotate(t)
	} else {
		// 優先度が高い方を持ち上げる
		if t.left.priority > t.right.priority {
			t = tr.rightRotate(t)
		} else {
			t = tr.leftRotate(t)
		}
	}

	return tr.delete(t, key)
}

func (tr *Treap) rightRotate(t *Node) *Node {
	s := t.left
	t.left = s.right
	s.right = t
	return s
}

func (tr *Treap) leftRotate(t *Node) *Node {
	s := t.right
	t.right = s.left
	s.left = t
	return s
}

// rootからスタートする
func (tr *Treap) biggerLowerBound(t *Node, x T) *Node {
	if t == nil {
		return nil
	} else if tr.less(t.key, x) {
		// 探索キーxが現在のノードキーより大きい場合、右を探索する
		return tr.biggerLowerBound(t.right, x)
	} else {
		// 探索キーxが現在のノードキー以下の場合、左を探索する
		node := tr.biggerLowerBound(t.left, x)
		if node != nil {
			return node
		} else {
			return t
		}
	}
}

// rootからスタートする
func (tr *Treap) biggerUpperBound(t *Node, x T) *Node {
	if t == nil {
		return nil
	} else if tr.less(t.key, x) || t.key == x {
		// 探索キーxが現在のノードキー以上の場合、右を探索する
		return tr.biggerUpperBound(t.right, x)
	} else {
		// 探索キーxが現在のノードキーより小さい場合、左を探索する
		node := tr.biggerUpperBound(t.left, x)
		if node != nil {
			return node
		} else {
			return t
		}
	}
}

// rootからスタートする
func (tr *Treap) smallerUpperBound(t *Node, x T) *Node {
	if t == nil {
		return nil
	} else if tr.less(t.key, x) || t.key == x {
		node := tr.smallerUpperBound(t.right, x)
		if node != nil {
			return node
		} else {
			return t
		}
	} else {
		return tr.smallerUpperBound(t.left, x)
	}
}

// rootからスタートする
func (tr *Treap) smallerLowerBound(t *Node, x T) *Node {
	if t == nil {
		return nil
	} else if tr.less(t.key, x) {
		node := tr.smallerLowerBound(t.right, x)
		if node != nil {
			return node
		} else {
			return t
		}
	} else {
		return tr.smallerLowerBound(t.left, x)
	}
}

func (tr *Treap) inorder(u *Node, res *[]T) {
	if u == nil {
		return
	}
	tr.inorder(u.left, res)
	*res = append(*res, u.key)
	tr.inorder(u.right, res)
}

func (tr *Treap) preorder(u *Node, res *[]T) {
	if u == nil {
		return
	}
	*res = append(*res, u.key)
	tr.preorder(u.left, res)
	tr.preorder(u.right, res)
}

/*******************************************************************/

/*********** I/O ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
)

func newReadString(ior io.Reader, sf bufio.SplitFunc) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+9)) // for Codeforces
	r.Split(sf)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

// ReadInt returns an integer.
func ReadInt() int {
	return int(readInt64())
}
func ReadInt2() (int, int) {
	return int(readInt64()), int(readInt64())
}
func ReadInt3() (int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64())
}
func ReadInt4() (int, int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64()), int(readInt64())
}

// ReadInt64 returns as integer as int64.
func ReadInt64() int64 {
	return readInt64()
}
func ReadInt64_2() (int64, int64) {
	return readInt64(), readInt64()
}
func ReadInt64_3() (int64, int64, int64) {
	return readInt64(), readInt64(), readInt64()
}
func ReadInt64_4() (int64, int64, int64, int64) {
	return readInt64(), readInt64(), readInt64(), readInt64()
}

func readInt64() int64 {
	i, err := strconv.ParseInt(ReadString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

// ReadIntSlice returns an integer slice that has n integers.
func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}

// ReadInt64Slice returns as int64 slice that has n integers.
func ReadInt64Slice(n int) []int64 {
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt64()
	}
	return b
}

// ReadFloat64 returns an float64.
func ReadFloat64() float64 {
	return float64(readFloat64())
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(ReadString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// ReadFloatSlice returns an float64 slice that has n float64.
func ReadFloat64Slice(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadFloat64()
	}
	return b
}

// ReadRuneSlice returns a rune slice.
func ReadRuneSlice() []rune {
	return []rune(ReadString())
}

/*********** Debugging ***********/

// ZeroPaddingRuneSlice returns binary expressions of integer n with zero padding.
// For debugging use.
func ZeroPaddingRuneSlice(n, digitsNum int) []rune {
	sn := fmt.Sprintf("%b", n)

	residualLength := digitsNum - len(sn)
	if residualLength <= 0 {
		return []rune(sn)
	}

	zeros := make([]rune, residualLength)
	for i := 0; i < len(zeros); i++ {
		zeros[i] = '0'
	}

	res := []rune{}
	res = append(res, zeros...)
	res = append(res, []rune(sn)...)

	return res
}

// Strtoi is a wrapper of strconv.Atoi().
// If strconv.Atoi() returns an error, Strtoi calls panic.
func Strtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(errors.New("[argument error]: Strtoi only accepts integer string"))
	} else {
		return i
	}
}

// PrintIntsLine returns integers string delimited by a space.
func PrintIntsLine(A ...int) string {
	res := []rune{}

	for i := 0; i < len(A); i++ {
		str := strconv.Itoa(A[i])
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}

// PrintIntsLine returns integers string delimited by a space.
func PrintInts64Line(A ...int64) string {
	res := []rune{}

	for i := 0; i < len(A); i++ {
		str := strconv.FormatInt(A[i], 10) // 64bit int version
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}

// PrintfDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintfDebug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// PrintfBufStdout is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func PrintfBufStdout(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}
