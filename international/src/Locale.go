package src

func locale() {
	//Locale 是一组描述世界上某一特定区域文本格式和语言习惯的设置的集合。locale 名通常由三个部分组成：
	//第一部分，是一个强制性的，表示语言的缩写，例如 "en" 表示英文或 "zh" 表示中文。
	//第二部分，跟在一个下划线之后，是一个可选的国家说明符，用于区分讲同一种语言的不同国家，
	//例如 "en_US" 表示美国英语，而 "en_UK" 表示英国英语。最后一部分，跟在一个句点之后，是可选的字符集说明符，
	//例如 "zh_CN.gb2312" 表示中国使用 gb2312 字符集。
	//
	//GO 语言默认采用 "UTF-8" 编码集，所以我们实现 i18n 时不考虑第三部分，
	//接下来我们都采用 locale 描述的前面两部分来作为 i18n 标准的 locale 名。

	//在 Linux 和 Solaris 系统中可以通过 locale -a 命令列举所有支持的地区名，读者可以看到这些地区名的命名规范。
	//对于 BSD 等系统，没有 locale 命令，但是地区信息存储在 /usr/share/locale 中。

	//设置 Locale
	//有了上面对 locale 的定义，那么我们就需要根据用户的信息 (访问信息、个人信息、访问域名等) 来设置与之相关的 locale，
	//我们可以通过如下几种方式来设置用户的 locale。

	//通过域名设置 Locale
	//设置 Locale 的办法之一是在应用运行的时候采用域名分级的方式，
	//例如，我们采用 www.asta.com 当做我们的英文站 (默认站)，而把域名 www.asta.cn 当做中文站。
	//这样通过在应用里面设置域名和相应的 locale 的对应关系，就可以设置好地区。
	//

	//通过 URL 就可以很明显的识别
	//用户可以通过域名很直观的知道将访问那种语言的站点
	//在 Go 程序中实现非常的简单方便，通过一个 map 就可以实现
	//有利于搜索引擎抓取，能够提高站点的 SEO

	//if r.Host == "www.asta.com" {
	//    i18n.SetLocale("en")
	//} else if r.Host == "www.asta.cn" {
	//    i18n.SetLocale("zh-CN")
	//} else if r.Host == "www.asta.tw" {
	//    i18n.SetLocale("zh-TW")
	//}
	//
	//当然除了整域名设置地区之外，我们还可以通过子域名来设置地区，例如 "en.asta.com" 表示英文站点，"cn.asta.com" 表示中文站点。
	//
	//
	//prefix := strings.Split(r.Host,".")
	//
	//if prefix[0] == "en" {
	//	i18n.SetLocale("en")
	//} else if prefix[0] == "cn" {
	//	i18n.SetLocale("zh-CN")
	//} else if prefix[0] == "tw" {
	//	i18n.SetLocale("zh-TW")
	//}

	//从域名参数设置 Locale
	//目前最常用的设置 Locale 的方式是在 URL 里面带上参数，例如 www.asta.com/hello?locale=zh 或者 www.asta.com/zh/hello 。
	//这样我们就可以设置地区：i18n.SetLocale(params["locale"])。

	//这种设置方式几乎拥有前面讲的通过域名设置 Locale 的所有优点，它采用 RESTful 的方式，以使得我们不需要增加额外的方法来处理。
	//但是这种方式需要在每一个的 link 里面增加相应的参数 locale，这也许有点复杂而且有时候甚至相当的繁琐。
	//不过我们可以写一个通用的函数 url，让所有的 link 地址都通过这个函数来生成，
	//然后在这个函数里面增加 locale=params["locale"] 参数来缓解一下。

	//也许我们希望 URL 地址看上去更加的 RESTfu l 一点，例如：www.asta.com/en/books (英文站点) 和
	//www.asta.com/zh/books (中文站点)，这种方式的 URL 更加有利于 SEO，而且对于用户也比较友好，
	//能够通过 URL 直观的知道访问的站点。那么这样的 URL 地址可以通过 router 来获取 locale
	//(参考 REST 小节里面介绍的 router 插件实现)：

	//mux.Get("/:locale/books", listbook)

	//从客户端设置地区
	//在一些特殊的情况下，我们需要根据客户端的信息而不是通过 URL 来设置 Locale，这些信息可能来自于客户端设置的喜好语言 (浏览器中设置)，
	//用户的 IP 地址，用户在注册的时候填写的所在地信息等。这种方式比较适合 Web 为基础的应用。
	//

	//当然在实际应用中，可能需要更加严格的判断来进行设置地区
	//
	//IP 地址
	//
	//另一种根据客户端来设定地区就是用户访问的 IP，我们根据相应的 IP 库，对应访问的 IP 到地区，目前全球比较常用的就是 GeoIP Lite Country 这个库。这种设置地区的机制非常简单，我们只需要根据 IP 数据库查询用户的 IP 然后返回国家地区，根据返回的结果设置对应的地区。
	//
	//用户 profile
	//
	//当然你也可以让用户根据你提供的下拉菜单或者别的什么方式的设置相应的 locale，然后我们将用户输入的信息，保存到与它帐号相关的 profile 中，当用户再次登陆的时候把这个设置复写到 locale 设置中，这样就可以保证该用户每次访问都是基于自己先前设置的 locale 来获得页面。
	//

}
