package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

var session *mgo.Session

const DB string = "test"

const POST string = "post"

func initDB() {
	var err error
	if session, err = mgo.Dial(fmt.Sprintf("testadmin:testadmin@127.0.0.1:27017/%s", DB)); err != nil { //todo 数据库连接管理
		log.Fatalf("mgo.Dial err : %v", err)
	}
}

//post
