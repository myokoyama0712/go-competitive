# 二分探索

Last Change: 2020-05-05 20:58:34.

めぐる式二分探索信者なので、基本的にはそれに愚直に従って解く。

## 苦手 or 重要問題集

- [ABC149 E.Handshake](https://atcoder.jp/contests/abc149/tasks/abc149_e)
  - 考え方も重要で、めぐる式だと結構実装でこんがらがる。
  - 境界値周りの調整方法については、自分にとって一見非自明であるため、非常に重要。
  - [maspyさんの解説](https://maspypy.com/atcoder-参加感想-2019-12-29abc-149)が非常にわかりやすい。
    - 「最適解の構造の観察」の部分はこの問題に限らずに幅広い応用が可能なため、超重要。
- [ARC037 C.億マス計算](https://atcoder.jp/contests/arc037/tasks/arc037_c)
  - 典型かつ超重要問題、そして苦手。
  - 復習してみたらABC149Eのほうが難しかった。
- [Indeedなう（予選A） C.説明会](https://atcoder.jp/contests/indeednow-quala/tasks/indeednow_2015_quala_3)
  - logが2つの形でできる。
  - なんか解説スライドはシンプルな方法を取っていたが正当性がよくわからなかった。
- [CODE FESTIVAL 2015 予選A D.壊れた電車](https://atcoder.jp/contests/code-festival-2015-quala/tasks/codefestival_2015_qualA_d)
  - 普通にやるとlogが2つ出てくるので、できれば工夫したい。
- [ABC032 D.ナップサック問題](https://atcoder.jp/contests/abc032/tasks/abc032_d)
  - 半分全列挙で二分探索を行う。
  - 単調性を持つようにデータを調整するところで貪欲的な考え方をする必要があり、個人的に苦手なので復習する価値が大きい。
- [DISCO2016本選 B.DDPC特別ビュッフェ2](https://atcoder.jp/contests/discovery2016-final/tasks/discovery_2016_final_b)
  - ちょっと特殊だがlogが2個つくタイプの二分探索。
- [ABC123 D.Cake 123](https://atcoder.jp/contests/abc123/tasks/abc123_d)
  - 基本的ではあるが、これも同率K番目のものの扱い方を適当にやるとWAになってしまう。

## 決め打ち二分探索

- [ABC063 D.Widespread](https://atcoder.jp/contests/abc063/tasks/arc075_b)
  - 割と初歩的ではあるが、オーバーフローの罠があるので注意が必要。

## 実装上の注意点

以下の3点に集約されると思う。
逆に、これらを肝に銘じておけば、正しい実装になるはず。

1. 最初に設定された `ng, ok` は絶対に `mid` には入らない。
2. `ok` は一回も動かない場合がある。
3. 最終的に `ok` が参照すべきものとなる。

