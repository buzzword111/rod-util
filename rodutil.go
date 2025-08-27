package rodutil

import (
	"context"
	"time"

	"github.com/go-rod/rod/lib/proto"
)

// WaitAndHandleDialog は、Rodのダイアログイベントを指定したタイムアウト時間内で待機し、
// ダイアログが表示された場合は handle 関数で処理し、そのダイアログ情報を返します。
// タイムアウトした場合は nil を返します。
func WaitAndHandleDialog(
	wait func() *proto.PageJavascriptDialogOpening,
	handle func(bool, string),
	timeout time.Duration,
) *proto.PageJavascriptDialogOpening {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})
	var dialog *proto.PageJavascriptDialogOpening

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
