// conversor.go
// exemplo 2 livro: Programando em Go
/*
exemplo de uso (terminal):

    $ go run conversor.go 10 12.5 6.3 q

        valor quilometro:  10  - milha:  16093.470878864446
        valor quilometro:  12.5  - milha:  20116.838598580554
        valor quilometro:  6.3  - milha:  10138.8866536846

*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	nargs := len(os.Args)

	if nargs < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	// unidade
	unOrigem := os.Args[nargs-1]
	// arrays
	vrOrigem := os.Args[1 : nargs-1]

	// verifica as unidades

	switch unOrigem {

	case "c", "f", "m", "q":
		break
	default:
		fmt.Println("Unidade desconhecida: ", unOrigem, "valido apenas: c, f, m, q")
		os.Exit(1)
	}

	for _, v := range vrOrigem {
		vr, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Println("Valor nao convertido para float64")
			os.Exit(1)
		}

		switch unOrigem {

		case "c":
			fmt.Println("valor celcius: ", vr, "- farh:", ((vr * 1.8000) + 32))

		case "f":
			fmt.Println("valor farh: ", vr, " - celcius:", ((vr - 32) * 1.8000))

		case "q":
			fmt.Println("valor quilometro: ", vr, " - milha: ", (vr / 0.00062137))

		case "m":
			fmt.Println("valor milha: ", vr, " - quilometro: ", (vr / 0.62137))

		default:
			fmt.Println("erro")
			os.Exit(1)
		}
	}
}
