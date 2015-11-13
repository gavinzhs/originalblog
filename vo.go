package main

import "gopkg.in/mgo.v2/bson"

type Post struct {
	Id   bson.ObjectId `bson:"_id" json:"id"`
	Desc string        `bson:"desc" json:"desc"`
}

func initDBIndex() {
	//    session.DB(DB).C(POST).EnsureIndex()
}
