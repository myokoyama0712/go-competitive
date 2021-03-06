# E8くんの100問セット

Last Change: 2020-09-04 15:10:40.

[Qiitaのこれ](https://qiita.com/e869120/items/eb50fdaece12be418faa#2-3-%E5%88%86%E9%87%8E%E5%88%A5%E5%88%9D%E4%B8%AD%E7%B4%9A%E8%80%85%E3%81%8C%E8%A7%A3%E3%81%8F%E3%81%B9%E3%81%8D%E9%81%8E%E5%8E%BB%E5%95%8F%E7%B2%BE%E9%81%B8-100-%E5%95%8F)
を解いていく。

## 001

制約の記載が不足している気がするが、普通にやれば通る。

## 002

約数列挙のライブラリの検証にでも。

## 003

今となっては難しいABC-Bのような気がする。

## 004

初めて解いたが、素直な全探索で良い。

## 005

ちょっとむずかしいが、ちゃんと模範解答通りだったので一安心。

## 006

この問題のように、「何について全探索するか？」というのを考えて工夫するのは大事。
当時参加していたときもちゃんと解けたが、当時よりもシンプルなコードが書けていたのは良い。

## 007: ★

方針建てるのもそうだが、バグらせたりして、かなり時間がかかってしまった。

3000個の格子点が与えられるので、その中から得られる正方形の内、最大のものを答える問題。

2点さえ決まれば、回転座標を利用してそこから正方形が作れるかどうかは判定できる。
。。と思ったが、実際には2点というよりは「あるベクトルが決まると確定する」という構造になっている（はず）。
まず、そこで全探索範囲をミスらないようにする必要がある。

こういった幾何の問題では、「自分が今どのような量を求めているのか？」というのを特に忘れないようにしたい。
さもなくば、迷子になってしまうことがある。

## 008: ★

これはかなり難しい気がする。
全探索というよりは、貪欲法のようにみえる。

正しく証明しようと思うと、絶対値の和の最小化を考える必要がある。
（スライドを読んでもこれの証明がよくわからず。。）

## 009

007に似ているが、こちらはだいぶ簡単だと思う。

## 010

ここからビット全探索。

クエリ形式なので、出てきた計算結果はメモしておく必要がある。

※この問題で、ktateishさんのbit全探索抽象化関数を拝借してみたが、非常に便利だった。

## 011

めんどうなだけで整理すると簡単。

## 012

抽象化していると、少し関数式に迷うぐらいで、簡単。

## 013

片方の軸のみビット全探索して、もう片方の軸は貪欲に考える、というのは、
ABCのE問題でもあった（チョコレートを最適に分割するやつ）。

ただし、実装はこちらのほうが圧倒的に簡単。

## 014

「ビット全探索である」と言われると簡単に思えてしまう問題な気がする。
前情報がなくても、極端な制約をみてちゃんときづけること！

## 015

順列全探索というカテゴリ。

浮動小数点が絡むのでめんどくさいが、やるだけ。

## 016

順列を辞書順に列挙するのをデフォルトにしておくと良い。

## 017: ★

手をつけるまでめちゃくちゃめんどくさく感じるが、計算量的に贅沢をすると結構シンプルになる。

構造体をキーとしたハッシュマップを使うと良い。

**※順列全列挙は `n=8` という制約のときに敏感になりたい。**
（`10! >= 3.5M` なので、 `n=10` ぐらいになるとむしろビット全探索のほうが有力になる。）

## 018

ここから二分探索。

この問題は「異なる数字の数」なので、二分探索でやろうとはあまりならない気がする。
しかし、めぐる式二分探索のような、抽象化した関数の練習には良い。

## 019: ★

二分探索という前情報がないと結構難しいかもしれない。
基本中の基本だが、二分探索は近傍のインデックスを求める手段だというのを、常に忘れずに。

円環上の問題は少しやっていて怖いが、この問題は一発で通せたので、それほどでもないけど。
とはいえ、紙の上での考察がないとちょっとつらい。

## 023

半分全列挙。

`n=40` でビット全探索＋半分全列挙みたいなのが主流だと思っていたが、こういうのもあるので注意したい。

## 024: ★

ここからDFS（深さ優先探索）。

この問題は基本中の基本だが、意外と普通の競プロの問題よりも正しく実装するのが大変。
ポインタ引数を用いるよりも、競プロ的にはグローバル変数に頼ってしまうほうが簡単かも。

「有向グラフのDFSは意外と大変」というのはちゃんと覚えておきたい。

## 025: ★

蟻本でもある、2次元グリッドの連結成分をDFSで求めるもの。

union findも重要な手法だが、DFSも問題に合わせて素早く、正しく実装できるのが安心。

## 026

累積和＋DFSみたいな問題で、いわゆる教育的という感触がある。

簡単だが、ふと読んだPDFにかかれている、以下のevimaさんのアドバイスは有用かもしれない。

> このような木に関する問題では、木の代わりに `1 − 2 − 3 − . . . − N` という直線的なグラフを考えると助け
> になることがあり、今回はその好例です。

## 027: ★★

めっちゃ難しい気がする、主に計算量解析が。

企業コンであった、いちごのケーキを切り分ける問題っぽい再帰を行う。

復習したほうがよく、そのときは解説をちゃんと読んだほうがいいが、なぜこれで間に合うのかがよくわからない。。

## 028

ここからBFS（幅優先探索）。

この問題のように、最短距離の基本はBFSであると心得る。

## 029

グリッドをBFSする。

グリッドを隣接リストのグラフに変換するライブラリや、
隣接リストグラフのBFSによる単一始点最短経路ライブラリのverifyに最適な問題。

## 030

これも029と同様に、verifyに便利。

たかだか10回BFS全探索をするだけなので、ある意味愚直にやれば良い。

## 031: ★

少し実装が重いという前情報があったが、割と整理してきれいに書けたと思う。

グリッドのBFSの亜種だが、テーブル構造が特殊なので、グラフ化しようとせずに素直にBFSを組んだほうが早そう。
（グラフを作るのも悪くなさそうだけど。）

グリッドのBFSは、キューに突っ込んだ時点でチェック済みになることに注意！
ダイクストラの感覚でやってはいけない。
（あちらはキューに突っ込んだ後も更新の可能性があるから、キューから取り出したときにフラグ更新する。）

## 032: ★（もうやらなくていい）

迷路の組み方が特殊だが、めんどくさいのでもうやらなくて大丈夫。

※グリッドを隣接リストに変換する手法のありがたさがわかる問題。

## 033

今回が3回目だが、過去2回WAを出していたのに対して、今回はちゃんと一発で通せた。

ややこしい数え上げなどせずに、ちゃんと経路復元して素直に1つずつ丁寧に数えるのが安全。

## 034

ここからしばらくはDP。

フィボナッチ数も立派なDPの問題（数え上げDPに該当する？）。

一応、分類はナップザックDPらしい。

## 035

もっとも基本的なナップザックDP。

## 036: ★

個数制限なしのナップザック問題。

基本問題とはいうものの、今再復習しても結構難しいと思う。
DAGを正しくイメージして、無駄をなくせないか？という思考をする上での基礎としたい。

## 037: ★

036と設定や求める量が異なるだけで、状態の考え方は似ており、また遷移は全く同じ。

## 038: ★

LCSを簡単に思える日は来るのだろうか。。

というか、LCSはナップザックDPの分類らしい。
遷移において、片方のidxの差が1のものはすべてナップザックDPなんだろうか？

## 039

先の基本問題よりは簡単だと思うが、面白い。

JOIの問題は大変なものも多いが、問題文が面白いと思う。

## 040

ABCであった、AGCTの文字列を避けるDP？みたいな問題だと思った。

とはいえ、個人的にはこちらのほうがだいぶ簡単だと思う。

## 041

これも素直に状態定義を考えれば、正しい解法に達する事ができる。

※どうも「各日の服の候補のうち、 `C[j]` が最大もしくは最小のものの2通りのみを試す」という方法も正しいらしい。
当然、証明無しでやるのは危険。

## 042

簡単だけど、なぜか添字がごっちゃになる気がする。

## 043

簡単。

## 044: ★★

初めてこのシリーズでわからない問題が出てきた。

けんちょんさんの記事をみて気づいたが、所詮は「個数制限なしのナップザック問題」に過ぎなかった。

考える必要のある正四面体数の数はかなり小さいことに注意すれば気付ける（というかcube rootなのはすぐ気づいたのに。。）。

この問題はメモリ制約が厳しいので、DP配列のメモリを節約する必要があるので、その練習にもいいかもしれない。

## 045

問題設定は複雑に見えるが、素直にナップザックDPをやるだけでよい。

初期値の条件を見逃したせいで、めちゃくちゃWAに悩まされた。
油断せず問題文はしっかり読もう。

ここでナップザックDPは終わり。

## 046: ★

ここからは「区間DP」と呼ばれるものになる。

一番の基本形は「連鎖行列積」と呼ばれる、行列の積において整数の掛け算の回数を最小化する問題。
基本なのに結構難しい。。

まず、 `A[n][m] * A[m][l]` の行列の積の掛け算の回数は `n*m*l` となるのは、もはや知識としてしまう。

区間DPでは `dp[i][j] := i番目からj番目までの行列の積で回数最小のものを記憶` とする。
これは掛け算のパターンをすべて試せばわかり、区間 `(i, j)` が決まると、その間の `k` を全探索すれば良い。

よって、この区間DPは計算量が `O(N^3)` となる。

## 047: ★★★

ここまでで一番難しく、答えをみてしまったが、面白い。

区間DPとしてもいい問題だろうけど、それ以上に円環を扱う問題の対策を学べる問題としていいと思った。

模範解答のコードを読んでみたり、おまけの方法でやってみるのも学びがあるはず。
時間をおいてから復習したい。

**円環のインデックスをそのまま扱うのは、すぬけさんレベルでもめんどくさいらしい（？）。**
**まずは直線に直せないか？という考察を行いたい。**

## 048: ★★★

これもやはり難しい。

`O(N^3)` の区間DPだが、直感的に遷移がすべて網羅できているのかどうか？というのが不安に感じられてしまう気がする。

**「区間の除去」みたいな問題は区間DPの典型的な設定であるとのこと。**

- [けんちょんさんの解説](https://drken1215.hatenablog.com/entry/2020/03/10/160500)
- [kamiさん？のブログ](https://algo-logic.info/range-dp/)

## 049: ★

ここからはbitDP。

巡回セールスマン問題は基本中の基本だが、なぜか閉路の長さを問われて困惑してしまった。
雑に考えると、始点の選び方をすべて試す必要があり、bitDP部分と併せて考慮すると、
計算量が `O(N^3)` になってしまうように感じてしまった。

しかし、よくよく考えると、ある最短の閉路が存在する場合、それはどこを始点にとっても形成することができる。
よって、始点は例えば `0` の場合など、一つだけ考えればよい。

## 050: ★★

時間制限の部分は特に難しく考える必要はないので、あまり気にする必要はない。

それよりも、最短経路を取るものの経路数を数える部分が難しい（と思う）。

自分は最短経路を求めてから、それを復元する要領で考えた。
一発でACできたのはかなり上出来だと思う。

とはいえ、順方向の最短経路を求めるbitDPにおいて、
`dp[S][j] := {最短経路の距離, その時の経路数}, dp[0][0] = {0, 1}`
とすることで、順方向のみで解けるっぽい。

**ひらめき？: 重みを持った樹形図？（状態ごとに通り数をまとめると、そのような解釈ができる、気がする）**

## 051

こういうのはbitDPとナップザックDPの複合問題と言えるのだと思う。
（`i` 番目から `i+1` 番目を求めるときに `i-1` 番目以前の結果は関係ない。）

「最初はJくんが鍵を持っている」という条件を見逃して、めちゃくちゃ悩んでしまった。

## 052: ★★★

難しくてわからなかったために答えを観たが、わかってしまうと初歩的なテクニックの組み合わせだけで解けるため、
かなり言い問題だと感じた。

基本にしたがって、愚直な方法からの効率化を考える。

`m` 種類の並べ方は `m!` 通りあるが、これを決めてしまったときの高速な計算方法がわかれば、
さらに効率化するためのヒントが得られる可能性がある。

結論から言うと、これは累積和を使うと `O(m)` で可能となる。
ある区間内で並べたい種類のぬいぐるみ `i` を決めた場合、区間内に存在する `i` の個数を `N[i]` とし、
`i` の総数を `T[i]` とすれば、その区間で取り除くべきぬいぐるみの数は `T[i] - N[i]` でよいとわかる。
よって、全種類についてこれらを前計算しておけば、左から順番にぬいぐるみの種類ごとに加算していけば、答えがわかる。

これをbitDPを使うことで高速化してやれば良い！

`dp[S] := 左から並べるぬいぐるみの集合を決めたときの、その区間で除去するべきぬいぐるみの最小個数`

遷移もかなり簡単に考えられる。
次に並べたいぬいぐるみについて、上述した除去すべき数を計算して、それを加算する形で良い。

## 056

ここから最短経路（ダイクストラ法）のパート。

この問題は完全にピュアな非負の辺のみの単一始点最短経路の問題なので、ライブラリチェックなどに活用したい。

## 057

これもオンラインクエリの形で都度ダイクストラすれば解けてしまう。

なんなら、TLが10秒なので、毎回わーシャルフロイドしてしまってもいいらしい（解説にはそう書かれているが、言語に寄っては厳しそう）。

## 058

グラフの構築までに、色々なステップがあり、そういった意味では実装が重めでめんどくさい。
ただし、ひとつひとつは明快なので、それほど嫌な気持ちになる実装パートはない。
よって、易しめだと感じる。

解説が結構勉強になる。

危険な街の列挙パートでは、自分は多点BFSと呼ばれるものをやった。
一方解説では、「`k` 個のゾンビの街から1本の道路でつながる仮想の街 `n(0-based)` を仮定し、そこを始点としてBFSする」
という方法を提案している。
これは、最小全域木の問題でたまにやる手法であり、グラフの拡張手段の典型手法としていつでも想起できるようにしておきたい。

また、最後のダイクストラで答えを出すところで調整をしたが、
解説通り、 `n-1` のゴールの街に向かう辺のコストは0とすべきだった。

## 059

これも058とほぼ同じだが、今回はエッジの数が大きく、そのままダイクストラするとMLEしてしまった。
要所要所で `int32` にしてもほとんど有効にならなかったあたり、おそらくpriority queueの影響が大きいのかもしれない。

なので `O(N^2)` のダイクストラを取り出したが、これでも `int64` ではMLEしてしまった。
これを `int32` に書き換えたところ、なんとかギリギリACだった。

`10^6` ぐらいになってくると、TLEのほかMLEにも気をつける必要があるかもしれない。
**512MBや1GBなら大丈夫だと思うが、256MBのときはかなり用心する必要がある。**

## 064

ここからは最小全域木。

この問題は最も基本的なので、クラスカル法等のライブラリチェック用にすると良い。

## 065: ★★

**これは最小全域木の重要な典型問題（有名問題）。**
過去に何回か似たようなものに出会ったのに解けなかったので、反省。

まず、以下は関連問題。

- [ABC173-F](https://atcoder.jp/contests/abc173/tasks/abc173_f)
- [APC001-D](https://atcoder.jp/contests/apc001/tasks/apc001_d)

**典型: `連結成分数 = 頂点数 - 辺の数` （ただし、木・森グラフに限定される）**

クラスカル法の途中の状態というのは、連結成分を一つずつ減らしながら作られる森の
コストの総和が最適な状態と言えるので、全体の連結成分数が `k` になった段階で止めて出力すれば良い。

※この点からも、最小全域木はやはりクラスカル法のほうが汎用性があるような。。？

## 066

これはグラフの構築パートがめんどくさいだけで、それが終わればあとはクラスカルやるだけになる。

強いて言うなら、、

- costがfloat64のクラスカルが必要
- フォーマット指定子によって小数点以下3桁までに出力を制限する必要がある

ぐらいが初めてのポイントだったが、大したことではないので、やはり改めて解き直す必要はないと思う。

## 068

これと次の問題の2問だけだが、「高速な素数判定法」というカテゴリ。

この問題は、試し割り法によって素因数分解を行うだけで良い。

基本問題なだけあって、出力が普段のコンテストのものとは違って、ちょっと特殊。

## 069

これも普通に `O(sqrt(n))` の素数判定を使い、愚直に全部調べれば良い。

## 070

この問題と次の問題の2問だけだが、「高速なべき乗計算」というカテゴリ。
要は「二分累乗法」のこと。

これは本当に1つのべき乗計算をやらせる基本問題。

## 071

二分累乗法と、剰余の基本性質を使う。

累積和の部分は苦手なタイプだったが、特にバグらせなかったので、その点は偉い。

## 076

ここからはしばらく累積和のカテゴリ。

この問題は累積和の基本問題。
半開区間に慣れるべし。

## 077

一見めんどくさそうだが、自然に累積和を半開区間で扱えるように定義すると、非常にシンプルに書くことができる。

## 078

2次元累積の基本的な問題。

制約が厳しいが、自分のライブラリでも、特に気を使う必要はなくすんなり通せた。

## 080

これも2次元累積和が使えれば特に苦労しない。

**2次元累積和の問題は計算量見積もりに注意したい。**

例えばこの問題は `h <= 125, w <= 125` で、すべての長方形領域を全探索することになり、
大体 `10^8` 程度の計算回数になるが、かなり高速に計算できる（300msで収まる）。
問題に寄っては丁寧な枝刈りが要求されるものもあるかもしれないが、フルフィードバックならとりあえず投げてしまうのも悪くない（多分）。

## 081

ここからしばらくは「いもす法」のカテゴリ。

これは最も基本になる問題。

## 082

問題文が冗長で苦しいが、普通に時刻の始点と終点をエンコードすれば、あとはいもす法部分はやるだけになる。
終点の扱いは注意が必要だが、別に難しくはない。

## 083

これも整理するのが少し面倒だが、整理し終わった後にやるいもす法は082のものと似ている。

## 084

## 088: ★

ここから2問は「圧縮」のカテゴリ（※座標圧縮ではない）。

この問題はスタックを用いたシミュレーションを行えば良い。
個数を圧縮してやるのがポイント。

## 089: ★★★

フリップも関わる、個人的重要問題。

難しくてわからなかったので、解答PDFを読んだ。

考察を進めると、以下の点が重要となる。

- ある区間を反転する場合、その区間を完全に包含し、かつ最大長の交互列のみを考えれば良い。
  - 反転区間を包含しない区間を考える代わりに、別の反転区間を完全に包含し、かつ答えが一緒になるような交互列の見方があるため。
- 反転する区間は、交互列になっている部分だけを考えれば良い。
  - 1つ目の考察と合わせると、連続部分を反転させても依然として連続になるため、そのような反転区間は考えなくて良くなる。
  - この時点で、考えられる候補を全探索するアルゴリズムは `O(N^2)` にまで落ちている。
    - これを思いつくだけでもけっこう大変だと思う。
- 交互列の長さを数えるには、反転したところから左右にさらに元の交互列をのぼせば良い。
  - よって、満点解法は「元の数列を交互列の長さに圧縮し、最大3つの要素を取り出してその和を計算したものの最大値」として求められる。

**フリップ系問題の典型問題の一つとして覚えておきたい。**

## 090

ここから2問は「単純な幾何計算」のカテゴリ。

この問題は二分探索が必要かと思いきや、別にすべてのペア間の距離を求めれば、そこから答えを計算することができた。

むしろ、二分探索する場合は境界条件を注意しないとWAになってしまった。
すなわち、接するのはOKなので、二点間の距離と半径の和が等しいときは「重ならない」という判定関数の元で二分探索しないと、
あるケースでWAになってしまう。

※二分探索の回数を増やせば大丈夫かと思ったが、1000回にしても駄目だった。
幾何問題で二分探索するときは慎重にやらないと駄目っぽい。

## 092

ここから3問は「考察に比べて実装がとても重い問題」カテゴリ。

この問題は、シンプルな落ちものパズルゲームのシミュレーションを行う問題。

以下の方針で一発ACできた（ちょっと時間かかり過ぎだが）。

- 落下のプロセスは、各マスの下の部分から処理し、再帰関数で実装する。
- 行の消滅プロセスでは、Union Findを活用して削除すべきマスを特定する。
  - [自分のブログのこの記事](https://maguroguma.hatenablog.com/entry/2019/12/17/003000#d-remove-one-element)を思い出したい。

今回のUnion Findの使い方は汎用性が高いので、よく覚えておきたい。

思いかもしれないが、あまり嫌な気持ちになる実装は求められなかった。

## 093

これも092とほとんど同じ問題。双子くんが解説を作ってくれているので、こちらをやるほうが良さそう。

落とす部分のシミュレーションは、単純に新しいスライスを作ってappendしていくだけで良かったっぽい。

計算量解析がむずい（特に「実は `k <= 3` のケースしか考えなくていいです」の部分とか）。

