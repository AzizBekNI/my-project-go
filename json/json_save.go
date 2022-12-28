package json

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Name struct {
	Family   string
	Personal string
}
type Email struct {
	Id     int
	Adress string
}
type Interest struct {
	Id    int
	Hobby string
}
type Person struct {
	Id        int
	FirstName string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.Name.Family + " ID : " + strconv.Itoa(p.Id)
}
func WriteMsg(msg interface{}){
	fmt.Println(msg)
}
func GetPersonEmail(p *Person, in int) {
	WriteMsg( p.Email[in].Adress)
}
func writeStart() {
	WriteMsg( "*----------------------------------*")
}
func checkError(err error) {
	if err != nil {
		WriteMsg(err.Error())
		os.Exit(1)
	}
}
func saveJson(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	defer outFile.Close()
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}
func Personal() {
	p := Person{
		Id:        1,
		FirstName: "Azizbek",
		Name: Name{
			Family:   "Azik",
			Personal: "Aziz",
		},
		Email: []Email{
			{Id: 1, Adress: "namozovazizbek777@gmail.com"},
			{Id: 2, Adress: "namozovazizbek@gmail.com"},
			{Id: 3, Adress: "namozovazizbek007@gmail.com"}},
		Interest: []Interest{
			{Id: 2, Hobby: "none"},
			{Id: 4, Hobby: "none"}},
	}
	WriteMsg("Start")
	writeStart()
	res := GetPerson(&p)
	WriteMsg(res)
	writeStart()
	i := 0
	GetPersonEmail(&p, i)
	writeStart()
	saveJson("./file_uchun/json.txt", &p)
}
