package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	Scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese el Número a Multiplicar")
		if Scanner.Scan() {
			numero, err := strconv.Atoi(Scanner.Text())
			if err != nil {
				fmt.Println("Hubo un error")
				continue
			} else {
				for i := 0; i <= 10; i++ {
					fmt.Println(numero, "*", i, "=", (numero * i))
				}
			}
			break
		}
	}
}
