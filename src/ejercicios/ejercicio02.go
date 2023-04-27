package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	var number int
	var err error
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingresa un numero y te dare su tabla de multiplicar:")
		if scan.Scan() {
			number, err = strconv.Atoi((scan.Text()))
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d \n", number, i, i*number)
	}

}
