package src

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
	"time"
)

type User struct {
	Uid     int      `orm:"PK"` //如果表的主键不是id, 那么需要加上pk注释, 显示的说这个字段是逐渐
	Name    string   `orm:"size(100)"`
	Profile *Profile `orm:"rel(one)"`      //OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` //设置一对多的反向关系
}

type Userinfo struct {
	Uid        int `orm:"PK"`
	Username   string
	Departname string
	Created    time.Time
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` //设置一对一的方向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User   `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tage `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

//注意一点，beego orm 针对驼峰命名会自动帮你转化成下划线字段，
//例如你定义了 Struct 名字为 UserInfo，那么转化成底层实现的时候是 user_info，
//字段命名也遵循该规则。

func init() {
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)

	//注册定义的model
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Tag))
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

	// 读取one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	//删除表
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//插入数据
	o := orm.NewOrm()
	var user User
	user.Name = "zxxx"
	user.Departname = "zxxx"
	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
	//插入后 user.Uid 就是插入成功之后的自增ID
	//同时插入多个对象: InsertMulti
	//类似sql语句
	// insert info table (name, age) values("slene", 28),("astaxie", 30),("unknown", 20)

	//第一个参数bulk 为并列插入的数量, 第二个为对象的slice
	//返回值为成功插入的数量

	//users := []User{
	//	{Name: "slene"},
	//	{Name: "astaxie"},
	//	{Name: "unknown"},
	//	...
	//}
	//successNums, err := o.InsertMulti(100, users)
	//bulk为1时, 将会顺序插入slice中的数据

	//更新数据
	//o := orm.NewOrm()
	//user := User{Uid:1}
	//if o.Read(&user) == nil {
	//	user.Name = "MyName"
	//	if num, err := o.Update(&user); err == nil {
	//		fmt.Println(num)
	//	}
	//}

	//Update 默认更新所有字段 可以更新制定字段
	// 只更新Name
	// o.Update(&user, "Name")
	//指定多个字段
	//o.Update(&user, "Field1", "Field2", ...)
	//// Where: 用来设置条件，支持多个参数，第一个参数如果为整数，相当于调用了 Where ("主键 =?", 值)。

	// 查询数据
	// beego orm 的查询接口毕竟灵活
	//example 1 根据主键获取数据
	//o := orm.NewOrm()
	//var user User
	//user := User{Id: 1}
	//err = o.Read(&user)
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(user.Id, user.Name)
	//}

	//example 2
	//o := orm.NewOrm()
	//var user User
	//qs := o.QueryTable(user) //返回 QuerySeter
	//qs.Filter("id", 1) // WHERE id = 1
	//qs.Filter("profile__age", 18) // WHERE profile.age = 18

	//example 3 WHERE IN 查询条件
	//qs.Filter("profile_age_in", 18, 20)
	// WHERE profile.age IN (18, 20)

	//example 4 更加复杂的条件
	//qs.Filter("profile_age_in", 18, 20).Exclude("profile_lt", 1000)
	// WHERE profile.age IN (18, 20) AND NOT profile_id < 1000

	//可以通过接口获取多条数据
	// exa1 根据条件 age > 17, 获取20位置开始的10条数据
	//var allusers []User
	//qs.Filter("profile_age_gt", 17)
	//WHERE profile.age > 17

	// exa2 limit 默认从10开始, 获取10条数据
	//qs.Limit(10, 20)
	//LIMIT 10 OFFSET 20 注意 和SQL 相反

	//删除数据
	//exa 1 删除单条数据
	//o := orm.NewOrm()
	//if num, err := o.Delete(&User{Id: 1}); err == nil {
	//	fmt.Println(num)
	//}
	//Delete 操作会对反向关系进行操作，此例中 Post 拥有一个到 User 的外键。
	//删除 User 的时候。如果 on_delete 设置为默认的级联操作，将删除对应的 Post

	//关联查询
	//type Post struct {
	//	Id int `orm:"auto"`
	//	Title string `orm:"size(100)"`
	//	User *User `orm:"rel(fk)"`
	//}
	//
	//var posts []*Post
	//qs := o.QueryTable("post")
	//num, err := qs.Filter("User__Name")

	//Group By 和 Having
	//qs.OrderBy("id", "-profile__age")
	//Order by id asc , profile.age DESC

	//qs.OrderBy("-profile__age", "profile")
	// ORDER BY profile.age DESC, profile_id ASC

	//GroupBy: 用来指定进行 groupby 的字段
	//Having: 用来指定having执行的时候的条件

	//使用原生sql
	//o := orm.NewOrm()
	//var r orm.RawSeter
	//r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")

	//复杂原生 sql使用

	//func (m *User) Query(name string) user []User{
	//	var o orm.Ormer
	//	var rs orm.RawSeter
	//	o = orm.NewOrm()
	//	rs = o.Raw("SELECT * FROM user "+
	//	"WHERE name=? AND uid>10 "+
	//	"ORDER BY uid DESC "+
	//	"LIMIT 100", name)
	//	// var user []User
	//	num, err := rs.QueryRows(&user)
	//	if err != nil{
	//		fmt.Println(err)
	//	} else{
	//		fmt.Println(num)
	//		// return user
	//	}
	//	return
	//}

}
