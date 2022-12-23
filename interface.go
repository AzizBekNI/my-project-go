// Bu mening birinch kamentaryam
package main

import "strconv"
type speker interface{
// interface bir funksiyadan turli tiplar yordamida chaqirishga yordam beradi bu misoldagi Interface funksiyasi
// go'yoki olingan tiplarimizning maydonlariga bordek. 16- qatordagi checker funksiyasi interface tipida bo'lmanligi
// sababli tiplarimiz uni tanimadi
	Interface() string 
} 
type uzb struct {} 
type laser int // bu ham bir tip 
type eng struct {}
func (e *eng)Interface() string {
	return ("hello")
}
func (u *uzb)Interface() string {
	return ("salom")
}
func (l *laser)Interface() string{
	return strconv.Itoa(8)
} 
func checker() string {
	return "check "
}
func Salom(){
	var e eng
	var u uzb
	var l laser 
	print(e.Interface())
	print(u.Interface())
	print(l.Interface())

}