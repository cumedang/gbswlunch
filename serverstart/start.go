package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/cumedang/gbswlunch/lunch"
)

const port string = ":8000"

var sanz Sanz
var a string

type Sanz struct {
	ddata string
	hdata string
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./serverstart/home.gohtml"))
	switch r.Method {

	case "POST":
		r.ParseForm()
		sanz.ddata = r.Form.Get("day")
		sanz.hdata = r.Form.Get("how")
		a = lunch.Lunch(sanz.ddata, sanz.hdata)
	}
	tmpl.Execute(rw, a)

}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
