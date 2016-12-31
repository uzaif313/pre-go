package main

import (
	"fmt"
)
type rubyist struct {
	name string
	exp int
}
func (g rubyist) getPosition() string{
	if (g.exp > 4){
	 return "Senior Dev"
	}else{
	 return "Beginer Dev"
	}
}
func main() {
	matt   := rubyist {name:"Matt",exp:20}
	robert := rubyist {name:"Robert",exp:8}
	uzaif  := rubyist {name:"uzaif",exp:1}
	fmt.Println(matt.getPosition())
	fmt.Println(uzaif.getPosition())
	fmt.Println(robert.getPosition())
}

