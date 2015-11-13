package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("请求的方法: %s", r.Method)
	//		if r.Method != "POST" {
	//			http.NotFound(w, r)
	//		}
	var desc, token string
	if r.Method == "POST" {
		desc = strings.TrimSpace(r.PostFormValue("desc"))
		token = strings.TrimSpace(r.PostFormValue("token"))
	} else if r.Method == "GET" {
		desc = strings.TrimSpace(r.FormValue("desc"))
		token = strings.TrimSpace(r.FormValue("token"))
	}
	log.Printf("token : %s", token)

	se := session.Copy()
	defer se.Close()
	if err := session.DB(DB).C(POST).Insert(&Post{bson.NewObjectId(), desc}); err != nil {
		panic(err)
	}

	//	showPosts(w, se)
	http.Redirect(w, r, "/post/list", 302)
}

func listPostHandler(w http.ResponseWriter, r *http.Request) {

	se := session.Copy()
	defer se.Close()

	showPosts(w, se)
}

func delPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("del post\n"))
}

func showPosts(w http.ResponseWriter, se *mgo.Session) {
	posts, err := listPost(se)
	if err != nil {
		log.Printf("list post err : %v", err)
		http.Error(w, "出错了", 500)
	}

	t, err := template.ParseFiles("temp/postlist.html")
	if err != nil {
		log.Printf("parse file err : %v", err)
	}
	err = t.Execute(w, posts)
	if err != nil {
		log.Printf("execute err : %v", err)
	}
}
