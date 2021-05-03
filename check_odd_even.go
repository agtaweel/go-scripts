package main

import (
	"fmt"
)

func main()  {
	var err error
	fmt.Println("Enter Number to check odd/even : ")
	var number int
	_,err = fmt.Scanln(&number)
	if err != nil{
		fmt.Println("Error scanning....")
	}
	if number%2 == 0{
		fmt.Println(number," is even!")
	}else {
		fmt.Println(number, " is odd!")
	}
}