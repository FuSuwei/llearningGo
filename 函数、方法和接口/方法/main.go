package main

import (
	"fmt"
	"os"
)

// 先打开文件对象
func main() {
	f, _ := os.Open("D:\\go-code\\test\\learnGo\\函数、方法和接口\\方法\\foo.dat")

	// 绑定到了 f 对象
	// func Close() error
	var Close = func() error {
		return (*os.File).Close(f)
	}
	// 绑定到了 f 对象
	// func Read(offset int64, data []byte) int
	var Read = func(data []byte)  (n int, err error){
		return (*os.File).Read(f, data)
	}
	bs := make([]byte,4, 4)
	n, err := Read(bs)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("----n", n)
	err = Close()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}