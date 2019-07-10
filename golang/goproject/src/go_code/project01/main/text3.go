package main

import "fmt"

// 本函数测试入口参数和返回值情况
func dummy(b int) int {

    // 声明一个c赋值进入参数并返回
    var c int
    c = b

    return c
}

// 空函数, 什么也不做
func void() {

}

func main() {

    // 声明a变量并打印
    var a int

    // 调用void()函数
    void()

    // 打印a变量的值和dummy()函数返回
    fmt.Println(a, dummy(0))
}
/*
第 6 行，dummy() 函数拥有一个参数，返回一个整型值，测试函数参数和返回值分析情况。
第 9 行，声明 c 变量，这里演示函数临时变量通过函数返回值返回后的情况。
第 16 行，这是一个空函数，测试没有任何参数函数的分析情况。
第 23 行，在 main() 中声明 a 变量，测试 main() 中变量的分析情况。
第 26 行，调用 void() 函数，没有返回值，测试 void() 调用后的分析情况。
第 29 行，打印 a 和 dummy(0) 的返回值，测试函数返回值没有变量接收时的分析情况。
*/