package http

import (
	"fmt"
	"log"
	"net/http"
)
func homePage(res http.ResponseWriter, req *http.Request){
    fmt.Fprintf(res, "Salom HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
func about(res http.ResponseWriter, req *http.Request){
	data := req.URL.Path[0:]
    res.Write([]byte("About me !"))
	res.Write([]byte(data))
}
func handleRequests() {
    http.HandleFunc("/", homePage)
	http.HandleFunc("/about", about)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
func Server(){
	handleRequests()
}