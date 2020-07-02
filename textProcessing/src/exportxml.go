package src

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)
}

// 输出
/*
	<?xml version="1.0" encoding="UTF-8"?>
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
*/

// 和我们之前定义的文件的格式一模一样之所以会有 os.Stdout.Write([]byte(xml.Header)) 这句代码的出现，
//是因为 xml.MarshalIndent 或者 xml.Marshal 输出的信息都是不带 XML 头的，为了生成正确的 xml 文件，
//我们使用了 xml 包预定义的 Header 变量。
//

//我们看到 Marshal 函数接收的参数 v 是 interface {} 类型的，
//即它可以接受任意类型的参数，那么 xml 包，根据什么规则来生成相应的 XML 文件呢？

//如果 v 是 array 或者 slice，那么输出每一个元素，类似 value
//如果 v 是指针，那么会 Marshal 指针指向的内容，如果指针为空，什么都不输出
//如果 v 是 interface，那么就处理 interface 所包含的数据
//如果 v 是其他数据类型，就会输出这个数据类型所拥有的字段信息
//

//生成的 XML 文件中的 element 的名字又是根据什么决定的呢？元素名按照如下优先级从 struct 中获取：
//如果 v 是 struct，XMLName 的 tag 中定义的名称
//类型为 xml.Name 的名叫 XMLName 的字段的值
//通过 struct 中字段的 tag 来获取
//通过 struct 的字段名用来获取
//marshall 的类型名称
//

//我们应如何设置 struct 中字段的 tag 信息以控制最终 xml 文件的生成呢
//XMLName 不会被输出
//tag 中含有 "-" 的字段不会输出
//tag 中含有 "name,attr"，会以 name 作为属性名，字段值作为值输出为这个 XML 元素的属性，如上 version 字段所描述
//tag 中含有 ",attr"，会以这个 struct 的字段名作为属性名输出为 XML 元素的属性，类似上一条，只是这个 name 默认是字段名了。
//tag 中含有 ",chardata"，输出为 xml 的 character data 而非 element。
//tag 中含有 ",innerxml"，将会被原样输出，而不会进行常规的编码过程
//tag 中含有 ",comment"，将被当作 xml 注释来输出，而不会进行常规的编码过程，字段值中不能含有 "--" 字符串
//tag 中含有 "omitempty", 如果该字段的值为空值那么该字段就不会被输出到 XML，空值包括：false、0、nil 指针或 nil 接口，
//		任何长度为 0 的 array, slice, map 或者 string
//tag 中含有 "a>b>c"，那么就会循环输出三个元素 a 包含 b，b 包含 c，例如如下代码就会输出
//
/*
   FirstName string   `xml:"name>first"`
   LastName  string   `xml:"name>last"`

   <name>
   <first>Asta</first>
   <last>Xie</last>
   </name>
*/
