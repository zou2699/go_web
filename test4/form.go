package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
	"learn_web/test3"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("./test4/login.gtpl")
		log.Println(t.Execute(w, token))
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
			username := r.Form["username"][0]
			password := r.Form["password"][0]
			if username == "zouhl" && password == "123" {
				fmt.Fprintf(w, "welcome zouhl")
			} else {
				fmt.Fprintf(w, "something error")
			}
		} else {
			fmt.Println("no token!")
			fmt.Fprintf(w, "no token")
		}
	}

}

func upload(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		log.Println("curtime:", curtime)
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("./test4/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(30 << 22)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test4/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}

func main() {
	http.HandleFunc("/", test3.SayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
