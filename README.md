# ToDo-list-tui

![](./image/sample.jpg)

## 内容

- Go 言語を使用
- TUI ライブラリである [tview](https://github.com/rivo/tview) を使用しています
- 未完成でコードも汚いです

## 開発期間

2022 8/7 ~ 8/10

## 使い方

```bash
go mod tidy
```

```bash
go run main.go
```

## 注意

TUI の表示が崩れる場合は、次の方法で解決するかもしれません

- フォントの変更
- [tview text view miss display bug. #281](https://github.com/rivo/tview/issues/281)

```bash
export LC_CTYPE="en_US.UTF-8"
```

## 参考にした記事

- [APG4b](https://atcoder.jp/contests/APG4b)
- [A Tour of Go](https://go-tour-jp.appspot.com/list)
- [【Go】tview による TUI ツール作成](https://zenn.dev/minefuto/articles/cafc02dd63f65d)
- [Go 言語で TUI を作ってみたい](https://zenn.dev/foxtail88/scraps/1e98aca98266ec)
