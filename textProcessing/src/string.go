package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aaa, err := strconv.ParseBool("false")
	checkError(err)
	bbb, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	ccc, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	ddd, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	eee, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(aaa, bbb, ccc, ddd, eee)
	return

	// func Contains(s, substr string) bool
	//字符串s 中是否包含substr, 返回bool值
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true

	//func Join(a []string, sep string) string
	//字符串链接 把slice a 通过sep链接起来
	//s := []string{"foo", "bar", "baz"}
	//fmt.Println(strings.Join(s, ", "))
	// foo, bar, baz

	//func Index(s, sepstring) int
	// 在字符串s 中 查找 sep所在位置, 返回位置值, 找不到返回-1
	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) //-1

	//func Repeat(s string, count int) string
	//重复s 字符串 count 次, 最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))
	//banana

	//func Replace(s, old, new string, n int) string
	//在字符串中,把old 字符串替换为new字符串, n标识替换的次数, 小于0标识全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	//oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//moo moo moo

	//func Split(s, sep string) []string
	//把字符串按照sep分割,返回slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	//["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	//["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	//[" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo o'Higgins"))
	// [""]

	//func Trim(s string, cutset string) string
	// 在s字符串的头部和尾部去除cutset 指定的字符串
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
	// Achtung

	//func Fields(s string) []string
	//去除s字符串的空格符, 并且按照空格分割返回slice

	fmt.Printf("Fields are: %q", strings.Fields(" foo bar baz "))
	// Fields are: ["foo" "bar" "baz"]

	//字符串转换
	//字符串串化的函数在strconv中, 如下也是只是列出以下常用的
	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	//fmt.Println(string(str))

	// Format 系列函数把其他类型的转换为字符串
	aa := strconv.FormatBool(false)                //false
	bb := strconv.FormatFloat(123.23, 'g', 12, 64) //123.23
	cc := strconv.FormatInt(1234, 10)              //1234
	dd := strconv.FormatUint(12345, 10)            //123456
	ee := strconv.Itoa(1023)                       //1023
	fmt.Println(aa, bb, cc, dd, ee)

	//parse 系列函数把字符串  转为其他类型
	//aaa, err := strconv.ParseBool("false")
	//checkError(err)
	//bbb, err := strconv.ParseFloat("123.23", 64)
	//checkError(err)
	//ccc, err := strconv.ParseInt("1234", 10, 64)
	//checkError(err)
	//ddd, err := strconv.ParseUint("12345", 10, 64)
	//checkError(err)
	//eee, err := strconv.Atoi("1023")
	//checkError(err)
	//fmt.Println(aaa, bbb, ccc, ddd, eee)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
