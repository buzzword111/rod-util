# rodutil

go-rodのユーティリティ関数集です。  
よく使う処理やダイアログ操作など、go-rodをより便利に使うための関数をまとめています。

## インストール

```
go get github.com/buzzword111/rod-util
```

## 使い方

examples/ex1 ディレクトリ内のサンプルコードを参照してください。

## 関数

### func WaitAndHandleDialog

```go
func WaitAndHandleDialog(
	page *rod.Page,
	timeout time.Duration,
) *proto.PageJavascriptDialogOpening
```

Rodのページで発生するJavaScriptダイアログ（alert, confirm, prompt）を、指定したタイムアウト時間内で待機し、自動的に「OK」を押して閉じ、そのダイアログ情報を返します。タイムアウトした場合は `nil` を返します。