/*
URL:
https://atcoder.jp/contests/joi2010ho/tasks/joi2010ho_b
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// var (
// 	n int
// 	A []int

// 	dp [10000 + 5][5000 + 5][2]int
// )

// const (
// 	ME, YOU = 0, 1
// )

// func main() {
// 	defer stdout.Flush()

// 	n = readi()
// 	A = readis(n - 1)

// 	for i := 0; i <= 10000; i++ {
// 		for j := 0; j <= 5000; j++ {
// 			for k := 0; k < 2; k++ {
// 				dp[i][j][k] = INF_B30
// 			}
// 		}
// 	}

// 	dp[0][0][0] = 0
// 	dp[0][0][1] = 0
// 	for i := 0; i < n-1; i++ {
// 		cost := A[i]
// 		for j := 0; j <= n/2; j++ {
// 			// 自分→自分
// 			ChMin(&dp[i+1][j+1][ME], dp[i][j][ME])
// 			// 相手→相手
// 			ChMin(&dp[i+1][j][YOU], dp[i][j][YOU])
// 			// 自分→相手
// 			ChMin(&dp[i+1][j+1][YOU], dp[i][j][ME]+cost)
// 			// 相手→自分
// 			ChMin(&dp[i+1][j][ME], dp[i][j][YOU]+cost)
// 		}
// 	}

// 	fmt.Println(Min(dp[n-1][n/2][ME], dp[n-1][n/2][YOU]))
// }

var (
	n int
	A []int

	dp [2][5000 + 5][2]int
)

const (
	ME, YOU = 0, 1
)

func main() {
	defer stdout.Flush()

	n = readi()
	A = readis(n - 1)

	for i := 0; i < 2; i++ {
		for j := 0; j <= 5000; j++ {
			for k := 0; k < 2; k++ {
				dp[i][j][k] = INF_B30
			}
		}
	}

	dp[0][1][ME] = 0
	dp[0][0][YOU] = 0
	for i := 0; i < n-1; i++ {
		cost := A[i]
		for j := 0; j <= n/2; j++ {
			// 自分→自分
			ChMin(&dp[1][j+1][ME], dp[0][j][ME])
			// 相手→相手
			ChMin(&dp[1][j][YOU], dp[0][j][YOU])

			// 自分→相手
			ChMin(&dp[1][j][YOU], dp[0][j][ME]+cost)
			// 相手→自分
			ChMin(&dp[1][j+1][ME], dp[0][j][YOU]+cost)
		}

		for j := 0; j <= n/2; j++ {
			for k := 0; k < 2; k++ {
				dp[0][j][k] = dp[1][j][k]
				dp[1][j][k] = INF_B30
			}
		}
	}

	fmt.Println(Min(dp[0][n/2][ME], dp[0][n/2][YOU]))
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

/*******************************************************************/

/********** common constants **********/

const (
	MOD = 1000000000 + 7
	// MOD          = 998244353
	ALPH_N  = 26
	INF_I64 = math.MaxInt64
	INF_B60 = 1 << 60
	INF_I32 = math.MaxInt32
	INF_B30 = 1 << 30
	NIL     = -1
	EPS     = 1e-10
)

// modi can calculate a right residual whether value is positive or negative.
func modi(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// modll can calculate a right residual whether value is positive or negative.
func modll(val, m int64) int64 {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

/********** FAU standard libraries **********/

//fmt.Sprintf("%b\n", 255) 	// binary expression

/********** I/O usage **********/

//str := reads()
//i := readi()
//X := readis(n)
//S := readrs()
//a := readf()
//A := readfs(n)

//str := ZeroPaddingRuneSlice(num, 32)
//str := PrintIntsLine(X...)

/*********** Input ***********/

var (
	// reads returns a WORD string.
	reads  func() string
	stdout *bufio.Writer
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

// readi returns an integer.
func readi() int {
	return int(_readInt64())
}
func readi2() (int, int) {
	return int(_readInt64()), int(_readInt64())
}
func readi3() (int, int, int) {
	return int(_readInt64()), int(_readInt64()), int(_readInt64())
}
func readi4() (int, int, int, int) {
	return int(_readInt64()), int(_readInt64()), int(_readInt64()), int(_readInt64())
}

// readll returns as integer as int64.
func readll() int64 {
	return _readInt64()
}
func readll2() (int64, int64) {
	return _readInt64(), _readInt64()
}
func readll3() (int64, int64, int64) {
	return _readInt64(), _readInt64(), _readInt64()
}
func readll4() (int64, int64, int64, int64) {
	return _readInt64(), _readInt64(), _readInt64(), _readInt64()
}

func _readInt64() int64 {
	i, err := strconv.ParseInt(reads(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

// readis returns an integer slice that has n integers.
func readis(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = readi()
	}
	return b
}

// readlls returns as int64 slice that has n integers.
func readlls(n int) []int64 {
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		b[i] = readll()
	}
	return b
}

// readf returns an float64.
func readf() float64 {
	return float64(_readFloat64())
}

func _readFloat64() float64 {
	f, err := strconv.ParseFloat(reads(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// ReadFloatSlice returns an float64 slice that has n float64.
func readfs(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = readf()
	}
	return b
}

// readrs returns a rune slice.
func readrs() []rune {
	return []rune(reads())
}

/*********** Output ***********/

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

// Printf is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func printf(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}

/*********** Debugging ***********/

// debugf is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func debugf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

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
