package main

import (
	"fmt"
	"./data"
)


func main() {
	devList := data.GetDevList();
	for _,dev := range devList {
		fmt.Println(dev)
	}
}
