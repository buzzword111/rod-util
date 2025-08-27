package rodutil

import (
	"context"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// WaitAndHandleDialog は、Rodのページで発生するJavaScriptダイアログ（alert, confirm, prompt）を、指定したタイムアウト時間内で待機し、
// ダイアログが表示された場合は自動的に「OK」を押して閉じ、そのダイアログ情報（*proto.PageJavascriptDialogOpening）を返します。
// タイムアウトした場合は nil を返します。
// この関数は、ダイアログの内容を取得したい場合や、自動でダイアログを閉じたい場合に利用します。
func WaitAndHandleDialog(
	page *rod.Page,
	timeout time.Duration,
) *proto.PageJavascriptDialogOpening {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})
	var dialog *proto.PageJavascriptDialogOpening

	wait, handle := page.MustHandleDialog()
	go func() {
		dialog = wait()
		close(done)
	}()

	select {
	case <-done:
		handle(true, "") // ダイアログ(alert, confirm, prompt)に対して「OK」を押す
		return dialog
	case <-ctx.Done():
		return nil
	}
}
