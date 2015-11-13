package main

import (
	"html/template"
	"log"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	temp, _ := template.ParseFiles("temp/postlist.html")
	err := temp.Execute(os.Stdout, "")
	if err != nil {
		log.Printf("execute err : %v", err)
	}
}
