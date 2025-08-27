package main

import (
	"fmt"
	"path/filepath"
	"time"

	rodutil "github.com/buzzword111/rod-util"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Chrome を起動
	url := launcher.New()
	url.Headless(false)
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

	// ダイアログのハンドラをセット
	wait, handle := page.MustHandleDialog()

	// goroutineでクリック実行（ダイアログ発生のトリガー）
	page.MustElement("#input").MustInput("1")
	go page.MustElement("#btn_valid").MustClick()

	// --- ここが呼び出し例 ---
	dialog := rodutil.WaitAndHandleDialog(wait, handle, 5*time.Second)
	if dialog != nil {
		fmt.Println("Dialog type:", dialog.Type)
		fmt.Println("Dialog message:", dialog.Message)
	} else {
		fmt.Println("alertダイアログが表示されませんでした（タイムアウト）")
	}
}
