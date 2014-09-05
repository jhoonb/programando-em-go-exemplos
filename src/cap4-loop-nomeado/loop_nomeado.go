// loop_nomeado.go
// exemplo 7 livro: Programando em Go
/*
exemplo de uso (terminal):

    $ go run loop_nomeado.go
*/

package main

import "fmt"

func main() {

	var i int

loop:
	for i = 0; i < 10; i++ {
		fmt.Printf("for i = %d\n", i)

		switch i {
		case 2, 3:
			fmt.Println("Quebrando o swich, i == %d", i)
			break
		case 5:
			fmt.Println("Quebrando o loop i == 5")
			break loop
		}
	}

	fmt.Println("fim")
}
