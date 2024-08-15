// マルチプレクサ Multiplexer
// リクエストをURL毎、対応ハンドラへの転送を行う。
package main

import "net/http"

func NewMux() http.Handler {
	mux := http.NewServeMux()
	// 状態確認用
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 静的解析のエラーを回避するため、明示的に戻り値を捨てる
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	return mux
}
