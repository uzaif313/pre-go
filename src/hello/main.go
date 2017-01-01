package main

import (
	"fmt"
	"os"
)

func main() {
		
		//Dynamic data- type inference
			args := os.Args 
			msg  := "Let's make realtime program"
		//Data type Declaration and init
			//var args [] string
			//args = os.Args
			var name string
			name = "Welcome to world of gopher"
		//conditional if in go
		if len(args) > 1{
			if args[1] == "1" {
				fmt.Println(name)
			}else{
				fmt.Println(args[1])
			}
		}else{
			fmt.Println(msg)
		}
		
}
