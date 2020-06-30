package src

//目前 Go 支持 mongoDB 最好的驱动就是 mgo，这个驱动目前最有可能成为官方的 pkg。
//http://labix.org/mgo
// 安装go get gopkg.in/mgo.v2

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Persion{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}
	result := Persion{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)
}
