// Bu mening birinch kamentaryam
package main

import (
	"fmt"
	//"strconv"
	"strings"
)

// interface bir funksiyadan turli tiplar yordamida chaqirishga yordam beradi bu misoldagi Interface funksiyasi
// go'yoki olingan tiplarimizning maydonlariga bordek. 16- qatordagi checker funksiyasi interface tipida bo'lmanligi
// sababli tiplarimiz uni tanimadi
type speker interface {
	talk() string
}
type uzb struct{}
type laser int // bu ham bir tip
type eng struct{}

func (e *eng) talk() string {
	return ("hello")
}
func (u *uzb) talk() string {
	return ("salom")
}
func (l laser) talk() string {
	return strings.Repeat("toot ", int(l))
}
// func checker() string {
// 	return "check "
// }
func shout(t speker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// type starship struct {
// 	laser 
// }
func Salom() {
	var e eng
	var u uzb
	 l := laser(3)
	fmt.Println(e.talk())
	fmt.Println(u.talk())
	fmt.Println(l.talk())
	shout(laser(2))
	shout(&eng{})
	shout(&uzb{})

}
