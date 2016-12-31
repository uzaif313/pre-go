package main

import (
	"fmt"
	"time"
	"os"
	"errors"
)

func main() {
	hourOfDay := time.Now().Hour()
	greeting,err  := getGreeting(hourOfDay)
	    if err != nil{	
		 fmt.Println(err)
		 os.Exit(1)
	    }else{
		 fmt.Println(greeting)
		}
}

func getGreeting(hour int)(string, error) {
	var msg string
	if hour < 7 {
		err := errors.New("Too Early for greet please sleep now")
		fmt.Println(err)	
		return msg,err
	}
	if hour < 12{
		msg = "Good Morning"
	}else if hour < 18 {
		msg = "Good Afternoon"
	}else {
		msg = "Good Evening"
	}
	return msg,nil
}