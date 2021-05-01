package main

import "fmt"
func add(number1 float32, number2 float32) float32{

	return number1 + number2
}

func subtract(number1 float32, number2 float32) float32 {
	return number1 - number2
}

func mul(number1 float32, number2 float32) float32{

	return number1 * number2
}

func div(number1 float32, number2 float32) float32 {
	return number1/number2
}
func main()  {
	fmt.Println("Enter First number : ")
	var number1,number2 float32
	var err error
	_, err =fmt.Scanln(&number1)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Enter Second number : ")
	_, err =fmt.Scanln(&number2)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	var operator string
	fmt.Println("Enter operator : ")
	_, err =fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	switch operator{
	case "+":
		fmt.Println("Result is : ", add(number1,number2))
	case "*":
		fmt.Println("Result is : ", mul(number1,number2))
	case "-":
		fmt.Println("Result is : ", subtract(number1,number2))
	case "/":
		fmt.Println("Result is : ", div(number1,number2))
	default:
		fmt.Println("Error operator.")
	}
}