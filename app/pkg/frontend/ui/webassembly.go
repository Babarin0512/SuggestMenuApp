package main
 
import (
	"syscall/js"
)
 
func main() {
	// グローバルオブジェクト(window)を取得する
	window := js.Global()
 
	// document オブジェクトを取得する
	document := window.Get("document")
 
	// bodyを取得する
	body := document.Get("body")
 
	// h1 のDOMを作成する
	h1 := document.Call("createElement", "h1")
	h1.Set("innerHTML", "Hello Webassembly!")
 
	// h1をbodyに追加する
	body.Call("appendChild", h1)
 
	// プログラムが終了しないように待機する
	select {}
 
}