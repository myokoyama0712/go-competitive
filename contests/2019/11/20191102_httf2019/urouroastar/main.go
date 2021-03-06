package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

/*********** I/O ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
)

func init() {
	ReadString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	// r.Buffer(make([]byte, 1024), int(1e+11)) // for AtCoder
	r.Buffer(make([]byte, 1024), int(1e+9)) // for Codeforces
	// Split sets the split function for the Scanner. The default split function is ScanLines.
	// Split panics if it is called after scanning has started.
	r.Split(bufio.ScanWords)

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

/*********** DP sub-functions ***********/

// ChMin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

// ChMax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
}

// NthBit returns nth bit value of an argument.
// n starts from 0.
func NthBit(num, nth int) int {
	return num >> uint(nth) & 1
}

// OnBit returns the integer that has nth ON bit.
// If an argument has nth ON bit, OnBit returns the argument.
func OnBit(num, nth int) int {
	return num | (1 << uint(nth))
}

// OffBit returns the integer that has nth OFF bit.
// If an argument has nth OFF bit, OffBit returns the argument.
func OffBit(num, nth int) int {
	return num & ^(1 << uint(nth))
}

// PopCount returns the number of ON bit of an argument.
func PopCount(num int) int {
	res := 0

	for i := 0; i < 70; i++ {
		if ((num >> uint(i)) & 1) == 1 {
			res++
		}
	}

	return res
}

/*********** Arithmetic ***********/

// Max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Max(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m < integer {
			m = integer
		}
	}
	return m
}

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m > integer {
			m = integer
		}
	}
	return m
}

// Sum returns multiple integers sum.
func Sum(integers ...int) int {
	s := 0

	for _, i := range integers {
		s += i
	}

	return s
}

// PowInt is integer version of math.Pow
// PowInt calculate a power by Binary Power (二分累乗法(O(log e))).
func PowInt(a, e int) int {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}

	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := PowInt(a, halfE)
		return half * half
	}

	return a * PowInt(a, e-1)
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Gcd returns the Greatest Common Divisor of two natural numbers.
// Gcd only accepts two natural numbers (a, b >= 1).
// 0 or negative number causes panic.
// Gcd uses the Euclidean Algorithm.
func Gcd(a, b int) int {
	if a <= 0 || b <= 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}
	if a < b {
		a, b = b, a
	}

	// Euclidean Algorithm
	for b > 0 {
		div := a % b
		a, b = b, div
	}

	return a
}

// Lcm returns the Least Common Multiple of two natural numbers.
// Lcd only accepts two natural numbers (a, b >= 1).
// 0 or negative number causes panic.
// Lcd uses the Euclidean Algorithm indirectly.
func Lcm(a, b int) int {
	if a <= 0 || b <= 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	// a = a'*gcd, b = b'*gcd, a*b = a'*b'*gcd^2
	// a' and b' are relatively prime numbers
	// gcd consists of prime numbers, that are included in a and b
	gcd := Gcd(a, b)

	// not (a * b / gcd), because of reducing a probability of overflow
	return (a / gcd) * b
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
		// str := strconv.FormatInt(A[i], 10)  // 64bit int version
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}

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

const MOD = 1000000000 + 7
const ALPHABET_NUM = 26
const INF_INT64 = math.MaxInt64
const INF_BIT60 = 1 << 60

var n, m, b int         // n: 座標のサイズ, m: マシーンの数, b: ブロックの数
var Blocks [40][40]bool // ブロックが設置あるか否か
var goal coord          // ゴール座標構造体
var Machines []machine  // マシーン構造体
var IsSet [40][40]bool  // 矢印を設置したか否か

type coord struct {
	y, x int
}

type machine struct {
	y, x int
	dir  rune
}

type direction struct {
	y, x int
	dir  rune
}

func main() {
	n, m, b = ReadInt3()
	gy, gx := ReadInt2()
	goal = coord{y: gy, x: gx}
	for i := 0; i < m; i++ {
		ry, rx := ReadInt2()
		tmp := ReadRuneSlice()
		c := tmp[0]

		Machines = append(Machines, machine{y: ry, x: rx, dir: c})
	}
	for i := 0; i < b; i++ {
		by, bx := ReadInt2()
		Blocks[by][bx] = true
	}

	bruteForceAStar()
}

func bruteForceAStar() {
	outputs := []direction{}

	for _, machine := range Machines {
		// マシーンごとに最初にぶつかる壁まで移動させる
		my, mx := machine.y, machine.x
		for {
			if machine.dir == 'L' {
				nmx := correctCoord(mx - 1)
				if Blocks[my][nmx] {
					break
				} else {
					mx = nmx
				}
			} else if machine.dir == 'R' {
				nmx := correctCoord(mx + 1)
				if Blocks[my][nmx] {
					break
				} else {
					mx = nmx
				}
			} else if machine.dir == 'U' {
				nmy := correctCoord(my - 1)
				if Blocks[nmy][mx] {
					break
				} else {
					my = nmy
				}
			} else if machine.dir == 'D' {
				nmy := correctCoord(my + 1)
				if Blocks[nmy][mx] {
					break
				} else {
					my = nmy
				}
			}
		}

		revRoute := AStar(my, mx)

		befDir := 'x'
		for i := len(revRoute) - 2; i >= 0; i-- {
			cy, cx := revRoute[i+1].y, revRoute[i+1].x
			ny, nx := revRoute[i].y, revRoute[i].x
			if !IsSet[cy][cx] {
				if (cy-ny == 1 || (cy == 0 && ny == n-1)) && befDir != 'U' {
					IsSet[cy][cx] = true
					outputs = append(outputs, direction{y: cy, x: cx, dir: 'U'})
					befDir = 'U'
				} else if (cy-ny == -1 || (cy == n-1 && ny == 0)) && befDir != 'D' {
					IsSet[cy][cx] = true
					outputs = append(outputs, direction{y: cy, x: cx, dir: 'D'})
					befDir = 'D'
				} else if (cx-nx == 1 || (cx == 0 && nx == n-1)) && befDir != 'L' {
					IsSet[cy][cx] = true
					outputs = append(outputs, direction{y: cy, x: cx, dir: 'L'})
					befDir = 'L'
				} else if (cx-nx == -1 || (cx == n-1 && nx == 0)) && befDir != 'R' {
					IsSet[cy][cx] = true
					outputs = append(outputs, direction{y: cy, x: cx, dir: 'R'})
					befDir = 'R'
				}
			} else {
				break
			}
		}
	}

	fmt.Printf("%d\n", len(outputs))
	for i := 0; i < len(outputs); i++ {
		fmt.Printf("%d %d %c\n", outputs[i].y, outputs[i].x, outputs[i].dir)
	}
}

type Node struct {
	pri          int // スコアと同等にセットする
	y, x         int // 座標
	status       int // 0: None -> 1: Open -> 2: Close
	rcost, hcost int // 実コスト、推定コスト
	parent       *Node
	hasParent    bool // 親ノードを持っているかどうかのフラグ
}

func NewNode(y, x, rcost int, parent *Node, hasParent bool) *Node {
	node := new(Node)

	node.y, node.x = y, x
	node.status = 0
	node.rcost = rcost
	node.hcost = calcHeulisticCost(y, x)
	node.pri = node.rcost + node.hcost
	node.parent = parent
	node.hasParent = hasParent

	return node
}

var nodeStatus [40][40]int

type NodePQ []*Node // Openリストとして用いる

func (pq NodePQ) Len() int           { return len(pq) }
func (pq NodePQ) Less(i, j int) bool { return pq[i].pri < pq[j].pri } // <: ASC, >: DESC
func (pq NodePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *NodePQ) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}
func (pq *NodePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// how to use
// temp := make(NodePQ, 0, 100000+1)
// pq := &temp
// heap.Init(pq)
// heap.Push(pq, &Node{pri: intValue})
// popped := heap.Pop(pq).(*Node)

// ゴールまでの推定コストを求める
func calcHeulisticCost(y, x int) int {
	return AbsInt(y-goal.y) + AbsInt(x-goal.x)
}

// ある始点からA*アルゴリズムでゴールまでの経路を求める
// ゴールまでの経路の逆順で返される
// スライス長が0の場合はゴールには到達不能
func AStar(sy, sx int) []coord {
	// 初期化
	temp := make(NodePQ, 0)
	openList := &temp
	heap.Init(openList)
	nodeStatus = [40][40]int{}
	delta := [][]int{
		[]int{0, -1}, []int{0, 1}, []int{-1, 0}, []int{1, 0},
	}

	snode := NewNode(sy, sx, 0, nil, false) // スタートノード
	heap.Push(openList, snode)
	nodeStatus[sy][sx] = 1 // Open

	var gnode *Node
	gflag := false
	for openList.Len() > 0 {
		currentNode := heap.Pop(openList).(*Node)
		ccost := currentNode.rcost
		cy, cx := currentNode.y, currentNode.x

		// ゴールノードが取り出せたらループを抜けて終了する
		if cy == goal.y && cx == goal.x {
			gnode = currentNode
			gflag = true
			break
		}

		// 移動可能なマスノードをOpenする
		for _, d := range delta {
			// 壁を取り払う
			ny, nx := correctCoord(cy+d[0]), correctCoord(cx+d[1])

			// ブロックに向かっても遷移できない
			if Blocks[ny][nx] {
				continue
			}

			// statusがNoneの場合のみOpenにして追加する
			if nodeStatus[ny][nx] == 0 {
				nnode := NewNode(ny, nx, ccost+1, currentNode, true)
				heap.Push(openList, nnode)
				nodeStatus[ny][nx]++ // statusをNoneからOpenへ更新
			}
		}

		// 注目中のOpenノードのstatusをClosedに更新
		nodeStatus[cy][cx]++
	}

	// ゴールに到達できたならば、その経路を返す
	if gflag {
		res := []coord{}
		cnode := gnode
		// for cnode.hasParent {
		for {
			res = append(res, coord{y: cnode.y, x: cnode.x})
			if cnode.hasParent {
				cnode = cnode.parent
			} else {
				break
			}
		}
		return res
	}
	return []coord{}
}

func correctCoord(xy int) int {
	if xy == n {
		return 0
	} else if xy == -1 {
		return n - 1
	}
	return xy
}

// ----------------

// for XorShift
var _gtx, _gty, _gtz, _gtw = 123456789, 362436069, 521288629, 88675123

// XorShiftによる乱数生成
// 下記URLを参考
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func randInt() int {
	tt := (_gtx ^ (_gtx << 11))
	_gtx = _gty
	_gty = _gtz
	_gtz = _gtw
	_gtw = (_gtw ^ (_gtw >> 19)) ^ (tt ^ (tt >> 8))
	return _gtw
}

func next(x int) int {
	if x == n-1 {
		return 0
	}
	return x
}
func bef(x int) int {
	if x == 0 {
		return n - 1
	}
	return x
}
