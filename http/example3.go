package http

import (
	"net/http"
	"strconv"
)

func Main_() {
	http.HandleFunc("/search", search)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":3000", nil)
}
func search(res http.ResponseWriter, req *http.Request) {
	sourc := req.FormValue("sourc")
	q := req.FormValue("q")
	data := "/search?q=" + q + "&sourc=" + sourc
	res.Write([]byte(data))
}
func add(res http.ResponseWriter, req *http.Request) {
	s1 := req.FormValue("s1")
	s2 := req.FormValue("s2")
	x, _ := strconv.Atoi(s1)
	x2, _ := strconv.Atoi(s2)
	data := x + x2
	//strconv.Itoa(data)
	
	res.Write([]byte(strconv.Itoa(data)))
}
