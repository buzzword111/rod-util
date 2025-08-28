# rodutil

go-rodのユーティリティ関数集です。  
よく使う処理やダイアログ操作など、go-rodをより便利に使うための関数をまとめています。

## インストール

```
go get github.com/buzzword111/rod-util
```

## 使い方

### サンプル

`examples/main.go`:

```go
package main

import (
	"fmt"
	"time"
	"path/filepath"

	rodutil "github.com/buzzword111/rod-util"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Chrome を起動
	url := launcher.New()
	url.Headless(false) // ヘッドレスをオフ
	launchURL := url.MustLaunch()
	browser := rod.New().ControlURL(launchURL)
	browser.Trace(true)
	browser.SlowMotion(1000)
	browser.MustConnect()
	defer browser.MustClose()

	// 相対パスでsample.htmlを開く
	relativePath := "./sample.html"
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}
	page := browser.MustPage("file://" + absPath)

	// goroutineでクリック実行（ダイアログ発生のトリガー）
	page.MustElement("#input").MustInput("123")
	page.MustElement("#btn_valid").MustClick()

	dialog := rodutil.WaitAndHandleDialog(page, 5*time.Second)
	if dialog != nil {
		fmt.Println("Dialog type:", dialog.Type) // → alert
		fmt.Println("Dialog message:", dialog.Message)
		return
	}
	fmt.Println("alertダイアログが表示されませんでした（タイムアウト）")
}
```

`examples/sample.html`:

```html
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Alert Sample</title>
</head>
<body>
  <button id="btn" onclick="alert('Hello!')">Click me</button>
  <input type="text" id="input" />
  <button id="btn_valid" onclick="document.getElementById('input').value.length > 2 ? alert('Hello!') : null;">Click me</button>
</body>
</html>
```

## 関数

### func WaitAndHandleDialog

```go
func WaitAndHandleDialog(
	page *rod.Page,
	timeout time.Duration,
) *proto.PageJavascriptDialogOpening
```

Rodのページで発生するJavaScriptダイアログ（alert, confirm, prompt）を、指定したタイムアウト時間内で待機し、自動的に「OK」を押して閉じ、そのダイアログ情報を返します。タイムアウトした場合は `nil` を返します。