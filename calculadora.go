package myCalculator

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MathStruct struct{}

func (MathStruct) Operate(entrada string, operador string) int {
	numeros := strings.Split(entrada, operador)
	num1 := parsear(numeros[0])
	num2 := parsear(numeros[1])

	switch operador {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		println(operador + " no está soportado")
		return 0
	}
}

func parsear(entrada string) int {
	numero, _ := strconv.Atoi(entrada) //convierte de string a enterow
	return numero
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

//subir el modulo a git:(crear el repo primero en github)
//go mod init github.com/usergit/reponame
