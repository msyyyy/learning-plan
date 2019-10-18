package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	dogP := &dog
	/*
	一个指针值（比如*Dog类型的值）可以被转换为一个unsafe.Pointer类型的值，反之亦然。
	一个uintptr类型的值也可以被转换为一个unsafe.Pointer类型的值，反之亦然。
	一个指针值无法被直接转换成一个uintptr类型的值，反过来也是如此。
	*/
	dogPtr := uintptr(unsafe.Pointer(dogP))

	// nsafe.Offsetof函数用于获取两个值在内存中的起始存储地址之间的偏移量，以字节为单位。
	// 这两个值一个是某个字段的值，另一个是该字段值所属的那个结构体值。我们在调用这个函数的时候，需要把针对字段的选择表达式传给它，比如dogP.name。
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Printf("nameP == &(dogP.name)? %v\n",
		nameP == &(dogP.name))
	fmt.Printf("The name of dog is %q.\n", *nameP)

	*nameP = "monster"
	fmt.Printf("The name of dog is %q.\n", dogP.name)
	fmt.Println()

	// 示例2。
	// 下面这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期。
	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)

}