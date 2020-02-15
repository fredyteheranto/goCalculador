package main

import"fmt"

func main() {
	const EXIT_OPTION int = 5
	var option int = 0

	printWelcome()

	for option != EXIT_OPTION {
		option = setMenu()

	}
}

func printWelcome() {
	var welcome string = fmt.Sprint(
		"*........................................*\n",
		"*                                        *\n",
		"*               CALCULADORA              *\n",
		"*                                        *\n",
		"*........................................*\n\n",
		"ESTA ES UN CALCULADORA CON GOLANG\n",
	)

	fmt.Println(welcome)
}

func setMenu() int {
	var option int
	var menu string = fmt.Sprint(
		"Seleccione una opcion: \n",
		"[1] Suma \n",
		"[2] Resta \n",
		"[3] Multiplicacion \n",
		"[4] Divicion \n",
		"[5] Salir \n\n",
		"> Seleccione una opcion: ",
	)

	fmt.Printf(menu)
	fmt.Scanf("%d", &option)

	return option
}

func evaluateOption(option int) {
	switch option {
	case 1:
		add()
	case 2:
		substract()
	case 3:
		multiply()
	case 4:
		divide()
	case 5:
		exit()
	default:
		fmt.Printf("  Wrong operation. Try again")
	}
}
func add() {
	var (
		firstNumber  int
		secondNumber int
		result int
	)

	fmt.Print("  Your option was sum\n\n")
	fmt.Print("  First number: ")
	fmt.Scanf("%d", &firstNumber)

	fmt.Print("  Second number: ")
	fmt.Scanf("%d", &secondNumber)

	result = firstNumber + secondNumber

	fmt.Printf("\n  The sum of %d and %d is %d\n\n", firstNumber, secondNumber, result)
}
func substract() {
	var (
		firstNumber  int
		secondNumber int
		result int
	)

	fmt.Print("  Your option was substraction\n\n")
	fmt.Print("  First number: ")
	fmt.Scanf("%d", &firstNumber)

	fmt.Print("  Second number: ")
	fmt.Scanf("%d", &secondNumber)

	result = firstNumber - secondNumber

	fmt.Printf("  The substraction of %d and %d is %d\n\n", firstNumber, secondNumber, result)
}

func multiply() {
	var (
		firstNumber  int
		secondNumber int
		result int
	)

	fmt.Print("  Your option was multiplication\n\n")
	fmt.Printf("  First number: ")
	fmt.Scanf("%d", &firstNumber)

	fmt.Print("  Second number: ")
	fmt.Scanf("%d", &secondNumber)

	result = firstNumber * secondNumber

	fmt.Printf("  The multiplication of %d and %d is %d\n\n", firstNumber, secondNumber, result)
}
func divide() {
	var (
		firstNumber  int
		secondNumber int
		result int
		modulus      int
	)

	fmt.Print("  Your option was division\n\n")
	fmt.Print("  First number: ")
	fmt.Scanf("%d", &firstNumber)

	fmt.Print("  Second number: ")
	fmt.Scanf("%d", &secondNumber)

	result = firstNumber / secondNumber
	modulus = firstNumber % secondNumber

	fmt.Printf("  The division of %d and %d is %d and the modulus is %d\n\n", firstNumber, secondNumber, result, modulus)
}

func exit() {
	var exitMsg = fmt.Sprint(
		"\nThanks for using the minimalist calculator!\n",
		"We hope to see you again soon :)\n",
	)

	fmt.Println(exitMsg)
}
