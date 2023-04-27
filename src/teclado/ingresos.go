package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var legend string
var err error

func IngresoNumeros() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1")
	if scan.Scan() {
		number1, err = strconv.Atoi((scan.Text()))
		if err != nil {
			panic("CRASH! DATO incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2")
	if scan.Scan() {
		number2, err = strconv.Atoi((scan.Text()))
		if err != nil {
			panic("CRASH! DATO incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese Leyenda")
	if scan.Scan() {
		legend = scan.Text()
	}

	fmt.Println(legend, number1*number2)
}
