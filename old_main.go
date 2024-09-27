package main

// export GOARCH=amd64 GOOS=windows
// export GO111MODULE=on
// GOARCH=amd64 GOOS=windows go build main.go && /home/mor/GOPATH/src/github.com/morentharia/py_sketches/venv/bin/python deploy.py

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/fatih/color"
	"github.com/k0kubun/pp"
	peparser "github.com/saferwall/pe"
	w "golang.org/x/sys/windows"
)

// https://github.com/Dabudabot/otus-re/blob/master/14/PELoader/PELoader.c
func old_main() {
	filename := "C:\\fasmw\\EXAMPLES\\HELLO\\HELLO.exe"
	pe, err := peparser.New(filename, &peparser.Options{})
	if err != nil {
		log.Fatalf("Error while opening file: %s, reason: %v", filename, err)
	}

	// Print with default helper functions
	color.Cyan("Prints text in cyan.")
	color.Cyan("Prints text in cyan.")
	color.Cyan("Prints text in cyan.")
	fmt.Println("fuck yeah")
	fmt.Println("fuck yeah")
	err = pe.Parse()
	if err != nil {
		log.Fatalf("Error while parsing file: %s, reason: %v", filename, err)
	}
	log.Printf("filename = %#v\n", filename)
	//---------------------------------
	// pp.Print(pe)

	// pe.RichHeader
	fmt.Printf("Magic is: 0x%x\n", pe.DOSHeader.Magic)
	fmt.Printf("Signature is: 0x%x\n", pe.NtHeader.Signature)
	// fmt.Printf("Machine is: 0x%x, Meaning: %s\n", pe.NtHeader.FileHeader.Machine, pe.ImageNtHeader.FileHeader.)
	pp.Print(pe.RichHeader)

	for _, sec := range pe.Sections {
		// pp.Print(sec)
		fmt.Printf("Section Name : %s\n", sec.Header.Name)
		fmt.Printf("Section VirtualSize : %x\n", sec.Header.VirtualSize)
		fmt.Printf("Section Flags : %x, Meaning: %v\n\n",
			sec.Header.Characteristics, sec.PrettySectionFlags())
	}
	return

	user32 := w.NewLazyDLL("user32.dll")
	MessageBox := user32.NewProc("MessageBoxW")

	// w.UnmapViewOfFile()
	_ = MessageBox

	fmt.Println("fuck yeah")
	fmt.Println("fuck yeah")
	fmt.Println("fuck yeah")
	r, _, _ := MessageBox.Call(
		0,
		uintptr(unsafe.Pointer(w.StringToUTF16Ptr("Hello fuck the World!"))),
		uintptr(unsafe.Pointer(w.StringToUTF16Ptr("Example"))),
		w.MB_OK,
	)

	fmt.Println("Return code:", r)
}
