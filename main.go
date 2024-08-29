package main

// GOARCH=amd64 GOOS=windows go build main.go && /home/mor/GOPATH/src/github.com/morentharia/py_sketches/venv/bin/python deploy.py

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	user32 := windows.NewLazyDLL("user32.dll")
	MessageBox := user32.NewProc("MessageBoxW")

	_ = MessageBox

	fmt.Println("fuck yeah")
	fmt.Println("fuck yeah")
	fmt.Println("fuck yeah")
	r, _, _ := MessageBox.Call(
		0,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("Hello fuck the World!"))),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("Example"))),
		windows.MB_OK,
	)

	fmt.Println("Return code:", r)
}
