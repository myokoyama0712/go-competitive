package main

import (
	"bufio"
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
	r.Buffer(make([]byte, 1024), int(1e+11))
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

// DigitSum returns digit sum of a decimal number.
// DigitSum only accept a positive integer.
func DigitSum(n int) int {
	if n < 0 {
		return -1
	}

	res := 0

	for n > 0 {
		res += n % 10
		n /= 10
	}

	return res
}

// DigitNumOfDecimal returns digits number of n.
// n is non negative number.
func DigitNumOfDecimal(n int) int {
	res := 0

	for n > 0 {
		n /= 10
		res++
	}

	return res
}

// Sum returns multiple integers sum.
func Sum(integers ...int) int {
	s := 0

	for _, i := range integers {
		s += i
	}

	return s
}

// Kiriage returns Ceil(a/b)
// a >= 0, b > 0
func Kiriage(a, b int) int {
	return (a + (b - 1)) / b
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

// Strtoi is a wrapper of `strconv.Atoi()`.
// If `strconv.Atoi()` returns an error, Strtoi calls panic.
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

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/all/ALDS1_8_D

type Node struct {
	key, priority       int
	right, left, parent *Node
}

var root *Node

func rightRotate(t *Node) *Node {
	s := t.left
	t.left = s.right
	s.right = t
	return s
}

func leftRotate(t *Node) *Node {
	s := t.right
	t.right = s.left
	s.left = t
	return s
}

func Insert(t *Node, key, priority int) *Node {
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

	if key < t.key {
		// 左の子へ移動
		t.left = Insert(t.left, key, priority) // 左の子へのポインタを更新
		// 左の子の方が優先度が高い場合右回転
		if t.priority < t.left.priority {
			t = rightRotate(t)
		}
	} else {
		// 右の子へ移動
		t.right = Insert(t.right, key, priority) // 右の子へのポインタを更新
		if t.priority < t.right.priority {
			// 右の子の方が優先度が高い場合左回転
			t = leftRotate(t)
		}
	}

	return t
}

// keyのノードが見つからなかった場合はnilを返す
func Find(k int) *Node {
	u := root
	for u != nil && k != u.key {
		if k < u.key {
			u = u.left
		} else {
			u = u.right
		}
	}
	return u
}

// 削除対象の節点を回転によって葉まで移動させた後に削除する
func delete(t *Node, key int) *Node {
	if t == nil {
		return nil
	}

	// 削除対象を検索
	if key < t.key {
		t.left = delete(t.left, key)
	} else if key > t.key {
		t.right = delete(t.right, key)
	} else {
		// 削除対象を発見、葉ノードとなるように回転を繰り返す
		return _delete(t, key)
	}

	return t
}

// 削除対象の節点の場合
func _delete(t *Node, key int) *Node {
	if t.left == nil && t.right == nil {
		// 葉の場合
		return nil
	} else if t.left == nil {
		// 右の子のみを持つ場合は左回転
		t = leftRotate(t)
	} else if t.right == nil {
		// 左の子のみを持つ場合は右回転
		t = rightRotate(t)
	} else {
		// 優先度が高い方を持ち上げる
		if t.left.priority > t.right.priority {
			t = rightRotate(t)
		} else {
			t = leftRotate(t)
		}
	}

	return delete(t, key)
}

// 削除のための回転処理が終わった後のdelete
func Delete(t *Node, key int) {
	var p *Node
	p = nil
	// 削除対象の葉ノードが見つかるまで下降する
	for t.key != key {
		p = t
		if key < t.key {
			t = t.left
		} else if key > t.key {
			t = t.right
		}
	}

	if p == nil {
		// 削除対象は根ノードなので、rootをnilにする
		root = nil
	} else {
		if p.right != nil && p.right.key == key {
			p.right = nil
		} else if p.left != nil && p.left.key == key {
			p.left = nil
		}
	}
}

func inorder(u *Node, res *[]int) {
	if u == nil {
		return
	}
	inorder(u.left, res)
	*res = append(*res, u.key)
	inorder(u.right, res)
}

func preorder(u *Node, res *[]int) {
	if u == nil {
		return
	}
	*res = append(*res, u.key)
	preorder(u.left, res)
	preorder(u.right, res)
}

func main() {
	n := ReadInt()

	for i := 0; i < n; i++ {
		s := ReadString()
		if s == "insert" {
			k, p := ReadInt2()
			// if root == nil {
			// 	root = Insert(root, k, p)
			// } else {
			// 	Insert(root, k, p)
			// }
			root = Insert(root, k, p)
		} else if s == "print" {
			res := []int{}
			inorder(root, &res)
			fmt.Printf(" ")
			fmt.Println(PrintIntsLine(res...))
			res = []int{}
			preorder(root, &res)
			fmt.Printf(" ")
			fmt.Println(PrintIntsLine(res...))
		} else if s == "find" {
			k := ReadInt()
			n := Find(k)
			if n != nil {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		} else if s == "delete" {
			k := ReadInt()
			root = delete(root, k)
			// node := delete(root, k)
			// if node != nil {
			// 	Delete(root, k)
			// }

			// res := []int{}
			// inorder(root, &res)
			// fmt.Printf(" ")
			// fmt.Println(PrintIntsLine(res...))
			// res = []int{}
			// preorder(root, &res)
			// fmt.Printf(" ")
			// fmt.Println(PrintIntsLine(res...))
		}
	}
}

// MODはとったか？
// 遷移だけじゃなくて最後の最後でちゃんと取れよ？

/*******************************************************************/
