# AGC016過去問感想

Last Change: 2020-02-16 00:36:13.

- A問題は文字列処理の圧縮に関する問題。22分ほどかかってしまった。
  - ある文字を選んで、それでスプリットしたときに得られる各スライスの長さのうち、最大のものがその文字で塗りつぶすための操作の最大となる。
    - ややオーバーキルな方法だったかもしれない。
  - シミュレーションでよい。
    - [解説放送](https://www.youtube.com/watch?v=kdQtQSgIYPI)をみるとシミュレーションのやり方はわかるようになる。
      - 本番でシミュレーションのやり方がわからなかったので反省。
      - **文字列圧縮のルールを、隣接するもののうちどちらか1つを選ぶ、と読み替えられるとずいぶん簡単になる。**
    - 文字列に含まれるものだけではなくて、26文字すべてについて試すのが良さそう。
      - その文字1種類だけになったか？もいちいち愚直にチェックすれば良さそう。
    -  `26*100*100` なので余裕そう。シミュレーションをやるならバグを減らすためにとにかく愚直に。

## B問題（@2020-02-16）

時間はかかったが、なんとか自力で導けた。
PDFを観た感じ少し着眼点が違う気がしたので、解説放送でも学んでいく。

**とにかく具体的な例を弄ることの大切さを学んだ気がした。**
**考え始めは考えつきやすい極端な例がいいと思う。**

- aloneなものに着目する、というのが主眼っぽい
  - 。。難しくて一見しただけでは理解しきれなかった。。

### 自分のやり方

初期状態を「すべての帽子が異なる色として考えて、見えている帽子の色のmaxを基準にして、他の帽子を既存の色で塗りつぶしていく」
という方法を考えた。

。。色々整理していくと、解説でいうaloneな帽子が作れる個数の範囲が求まるので、
それと照らし合わせて範囲内に収まっているのならYes、そうでないならNoと考えた。

