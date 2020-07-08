package src

import (
	"fmt"
	"time"
)

var locales map[string]map[string]string

func test() {
	//本地化文本消息
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	locales["en"] = en
	cn := make(map[string]string, 10)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	locales["zh-CN"] = cn
	lang := "zh-CN"
	fmt.Println(msg(lang, "pea"))
	fmt.Println(msg(lang, "bean"))

	//本地化日期和时间
	//因为时区的关系，同一时刻，在不同的地区，表示是不一样的，而且因为 Locale 的关系，时间格式也不尽相同，
	//例如中文环境下可能显示：2012年10月24日 星期三 23时11分13秒 CST，而在英文环境下可能显示:
	//Wed Oct 24 23:11:13 CST 2012。这里面我们需要解决两点:
	//时区问题
	//格式问题
	//$GOROOT/lib/time 包中的 timeinfo.zip 含有 locale 对应的时区的定义，为了获得对应于当前 locale 的时间，
	//我们应首先使用 time.LoadLocation(name string) 获取相应于地区的 locale，比如 Asia/Shanghai 或 America/Chicago
	//对应的时区信息，然后再利用此信息与调用 time.Now 获得的 Time 对象协作来获得最终的时间。
	//详细的请看下面的例子 (该例子采用上面例子的一些变量):
	//
	en["time_zone"] = "America/Chicago"
	cn["time_zone"] = "Asia/Shanghai"

	loc, _ := time.LoadLocation(msg(lang, "time_zone"))
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))

	en["date_format"] = "%Y-%m-%d %H:%M:%S"
	cn["date_format"] = "%Y年%m月%d日 %H时%M分%S秒"
	fmt.Println(date(msg(lang, "date_format"), t))

	//本地化货币值
	//各地货币表示也不一样, 处理方式与日期差不多
	en["money"] = "USD %d"
	cn["money"] = "￥%d元"

	fmt.Println(money_format(msg(lang, "date_format"), 100))

}
func money_format(formate string, money int64) string {
	return fmt.Sprintf(formate, money)
}
func date(fomate string, t time.Time) string {
	year, month, day = t.Date()
	hour, min, sec = t.Clock()
	// 解析相应的 %Y %m %d %H %M %S 然后返回信息
	// %Y 替换成 2012
	// %m 替换成 10
	// %d 替换成 24
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}
