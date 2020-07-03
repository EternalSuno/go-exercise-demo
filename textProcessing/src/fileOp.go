package src

import (
	"fmt"
	"os"
)

//目录操作
//文件操作的大多数函数都是在 os 包里面，下面列举了几个目录操作的：
//
//func Mkdir(name string, perm FileMode) error
//
//创建名称为 name 的目录，权限设置是 perm，例如 0777
//
//func MkdirAll(path string, perm FileMode) error
//
//根据 path 创建多级子目录，例如 astaxie/test1/test2。
//
//func Remove(name string) error
//
//删除名称为 name 的目录，当目录下有文件或者其他目录时会出错
//
//func RemoveAll(path string) error
//
//根据 path 删除多级子目录，如果 path 是单个名称，那么该目录下的子目录全部删除。
//

func filetest() {
	os.Mkdir("astaxie", 0777)
	os.Mkdir("astaxie/test1/test2", 0777)
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("astaxie")
}

//文件操作
//建立与打开文件
//新建文件可以通过如下两个方法
//
//func Create(name string) (file *File, err Error)
//
//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是 0666 的文件，返回的文件对象是可读写的。
//
//func NewFile(fd uintptr, name string) *File
//
//根据文件描述符创建相应的文件，返回一个文件对象
//

//通过如下两个方法来打开文件：
//
//func Open(name string) (file *File, err Error)
//
//该方法打开一个名称为 name 的文件，但是是只读方式，内部实现其实调用了 OpenFile。
//
//func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
//
//打开名称为 name 的文件，flag 是打开的方式，只读、读写等，perm 是权限
//

//写文件
//写文件函数：
//
//func (file *File) Write(b []byte) (n int, err Error)
//
//写入 byte 类型的信息到文件
//
//func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
//
//在指定位置开始写入 byte 类型的信息
//
//func (file *File) WriteString(s string) (ret int, err Error)
//
//写入 string 信息到文件
//

func filetestt() {
	userFile := "astaxie.txt"
	fout, err := os.Create(userfile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}

//读文件
//func (file *File) Read(b []byte) (n int, err Error)
//读取数据到 b 中
//func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
//从 off 开始读取数据到 b 中

func readfile() {
	userFile := "asatxie.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

//删除
//func Remove(name string) Error
