# Codeforces Raif Round div2 感想

Dがつらかったが、なんとかギリギリ修正して通せたので、レートは+45できた。

## A問題

+2するかどうかだけ気をつける。だいぶ簡単。

## B問題

これかなり難しいと思った。  
強連結成分分解できれば一瞬で溶けると思ったので、覚えたてながら使ってみてよかった。

ただし、模範解答に従ってupsolveは必要。

## C問題

スタックを使って `A` をできるだけ使うように貪欲に進める。

## D問題

構築だけど、これはどういう思考のもとで解けばいいのだろう。。？

自分が考えた方法は「左の列から決めていく」というもの。  
具体的には、 `2, 3` の設定方法について、以下のように進める（ `1, 0` についてはほぼ自明）。

- 列が `3` だったら、右にある `0` 以外のものとマッチングできるので、
できるだけ右にあるものを `3,2,1` の優先順位で選ぶ。
- 列が `2` のときは、右にある `1` のみをマッチングできるので、できるだけ右にあるものからマッチングさせていく。
- `i` 列目を観ているときは、基本的には `i` 行目を埋めていくことを考える。
  - まだ `i` 列目にターゲットが存在しないときは、 `(i, i)` にターゲットを設置する。すでに存在すれば何もしない。
  - マッチングした `tid` 列目については `(i, tid)` に2つめのターゲットを設置する。
  - 3つ目については `(tid, tid)` に設置する。

