package main

import (
	"log"
	"net/http"
)

func main() {
	print("original blog start")

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
