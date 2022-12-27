package self

import "fmt"

type location struct {
	long, lat interface{}
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func GetLocation(){
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity)
}