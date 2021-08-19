package views

import (
	"net/http"
	"path"
)

type Context struct {
}

func Helloworld(response http.ResponseWriter, request *http.Request) {
	// vars := mux.Vars(request)
	// response.WriteHeader(http.StatusOK)
	// fmt.Fprintf(response, "Hello World")
	// response.Write([]byte("Hello world"))
	// http.FileServer(http.Dir("./static"))
	// p := "./static/index.html"
	// // http.ServeFile(response, request, p)
	// templates := template.New("static")
	// templates.New("doc").Parse(p)
	// context := Context{}
	// templates.Lookup("doc").Execute(response, context)
	p := path.Dir("./static/index.html")
	// set header
	response.Header().Set("Content-type", "text/html")
	http.ServeFile(response, request, p)
}
