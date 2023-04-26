package main

import "fmt"

/* 切片 */
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high) --左包含，右不包含

	// len() 函数求长度
	// cap() 函数求切片的容量
	/**
	 * cap()函数是从 s 的第一个元素到底层数组（a）末尾的元素数量
	 * 在这个例子中，s 从下标 1 开始切 a，到 a 数组最后一个元素（下标 4）结束，共有 4 个元素
	 */
	fmt.Printf("a:%v | len(a):%v | cap(a):%v\n", a, len(a), cap(a)) // a:[1 2 3 4 5] | len(a):5 | cap(a):5
	fmt.Printf("s:%v | len(s):%v | cap(s):%v\n", s, len(s), cap(s)) // s:[2 3] | len(s):2 | cap(s):4
}
