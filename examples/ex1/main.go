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
	// url.Headless(false)
	launchURL := url.MustLaunch()
	browser := rod.New().ControlURL(launchURL)
	// browser.Trace(true)
	// browser.SlowMotion(1 * time.Second)
	browser.MustConnect()
	defer browser.MustClose()

	relativePath := "./sample.html"
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}
	page := browser.MustPage("file://" + absPath)

	page.MustElement("#input").MustInput("123")
	go page.MustElement("#btn_valid").MustClick()

	dialog := rodutil.WaitAndHandleDialog(page, 5*time.Second)
	if dialog != nil {
		fmt.Println("Dialog type:", dialog.Type)
		fmt.Println("Dialog message:", dialog.Message)
		return
	}
	fmt.Println("alertダイアログが表示されませんでした（タイムアウト）")
}
