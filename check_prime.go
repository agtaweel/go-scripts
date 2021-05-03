package main

import "fmt"

func main()  {
	var number, flag, temp int
	var err error
	fmt.Println("Enter the Number to check Prime : ")
	_,err = fmt.Scanln(&number)
	if err != nil || number<2{
		fmt.Println("Error scanning")
		return
	}
	temp = number/2
	for i := 2;i<temp;i++ {
		if number%i ==0{
			flag =1
			break
		}
	}
	if flag ==0{
		fmt.Println(number, " is prime number!")
	}else {
		fmt.Println(number, " is Not prime number!")
	}
}