package src

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)

	//注册定义的model
	orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	//创建table
	orm.RunSyncdb("default", false, true)

}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	//插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//更新表
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

}
