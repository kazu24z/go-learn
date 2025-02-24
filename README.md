## このリポジトリについて

このリポジトリは、Go言語の学習を目的として作成されています。


## 使い方

1. answersディレクトリにsample, sample2のように問題の解答ごとにディレクトリを切り、そのなかに.goファイルを作成します。
2. registry.goに例を参考に、問題ごとの解答をimportします。
3. 以下のコマンドで実行します。

```bash
go run main.go <ディレクトリ名>
```

## エイリアスの登録
answersの下にサブディレクトリを作成できるようにしていますが、都度フルパス入力は面倒です。
なので、エイリアス登録してコマンド実行できます。

1. main.go > Prefixesにパスとそのエイリアスを登録

```
var Prefixes = map[string]string{
    "t-": "tutorial/",
}
```

2. 実行

```
go run main.go t-no1
```
