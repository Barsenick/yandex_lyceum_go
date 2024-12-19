package main

import (
	"fmt"
	"net/http"
	"strings"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	lower := strings.ToLower(name)
	switch lower {
	case "":
		fmt.Fprint(w, `Я не знаю кто ты. Добавь в конце URL "?name=*твоё имя*". Если я тебя знаю ты увидишь секретное сообщение.`)
	case "николай":
		fmt.Fprint(w, `абаюнда🥵`)
	case "коля":
		fmt.Fprint(w, `абаюнда🥵`)
	case "сергей":
		fmt.Fprint(w, `го в рокет?`)
	case "sigma":
		fmt.Fprint(w, `Ладно. В конце URL перед "?name=" добавь "/secret/sigma".`)
	case "julianiyaggnew":
		fmt.Fprint(w, `Sigma Rizz on Kondicii`)
	default:
		fmt.Fprint(w, `либо я тебя не знаю либо было лень добавлять лол`)
	}
}

func SigmaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `этот тип думал что тут будет секретка`)
}

func main() {
	http.HandleFunc("/", NameHandler)
	http.HandleFunc("/secret/sigma", SigmaHandler)
	http.ListenAndServe(":8080", nil)
}
