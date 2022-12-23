package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}
type rover struct{}

func (r rover) talk() string {
	return "whir whir"
}
func shouter(t talker) {
	lounder := strings.ToUpper(t.talk())
	fmt.Println(lounder)
}
func Help(){
	t := rover{}
	shouter(t)
}
