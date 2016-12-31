package main

import (
	"fmt"
)

func main() {
	//  String Array with size 3
	//	var techStack [3]string
	//	techStack[0] = "MEAN"
	//	techStack[1] = "MERN"
	//	techStack[2] = "React on Rails"
	//	fmt.Println(techStack)
	
	 // Slice declaration 
	 // var techStack []string
	  
	 // techStack  = append(techStack,"MERN")
	 // techStack = append(techStack,"Nodejs & express")
	 // techStack = append(techStack,"MEAN ")
	 //techStack = append(techStack,"socket.io with react")
	 // fmt.Println(techStack)
	 
	 //Declare and initailze slice
	 var techStack = []string{"Rails","Spring Boot","MEAN","LAMP","MERN"}
	 //iterate slice with for loop using range
	 for i:= range(techStack) {
		fmt.Println(techStack[i])
	 }
 }

