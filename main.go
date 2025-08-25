package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Chrome を起動
	url := launcher.New()
	url.Headless(false) // 👈 ヘッドレスをオフ
	launchURL := url.MustLaunch()
	browser := rod.New().ControlURL(launchURL)
	browser.Trace(true)
	browser.SlowMotion(1000)
	browser.MustConnect()
	defer browser.MustClose()

	// ローカルのHTMLを開く（パスを環境に合わせて変更）
	page := browser.MustPage("file:///Users/buzzword111/Programs/Go/20250825_rod_alert_sample/sample.html")

	// ダイアログのハンドラをセット
	wait, handle := page.MustHandleDialog()

	// goroutineでクリック実行（ダイアログ発生のトリガー）
	page.MustElement("#input").MustInput("1")
	go page.MustElement("#btn_valid").MustClick()

	// ダイアログが出るのを待機
	dialog := wait()

	fmt.Println("Dialog type:", dialog.Type)       // → alert
	fmt.Println("Dialog message:", dialog.Message) // → Hello!

	// OKを押す（dismiss する場合は false）
	handle(true, "")
}
