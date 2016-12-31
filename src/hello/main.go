package main

import (
	"fmt"
)
type rubyist struct {
	name string
	exp int
	isExpert bool
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
	aman  := rubyist {name:"aman",exp:0}
	//fmt.Println(matt.getPosition())
	//fmt.Println(uzaif.getPosition())
	//fmt.Println(robert.getPosition())
	var devs = []rubyist {uzaif, matt, robert, aman}
	for i:=range(devs){
		checkExpert(&devs[i])
		fmt.Println(devs[i])
	}

}
func checkExpert(r *rubyist){
	r.isExpert = r.exp > 3
}
