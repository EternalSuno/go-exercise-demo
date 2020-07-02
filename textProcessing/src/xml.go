package src

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}

//XML 本质上是一种树形的数据格式，而我们可以定义与之匹配的 go 语言的 struct 类型，
//然后通过 xml.Unmarshal 来将 xml 中的数据解析成对应的 struct 对象。如上例子输出如下数据
//
//{{ servers} 1 [{{ server} Shanghai_VPN 127.0.0.1} {{ server} Beijing_VPN 127.0.0.2}]
//<server>
//<serverName>Shanghai_VPN</serverName>
//<serverIP>127.0.0.1</serverIP>
//</server>
//<server>
//<serverName>Beijing_VPN</serverName>
//<serverIP>127.0.0.2</serverIP>
//</server>
//}
//上面的例子中，将 xml 文件解析成对应的 struct 对象是通过 xml.Unmarshal 来完成的，这个过程是如何实现的？
//可以看到我们的 struct 定义后面多了一些类似于 xml:"serverName" 这样的内容，这个是 struct 的一个特性，
//它们被称为 struct tag，它们是用来辅助反射的。我们来看一下 Unmarshal 的定义：
//
//func Unmarshal(data []byte, v interface{}) error
//
//我们看到函数定义了两个参数，第一个是 XML 数据流，第二个是存储的对应类型，
//目前支持 struct、slice 和 string，XML 包内部采用了反射来进行数据的映射，
//所以 v 里面的字段必须是导出的。Unmarshal 解析的时候 XML 元素和字段怎么对应起来的呢？
//这是有一个优先级读取流程的，首先会读取 struct tag，如果没有，那么就会对应字段名。
//必须注意一点的是解析的时候 tag、字段名、XML 元素都是大小写敏感的，所以必须一一对应字段。
//
//Go 语言的反射机制，可以利用这些 tag 信息来将来自 XML 文件中的数据反射成对应的 struct 对象，
//关于反射如何利用 struct tag 的更多内容请参阅 reflect 中的相关内容。
//
//解析XML到struct的时候遵循如下的规则:
//如果 struct 的一个字段是 string 或者 [] byte 类型且它的 tag 含有 ",innerxml"，
//Unmarshal 将会将此字段所对应的元素内所有内嵌的原始 xml 累加到此字段上，
//如上面例子 Description 定义。最后的输出是
//
/*
	<server>
	<serverName>Shanghai_VPN</serverName>
	<serverIP>127.0.0.1</serverIP>
	</server>
	<server>
	<serverName>Beijing_VPN</serverName>
	<serverIP>127.0.0.2</serverIP>
	</server>
*/
//如果 struct 中有一个叫做 XMLName，且类型为 xml.Name 字段，
//那么在解析的时候就会保存这个 element 的名字到该字段，如上面例子中的 servers。
//如果某个 struct 字段的 tag 定义中含有 XML 结构中 element 的名称，
//那么解析的时候就会把相应的 element 值赋值给该字段，如上 servername 和 serverip 定义。
//如果某个 struct 字段的 tag 定义了中含有 ",attr"，
//那么解析的时候就会将该结构所对应的 element 的与字段同名的属性的值赋值给该字段，如上 version 定义。
//如果某个 struct 字段的 tag 定义 型如 "a>b>c",
//则解析的时候，会将 xml 结构 a 下面的 b 下面的 c 元素的值赋值给该字段。
//如果某个 struct 字段的 tag 定义了 "-",
//那么不会为该字段解析匹配任何 xml 数据。
//如果 struct 字段后面的 tag 定义了 ",any"，
//如果他的子元素在不满足其他的规则的时候就会匹配到这个字段。
//如果某个 XML 元素包含一条或者多条注释，那么这些注释将被累加到第一个 tag 含有 ",comments" 的字段上，
//这个字段的类型可能是 [] byte 或 string, 如果没有这样的字段存在，那么注释将会被抛弃。
//
//上面详细讲述了如何定义 struct 的 tag。 只要设置对了 tag，那么 XML 解析就如上面示例般简单，
//tag 和 XML 的 element 是一一对应的关系，如上所示，我们还可以通过 slice 来表示多个同级元素。

//注意： 为了正确解析，go 语言的 xml 包要求 struct 定义中的所有字段必须是可导出的（即首字母大写)
//
// 输出XML
// 假若我们不是要解析如上所示的 XML 文件，而是生成它，那么在 go 语言中又该如何实现呢？
//xml 包中提供了 Marshal 和 MarshalIndent 两个函数，来满足我们的需求。
//这两个函数主要的区别是第二个函数会增加前缀和缩进，函数的定义如下所示：
//
// func Marshal(v interface{}) ([]byte, error)
// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

//两个函数第一个参数都是用来生成XML的结构定义类型数据, 都是返回生成的XML数据流.

func test() {

	/*	<?xml version="1.0" encoding="utf-8"?>
		<servers version="1">
		<server>
		<serverName>Shanghai_VPN</serverName>
		<serverIP>127.0.0.1</serverIP>
		</server>
		<server>
		<serverName>Beijing_VPN</serverName>
		<serverIP>127.0.0.2</serverIP>
		</server>
		</servers>
		上面的 XML 文档描述了两个服务器的信息，包含了服务器名和服务器的 IP 信息，接下来的 Go 例子以此 XML 描述的信息进行操作。
	*/

	//解析xml
	//可以通过 xml 包的 Unmarshal 函数来达到我们的目的
	//func Unmarshal(data []byte, v interface{}) error
	//data 接收的是 XML 数据流，v 是需要输出的结构，定义为 interface，也就是可以把 XML 转换为任意的格式。
	//我们这里主要介绍 struct 的转换，因为 struct 和 XML 都有类似树结构的特征。

}
