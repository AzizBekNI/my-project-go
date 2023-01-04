package http

import (
	"net/http"
)

func Vazifa() {
	//http.HandleFunc("/password",password)
	http.HandleFunc("/auto", login)
	http.ListenAndServe(":3000", nil)
}

//	func password(res http.ResponseWriter, req *http.Request){
//		password := req.FormValue("password")
//		res.Write([]byte(password))
//	}
func login(res http.ResponseWriter, req *http.Request) {
	password := req.FormValue("password")
	login := req.FormValue("login")
	if login == "admin" && password == "123456" {
		res.Write([]byte("Xush kelibsiz !"))
	} else {
		res.Write([]byte("login yoki parol xato"))
	}
}
