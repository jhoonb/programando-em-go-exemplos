// loop_infinito.go
// exemplo 6 livro: Programando em Go
/*
exemplo de uso (terminal):

    $ go run loop_infinito.go
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	n := 0

	for {
		// n = n + 1
		n++

		i := rand.Intn(4200)
		fmt.Println(i)

		if i%42 == 0 {
			break
		}
	}
	fmt.Printf("Saída após %d iterações. \n", n)
}
