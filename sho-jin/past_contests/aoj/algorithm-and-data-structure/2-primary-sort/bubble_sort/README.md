# バブルソート

- 挿入ソートと同じように、前半がソート済み配列となるような形になるが、ソート済みの部分が必ずグローバルに見て完成した形になる。
  - 泡のように、小さい値が一番前にくるように要素間での交換がなされるため。
- `A[j] < A[j-1]` とすれば、要素のキーが同じ場合にそれらの順番が入れ替わらないため安定ソートだが、等号を含めてしまうと安定ソートではなくなってしまう。
- バブルソートの交換回数は反転数または転倒数と呼ばれるもので、列の乱れの具合を表す数値として知られている。

