package main

/* 循环 */
import "fmt"

func main() {
	/* 循环跳出 */
	fmt.Println("--------- break ---------")
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

	fmt.Println("--------- continue ---------")
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

	fmt.Println("--------- goto ---------")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")

	/* 循环遍历 */
	fmt.Println("--------- for range遍历 ---------")
	var a = [...]string{"北京", "上海", "深圳"}
	for index, value := range a {
		fmt.Println(index, value)
	}
}
