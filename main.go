package main

import (
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	print("original blog start")

	createPostHandler := func(w http.ResponseWriter, r *http.Request) {
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
		if err := session.DB(DB).C(POST).Insert(&Post{bson.NewObjectId(), desc}); err != nil {
			panic(err)
		}
		w.Write([]byte("创建post\n"))
	}

	listPostHandler := func(w http.ResponseWriter, r *http.Request) {
		//		t := template.New("listpost")
		t, err := template.ParseFiles("temp/postlist.html")
		if err != nil {
			log.Printf("parse file err : %v", err)
		}
		err = t.Execute(w, "no need data")
		if err != nil {
			log.Printf("execute err : %v", err)
		}
		//		w.Write([]byte("list post\n"))
	}

	delPostHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("del post\n"))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome\n"))
	})
	http.HandleFunc("/post/create", createPostHandler)
	http.HandleFunc("/post/list", listPostHandler)
	http.HandleFunc("/post/del", delPostHandler)

	initDB()
	initDBIndex()

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("listener and serve err : %v", err)
	}
}
