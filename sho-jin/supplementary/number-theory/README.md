# 整数論

Last Change: 2020-12-31 15:17:20.

## 参考

- [拡張ユークリッドの互除法〜一次不定方程式 `ax + by = c` の解き方〜](https://qiita.com/drken/items/b97ff231e43bce50199a)
- [AtCoder版！マスター・オブ・整数（素因数分解編）](https://qiita.com/drken/items/a14e9af0ca2d857dad23)

## ユークリッドの互除法

### ユークリッドの互除法周りで覚えるべきこと

以下の3つを覚えてしまう。

1. `gcd(a, b) = gcd(a - q * b, b)`
2. `gcd(a, b) = gcd(a-b, b)`
3. `a = c (mod b) => gcd(a, b) == gcd(c, b)`

1はユークリッドの互除法の本質そのもの。
問題のサイズが小さくできる。
概略としては `gcd(a, b) | gcd(a - q * b, b) && gcd(a - q * b, b) | gcd(a, b)` を示すことで導かれる。

2は1の `q = 1` の特殊ケース。

3はつまりは `a, c` の `b` で割ったときのあまりが等しいので、
互除法のアルゴリズムの最初のステップで同じ問題に帰着することから自然とわかる。

### 一次不定方程式（ベズー等式）の整数解を持つための必要十分条件

`a*x + b*y = c (a, b, c != 0)` が整数解を持つ `<=>` `c` が `gcd(a, b)` の倍数

`=>` はかんたんに示せるが、 `<=` は `a*x + b*y = gcd(a, b)` は整数解を持つことを示しておく必要がある。
これは、複数の示し方があるものの、いずれもテクニックが必要。

#### `a*x + b*y = gcd(a, b)` は整数解を持つことの証明

ちょっとむずかしいが、これが頭に入っていると競技のnumber theory系の問題に取り組みやすくなるので、
ぜひとも理解して覚えておきたい。

1. 拡張ユークリッドの互除法で具体的な解を導く
2. `a*x + b*y` で表される最小の正の整数を `d0` とおいて `d0 == gcd(a, b)` を示す
3. `a*x + b*y = 1` が整数解を持つ `<=>` `a, b` が互いに素: これを示し、利用する。

##### 1つ目の方法

拡張ユークリッドの互除法も、問題のサイズが小さくなる過程を理解することが重要。

`d = gcd(a, b), a = b*q + r` とする。
これらを元の方程式に代入すると、
`(b * q + r)*x + b*y = d <=> b*(q * x + y) + r*x = d`
から、 `(a, b) => (b, r)` のように、問題のサイズが小さくなっていることがわかる。

##### 2つ目の方法

難しい。
`d0 | gcd(a, b) && gcd(a, b) | d0` を導く。

##### 3つ目の方法

[高校数学の美しい物語](https://mathtrain.jp/axbyc)で説明されている。

`=>` は対偶がかんたんに示せるので、それを利用する。

`<=` が大事。
**`a, b` が互いに素なとき、 `a, 2a, 3a, ..., (b-1)a` を `b` で割ったあまりは全て異なるため、あまりが `1` となるものが存在する。**
よって `m*a = b*n + 1` と表され、ここから自然と導かれる。

※太線部分は背理法でかんたんに示せる。

※太線部分は「片方 `a` の整数倍を小さい方から列挙し、もう片方 `b` でそれらのあまりをとって並べると、
`b` 個あるあまりの種類が `b` 個並べられるまで衝突しない」というように、視覚的（？）に頭に思い描けると良い、かもしれない。。

※3つ目の命題が示せた後は、 `a = p*d, b = q*d (gcd(p, q) == 1)` とおくと、
もとの示したい命題はかんたんに示せる。

## 手書きメモ

![ユークリッドの互除法1](./images/euclid-1.jpg)

![ユークリッドの互除法2](./images/euclid-2.jpg)

![ユークリッドの互除法3](./images/euclid-3.jpg)

![ユークリッドの互除法4](./images/euclid-4.jpg)

![ユークリッドの互除法5](./images/euclid-5.jpg)

---

## オイラー関数

> 正の整数 `N` が与えられたとき、 `1, 2, ..., N` のうち `N` と互いに素であるものの個数を
> `Phi(N)` と表します。
> これをオイラー関数と呼びます。

（時間計算量的に）素因数分解が可能ならば求められる。

`N = p[0]^e[0] * p[1]^e[1] * ... * p[n-1]^e[n-1]` と素因数分解されるとき、
オイラー関数の値は、以下のようにして求まる。

`Phi(N) = N * (1-1/p[0]) * (1-1/p[1]) *...* (1-1/p[n-1])`

---

## 合同方程式

[ABC186-E](https://atcoder.jp/contests/abc186/tasks/abc186_e)で初めて問われた。

合同方程式 `a*x = b (mod m)` の解き方は[けんちょんさんの本問題の解説](https://drken1215.hatenablog.com/entry/2020/12/20/015100)
に詳しく書かれている。

※ この問題は中国剰余定理を応用することでも解けるらしいので、頭の片隅に入れておく。

### 解き方

`gcd(a, m) == 1` すなわち `a, m` が互いに素な場合。  
このとき、 `a` の逆元が求まれば `x = a^(-1)*b` が答え。  
`m` が素数でないときの逆元は拡張ユークリッドの互除法によって `extgcd(a, m)` とすることで求まる。

これは、不定方程式 `a*x + m*y = 1` の解を求めており、これの両辺を `m` であまりを取ると `a*x = 1 (mod m)` となる。
よって、このときの解 `x` が求めたい逆元となる。

`gcd(a, m) = g > 1` のときは少し複雑になる。  
結論から言うと、以下のような形式になる。

- `b` が `g` の倍数でないときは解が存在しない。
- `b` が `g` の倍数のときは `a', b', m' = a/g, b/g, m/g` として `a'*x = b' (mod m')` を解けば良い（ `gcd(a', m') == 1` なので先述の方法によって必ず解ける）

#### 証明

これの証明は順を追っていけばそれほど難しくない。

合同方程式を言い換えると `a*x - b` は `m` で割り切れると解釈できるため、
「 `a*x - b == m*y` を満たす `y` が存在する」という言い換えが可能。  
これはさらに `b == g * (a'*x - m'*y)` となることから、 `b` が `g` で割り切れるかどうか、という観点になる。  
割り切れないときはこれを満たす `x, y` は存在しないことは明らか。  
割り切れるときは `b = g*b'` となるため、もとの不定方程式は `a'*x - b' = m'*y` となる。  
これを改めて合同方程式の形に直すと `a'*x = b' (mod m')` が得られて先述の逆元を求めることによって解ける形になる。

### `extgcd` のインターフェースについて復習

`extgcd(a, b)` : `a*x + b*y = gcd(a, b)` という不定方程式について解くことで、

- `gcd(a, b)`
- 不定方程式のある解 `(x, y)`

を一挙に求めるもの、という認識を持っておく。


