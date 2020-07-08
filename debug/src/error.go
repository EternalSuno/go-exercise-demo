package src

import (
	"errors"
	"log"
	"os"
)

func dealError() {
	//Go 定义了一个叫做 error 的类型，来显式表达错误。在使用时，通过把返回的 error 变量与 nil 的比较，来判定操作是否成功。
	//例如 os.Open 函数在打开文件失败时将返回一个不为 nil 的 error 变量
	//
	f, err := os.Open("filename.txt")
	if err != nil {
		log.Fatal(err)
	}
	//类似于 os.Open 函数，标准包中所有可能出错的 API 都会返回一个 error 变量，以方便错误处理，
	//这个小节将详细地介绍 error 类型的设计，和讨论开发 Web 应用中如何更好地处理 error。
	//
	//Error 类型
	//error 类型是一个接口类型
	//type error interface {
	//	Error() string
	//}
	//error 是一个内置的接口类型，我们可以在 /builtin/ 包下面找到相应的定义。
	//而我们在很多内部包里面用到的 error 是 errors 包下面的实现的私有结构 errorString
	type errorString struct {
		s string
	}
	func (e *errorString) Error() string {
		return e.s
	}
	//你可以通过 errors.New 把一个字符串转化为 errorString，以得到一个满足接口 error 的对象，
	//其内部实现如下：
	func New(text string) error {
		return &errorString{text}
	}
	//下面这个例子演示了如何使用 errors.New:
	func Sqrt(f float64) (float64, error) {
		if f < 0 {
			return 0, errors.New("math: square root of negative number")
		}
	}



}

//func Open(name string) (file *File, err error)
