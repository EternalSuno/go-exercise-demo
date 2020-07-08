package src

import (
	"fmt"
	"os"
	"path"
	"time"
)

func web() {
	//国际化站点
	// 管理多个本地包
	//开发一个应用的时候，首先我们要决定是只支持一种语言，还是多种语言，如果要支持多种语言，我们则需要制定一个组织结构，
	//以方便将来更多语言的添加。在此我们设计如下：Locale 有关的文件放置在 config/locales 下，假设你要支持中文和英文，
	//那么你需要在这个文件夹下放置 en.json 和 zh.json。大概的内容如下所示：
	//
	//# zh.json
	//{
	//	"zh": {
	//		"submit": "提交",
	//		"create": "创建"
	//	}
	//}
	//	en.json
	//{
	//	"en": {
	//		"submit": "Submit",
	//		"create": "Create"
	//	}
	//}

	//为了支持国际化，在此我们使用了一个国际化相关的包 —— go-i18n，
	//首先我们向 go-i18n 包注册 config/locales 这个目录，以加载所有的 locale 文件

	//fmt.Println(Tr.Translate("submit"))
	////输出Submit
	//Tr.SetLocale("zh")
	//fmt.Println(Tr.Translate("submit"))
	////输出“提交”

	// 自动加载本地包
	//了如何自动加载自定义语言包，其实 go-i18n 库已经预加载了很多默认的格式信息，
	//例如时间格式、货币格式，用户可以在自定义配置时改写这些默认配置，请看下面的处理过程：

	// locale=zh 的情况下，执行如下代码：

	fmt.Println(Tr.Time(time.Now()))
	// 输出：2009 年 1 月 08 日 星期四 20:37:58 CST

	fmt.Println(Tr.Time(time.Now(), "long"))
	// 输出：2009 年 1 月 08 日

	fmt.Println(Tr.Money(11.11))
	// 输出: ￥11.11

}

//// 加载默认配置文件，这些文件都放在 go-i18n/locales 下面
//// 文件命名 zh.json、en.json、en-US.json 等，可以不断的扩展支持更多的语言
func (il *IL) loadDefaultTranslations(dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		fullPath := path.Join(dirPath, name)

		fi, err := os.Stat(fullPath)
		if err != nil {
			return err
		}

		if fi.IsDir() {
			if err := il.loadTranslations(fullPath); err != nil {
				return err
			}
		} else if locale := il.matchingLocaleFromFileName(name); locale != "" {
			file, err := os.Open(fullPath)
			if err != nil {
				return err
			}
			defer file.Close()
			if err := il.loadTranslation(file, locale); err != nil {
				return err
			}
		}
	}
	return nil
}
