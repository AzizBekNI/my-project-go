package http

import (
	"fmt"
	"net/http"
)
func HandleMsg(){
mux := http.NewServeMux()
msg1 := &messageHandler{"Its first msg"}
msg2 := &messageHandler{"It's second msg"}
mux.Handle("/bir", msg1)
mux.Handle("/ikki", msg2)
http.ListenAndServe(":3000", mux)
}
type messageHandler struct{
	message string
}
func (msg *messageHandler)ServeHTTP(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, msg.message)
}