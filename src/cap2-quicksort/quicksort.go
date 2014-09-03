// quicksort.go
// exemplo 3 livro: Programando em Go
/* 
exemplo de uso (terminal):

    $ go run quicksort.go 6 5 1 0 9
        [0 1 5 6 9]
*/


package main

import (
	"fmt"
	"os"
	"strconv"
)

func quicksort(d []int) []int {

	if len(d) <= 1 {
		return d
	}

	auxd := make([]int, len(d))
	copy(auxd, d)

	idPivo := len(auxd) / 2
	pivo := auxd[idPivo]
	auxd = append(auxd[:idPivo], auxd[idPivo+1:]...)

	menores, maiores := particionar(auxd, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)

}

func particionar(nums []int, pivo int) (menores []int, maiores []int) {

	for _, n := range nums {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}

func main() {

	args := os.Args[1:]
	d := make([]int, len(args))

	//converte os d para int e add no array
	for i, j := range args {
		n, err := strconv.Atoi(j)

		if err != nil {
			fmt.Println("erro")
			os.Exit(1)
		}

		d[i] = n
	}

	//chama o quicksort
	fmt.Println(quicksort(d))

}
