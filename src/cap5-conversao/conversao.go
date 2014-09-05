// conversao.go
// exemplo 10 livro: Programando em Go
/*
exemplo de uso (terminal):

    $ go run conversao.go
*/

package main

import "fmt"

type ListaDeCompras []string

func imprimirSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

func imprimirLista(lista ListaDeCompras) {
	fmt.Println("Lista de compras: ", lista)
}

// exemplo de conversao entre tipos
func main() {
	lista := ListaDeCompras{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}

	imprimirSlice([]string(lista))
	imprimirLista(ListaDeCompras(slice))
}
