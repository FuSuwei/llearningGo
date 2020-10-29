package main

//type UpperWriter struct {
//	io.Writer
//}
//
//func (p *UpperWriter) Write(data []byte) (n int, err error) {
//	return p.Writer.Write(bytes.ToUpper(data))
//}
//
//func main() {
//	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
//}


// 自定义输出
//type UpperString string
//
//func (s UpperString) String() string {
//	return strings.ToUpper(string(s))
//}
////
////type fmt.Stringer interface {
////String() string
////}
//
//func main() {
//	fmt.Fprintln(os.Stdout, UpperString("hello, world大地"))
//}

// 伪造内部方法
import (
	"fmt"
	"testing"
)

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func main() {
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
}