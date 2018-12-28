# go-sangoku
中国三国時代の魏・呉・蜀皇帝の名前を返すコマンドラインツール

## Install
```
$ go get github.com/jyouj/go-sangoku
```

## Usage
魏の皇帝の名前を知りたい時
```
$ go-sangoku -gi 1
曹丕/魏の文帝
```

蜀の皇帝の名前が知りたい時
```
$ go-sangoku -s 2
劉禅
```

呉の皇帝の名前が知りたい時
```
$ go-sangoku -go 3
孫休
```

それ以外
```
$ go-sangoku
漢王朝が滅亡し、群雄割拠の時代に移りました！
```
