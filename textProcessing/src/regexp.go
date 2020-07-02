package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//正则表达式是一种进行模式匹配和文本操纵的复杂而又强大的工具。虽然正则表达式比纯粹的文本匹配效率低，但是它却更灵活。
//按照它的语法规则，随需构造出的匹配模式就能够从原始文本中筛选出几乎任何你想要得到的字符组合。
//如果你在 Web 开发中需要从一些文本数据源中获取数据，那么你只需要按照它的语法规则，随需构造出正确的模式字符串就能够从原数据源提取出有意义的文本信息。
//
//Go 语言通过 regexp 标准包为正则表达式提供了官方支持，如果你已经使用过其他编程语言提供的正则相关功能，
//那么你应该对 Go 语言版本的不会太陌生，但是它们之间也有一些小的差异，因为 Go 实现的是 RE2 标准，
//除了 \C，详细的语法描述参考：http://code.google.com/p/re2/wiki/Syntax
//
//其实字符串处理我们可以使用 strings 包来进行搜索 (Contains、Index)、替换 (Replace) 和解析 (Split、Join) 等操作，
//但是这些都是简单的字符串操作，他们的搜索都是大小写敏感，而且固定的字符串，如果我们需要匹配可变的那种就没办法实现了，
//当然如果 strings 包能解决你的问题，那么就尽量使用它来解决。因为他们足够简单、而且性能和可读性都会比正则好。
//

//通过正则判断是否匹配
// regexp 包中含有三个函数用来判断是否匹配, 如果匹配返回true, 否则返回false
//func Match(pattern string, b []byte) (matched bool, error error)
//func MatchReader(pattern string, r io.RuneReader) (matched bool, error error)
//func MatchString(pattern string, s string) (matched bool, error error)

//上面的三个函数实现了同一个功能，就是判断 pattern 是否和输入源匹配，匹配的话就返回 true，
//如果解析正则出错则返回 error。三个函数的输入源分别是 byte slice、RuneReader 和 string。
//

//验证IP

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

//验证字符串
func regstring() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}
}

// 通过正则获取内容
//Match 模式只能用来对字符串的判断，而无法截取字符串的一部分、过滤字符串、或者提取出符合条件的一批字符串。
//如果想要满足这些需求，那就需要使用正则表达式的复杂模式。

func regexptest() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	//将html 标签全转为小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除 STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除 SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并进行换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))

	//从这个示例可以看出，使用复杂的正则首先是 Compile，它会解析正则表达式是否合法，
	//如果正确，那么就会返回一个 Regexp，然后就可以利用返回的 Regexp 在任意的字符串上面执行需要的操作。

	//解析正则表达式几个方法
	//func Compile(expr string) (*Regexp, error)
	//func CompilePOSIX(expr string) (*Regexp, error)
	//func MustCompile(str string) *Regexp
	//func MustCompilePOSIX(str string) *Regexp

	//CompilePOSIX 和 Compile 的不同点在于 POSIX 必须使用 POSIX 语法，它使用最左最长方式搜索，
	//而 Compile 是采用的则只采用最左方式搜索 (例如 [a-z]{2,4} 这样一个正则表达式，
	//应用于 "aa09aaa88aaaa" 这个文本串时，CompilePOSIX 返回了 aaaa，而 Compile 的返回的是 aa)。
	//前缀有 Must 的函数表示，在解析正则语法的时候，如果匹配模式串不满足正确的语法则直接 panic，而不加 Must 的则只是返回错误。
	//

	//
	//func (re *Regexp) Find(b []byte) []byte
	//func (re *Regexp) FindAll(b []byte, n int) [][]byte
	//func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
	//func (re *Regexp) FindAllString(s string, n int) []string
	//func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
	//func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
	//func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
	//func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
	//func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
	//func (re *Regexp) FindIndex(b []byte) (loc []int)
	//func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)
	//func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int
	//func (re *Regexp) FindString(s string) string
	//func (re *Regexp) FindStringIndex(s string) (loc []int)
	//func (re *Regexp) FindStringSubmatch(s string) []string
	//func (re *Regexp) FindStringSubmatchIndex(s string) []int
	//func (re *Regexp) FindSubmatch(b []byte) [][]byte
	//func (re *Regexp) FindSubmatchIndex(b []byte) []int
}

func exm() {
	a := "I am learning Go language"
	re, _ := regexp.Compile("[a-z]{2,4}")

	// 查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	//查找符合正则的所有 slice, n 小于 0 表示 返回全部符合的字符串， 不然就是返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	//查找符合条件的index位置， 开始位置和结束位置
	index := re.FindIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	//查找 Submatch, 返回数组，第一个元素是匹配的全部元素，第二个元素是第一个 () 里面的，第三个是第二个 () 里面的
	// 下面的输出第一个元素是 "am learning Go language"
	// 第二个元素是 " learning Go "，注意包含空格的输出
	// 第三个元素是 "uage"
	//

	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	//定义和上面的 FindIndex 一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	// FindAllSubmatch, 查找所有符合条件的子匹配

	submatchall := re2.FindAllStringSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex, 查找所有字匹配的 index

	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)

}
