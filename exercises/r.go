package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s1 []int         // nil slice
	s2 := []int{}        // empty, non-nil
	s3 := make([]int, 0) // empty, non-nil
	s4 := make([]int, 0, 5)

	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	h2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	h3 := (*reflect.SliceHeader)(unsafe.Pointer(&s3))
	h4 := (*reflect.SliceHeader)(unsafe.Pointer(&s4))

	fmt.Printf("s1  len=%d cap=%d nil=%v data=%#x\n", len(s1), cap(s1), s1 == nil, h1.Data)
	fmt.Printf("s2  len=%d cap=%d nil=%v data=%#x\n", len(s2), cap(s2), s2 == nil, h2.Data)
	fmt.Printf("s3  len=%d cap=%d nil=%v data=%#x\n", len(s3), cap(s3), s3 == nil, h3.Data)
	fmt.Printf("s4  len=%d cap=%d nil=%v data=%#x\n", len(s4), cap(s4), s4 == nil, h4.Data)

	// Mutations to show pointer changes
	s2 = append(s2, 1) // grow from empty
	fmt.Printf("s2' len=%d cap=%d nil=%v data=%#x\n", len(s2), cap(s2), s2 == nil, (*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	s4 = append(s4, 42) // within preallocated cap
	fmt.Printf("s4' len=%d cap=%d nil=%v data=%#x\n", len(s4), cap(s4), s4 == nil, (*reflect.SliceHeader)(unsafe.Pointer(&s4)).Data)
}
