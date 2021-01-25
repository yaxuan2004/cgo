package main

import (
	"C"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

var c chan int

func decrement(n int) {
	for n > 0 {
		n -= 1
	}
}

//export mycal
func mycal() int {
	var t int = 0
	for i := 0; i < 10000000; i++ {
		t += i
	}
	return t
}

//export qrcodeText
func qrcodeText() int {
	qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 256, "qr.png")
	return 0
}

//export count_time
func count_time() *C.char {
	start := time.Now()
	decrement(100000000)
	total_time := time.Since(start).String()
	return C.CString(total_time)
}

func main() {}
