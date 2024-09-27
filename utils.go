package main

import "unsafe"

func Ptr[T any](v T) uintptr {
	return (uintptr)(unsafe.Pointer(&v))
}
