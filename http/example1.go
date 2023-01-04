package http

import (
	"io"
	"net/http"
)
func Handle_(){
	var t tash
	var i ist
	mux := http.NewServeMux()
	mux.HandleFunc("/tashkent", t.handle_)
	mux.HandleFunc("/istanbul", i.handle_)
	http.ListenAndServe(":3000", mux)
}
type tash int
func (city tash)handle_(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "Salom Tashkent")
}
type ist int
func (city ist)handle_(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "Salom Istanbul")
}
