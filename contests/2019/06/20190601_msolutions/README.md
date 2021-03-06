# M-SOLUTIONS 2019 感想

簡単すぎる2問の完答でパフォ1130ぐらいの微下がり。
Dの500が考え直したらあまりにも単純だったため、正しく反省したい。

- A問題は三角形の内角の和を計算させるもの。
- B問題はちょっとかんがえてif文を書くだけ。
- C問題は期待値に関する問題。答え方の形式がみかけないものだったので避けた。
- D問題は木を題材とした構築問題。
  - 次数の小さいところ、もしくは大きいところから埋めていく、と考えたが、埋まっていくところから実質次数が減っていくため、これをうまく実現する方法がなかった。
    - あるのかもしれないがもっと簡単に考えて良い。
  - **実は根から大きい数を幅優先で割り当てていけば良いだけ。**
  - ここに至るまでの考察すべき内容を考えてみる。
    - 最大値はどう頑張っても和に組み込めない（枝に書き込めない）
      - これはまっさきに気づいた。
    - **ありえる和の最大値は `Sum(C...) - Max(C...)` である。これ以上は大きくできないのは自明。**
      - これはけんちょんさんがツイートしていた内容。
      - コンテスト中は和がどうなるか想像できていなかったので、ここまでちゃんと考察できていれば、違う思考ができたかもしれない。
      - idsigmaさんも「上界・下界」といったワードをつぶやいており、確かに限界を把握しておくのは考察を進める上で大事だと思った。
    - **そもそもDFSに比べてBFSの経験値が低く、順に埋めていく方法のイメージの幅が狭かった気がする。**
      - 深さ優先は頭によぎったが、どう考えてもいい埋め方にはならないと棄却してしまった。
      - BFSも今回コンテスト後に適当に書いてみたら、むしろDFSよりも簡単に書ける気がするので、グリッドだけでなく、グラフ・木でもちゃんと念頭に置いて置けるようにすること！
  - ところで、PDFの想定解法はなんか難しい（`O(N^2)` 解法が要求されている。）
    - ソート＋BFSはもちろんソート部分がネックになっているので `O(NlogN)` で済む。
    - 今回の問題は制約も惑わす要素に鳴っている気がする。
  - 解説放送を見るとDFSでも行けるらしい。。
    - 「大きいものが連結していること」が最大のポイントらしいが、あまりしっくりは来ない。。
    - `ci` が大きいものを連結成分にギュッとまとめる、というのはわかりやすい気がする。
  - **典型: 上限・下限等の限界値を考え、それを実現することを目指す。**
  - というか最もダメだったのは、**1つの手法に固執し続けてしまったこと。**
    - 詰まったらもうその手法はダメだと諦めて思考をリセットすることを考えよう。

---

## 順位表の使い方

ちょっと前から順位表を見ると集中力が散漫になるのでやめていたが、
順位表を見ることは毒にも薬にもなると考えられる気がするので、ちょっと整理してみる。

- 特に迷うまでは順位表は見ない。
  - 意味がない。順位表を見て提出を決めるかどうかはナンセンスだと思う。
- 迷って迷って考えられることがなかったら見て良い。
  - この状態で初めて順位表からメタ読みができるようになる気がする。少なくとも自分は。

## Cについて（2019-08-21に復習）

- 確率 `p` で成功する試行について、1回成功するまでに要する試行回数の期待値は `1/p` というのが最重要前提知識。
- この他に、引き分けでないことがわかった上での条件付き確率、それを利用した条件付き期待値の整理が求められる。
  - 慣れないとまだまだ難しそう。。

### 参考資料

- わかりやすいと思ったのは、[kmjpさんのブログ解説記事](http://kmjp.hatenablog.jp/entry/2019/06/04/1000)
- [mathさんの期待値と条件付き確率の記事](https://math314.hateblo.jp/entry/2013/12/12/232045)も一般的な話の勉強にとても役立った。

#### Twitterでの質問の下書き

---

確率・期待値が苦手すぎなので、競技プログラミングQ&Aサイト「Twitter」に投げてみようと思う。

---

M-SolutionsプロコンC問題の解法

X(M): 「勝敗のつくゲーム」がちょうどM回行われる確率
Y(M): 「勝敗のつくゲーム」がちょうどM回行われる、、までに引き分けを含めて「ゲームが行われる回数の期待値」

\sum_X(M)*Y(M) が答え。

（続く）

---

Y(M)を「勝敗のつくゲームがM回行われるときの条件付き期待値」と考えると納得できる。
Y(M)の式もM回表が出るまでのコイントスの回数の期待値をイメージすれば直感的に納得できる。

（続く）

---

X(M)の計算がわかるようでやっぱりわからない。
p=A/(A+B)と(1-p)=B/(A+B)で反復試行の確率で考えているように見えるけど、pってどういう確率と解釈したらいいんですか？

##### 着地点

- やはり当初の考えどおり、上記の `p` は「勝敗がつくとわかっている」条件のもとでの、Aが勝つ条件付き確率、という認識で良い。
- なぜ引き分けの可能性を考えないかというと、 `Y(M)` の中にM回勝敗が決るゲームが行われる間での、 **引き分けが起こる回数も考慮されきっているから。**
  - むしろ `X(M)` の中では引き分けの可能性を考えてはいけない、ということになる。
- 厳密に示すにはただただ計算するしかない、とのこと。
  - そのために良い記事がこちらの[ふるやんさんの解説記事](https://www.creativ.xyz/best-of-2n-1-1080/)。

