# ABC027 過去問感想

Last Change: 2020-02-26 23:51:25.

## A問題（@2020-02-09）

出現した数字のカウントを持っておけば自然と導出できる。

サンプルが優しい。

## B問題（@2020-02-04）

結構考えた。
全探索すると `O(2^n)` とかになるので無理（DPとかできるのかな？）

考察を進めると、平均とその整数倍を、今見ている島までの累積和と比較することで、
橋をかけるべきかどうかの判断が可能となる。

貪欲法に分類される？

## C問題（@2020-02-26）

ゲーム系問題。

一回諦めて解法を観たが、完全二分木を使った解法がトリッキー過ぎて宇宙が広がったので、
ググってブログを探してみた。

maspyさんのブログを見ると、自分が当初考えた「定石通り、選択肢が限られた後ろから考える」という手法でも解けるらしいとわかったので、
改めて考察を進めた。

以下のように考えると行ける。

1. `2*x > n` となるような `x` が回ってくると負け。つまり `n/2 < x <= n`
2. 2つの選択肢のうち、いずれかが↑で求めた範囲に落とせるのであれば勝ち。つまり `(lb-1)/2 < x <= ub/2`
  - 「いずれか」なので数直線上のORを取るとこのような範囲となる（わからないならば紙に書いてみるといい）
3. 2つの選択肢の両方ともが↑で求めた範囲に落ちてしまうのであれば負け。つまり `lb/2 < x <= (ub-1)/2`
  - 「両方とも」なので数直線上のANDを取るとこのような範囲となる（わからないならばｒｙ）
4. 以下繰り返し。

初期値が `x=1` なので、これが負け範囲に落ちるのか勝ち範囲に落ちるのかを、 `lb==0` となるまで続けて調べれば良い。

ゲーム系は後ろから考えるのが定石だが、その勝ち負けの更新の振る舞いも改めて強く認識しておくと良い。
つまり、

- 複数選択肢があって、1つでも「負け範囲」に落とせるのならば、それは「勝ち確定」
- 複数選択肢があって、全てが「勝ち範囲」に落ちてしまうのならば、それは「負け確定」

という振る舞いを認識しておくと、つまらない考察ミスが減らせる（かもしれない）。

