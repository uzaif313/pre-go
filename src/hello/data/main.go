package data
type getPositioner interface {
	GetPosition() string
}
type rubyist struct {
	name string
	exp int
	isExpert bool
}
type gopher struct{
	name string
	gitStar int
}


func (g gopher) GetPosition() string{
	if (g.gitStar > 4){
	 return "Intermdiate folk"
	}else{
	 return "Beginer folk"
	}
}


func (g rubyist) GetPosition() string{
	if (g.exp > 4){
	 return "Senior Dev"
	}else{
	 return "Beginer Dev"
	}
}


func GetDevList()[] getPositioner{
	matt   :=  rubyist {name:"Matt",exp:20}
	robert :=  rubyist {name:"Robert",exp:8}
	uzaif  :=  gopher {name:"uzaif",gitStar:10}
	aman   :=  rubyist {name:"aman",exp:0}

	list := []getPositioner{&matt, &robert, &uzaif, &aman}
	return list
}
