// lista.go
// exemplo 11 livro: Programando em Go
/*
exemplo de uso (terminal):

    $ go run lista.go
*/

package main

import "fmt"

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {

	l := *lista
	removido := l[indice]
	*lista = append(l[0:indice], l[indice+1:]...)

	return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{} {
	return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFim() interface{} {
	return lista.RemoverIndice(len(*lista) - 1)
}

func main() {
	lista := ListaGenerica{1, "café", 42, true, 23, "Bola", 3.14, false}

	fmt.Printf("Lista original: \n%v \n\n", lista)

	fmt.Printf("removendo do inicio: %v, apos remocao: \n %v \n\n", lista.RemoverInicio(), lista)

	fmt.Printf("removendo do fim: %v, apos remocao: \n %v \n\n", lista.RemoverFim(), lista)

	fmt.Printf("removendo do Indice: %v, apos remocao: \n %v \n\n", lista.RemoverIndice(3), lista)

	fmt.Printf("removendo do indice: %v, apos remocao: \n %v \n\n", lista.RemoverIndice(0), lista)

	fmt.Printf("removendo do ultimo indice: %v, apos remocao: \n %v \n\n", lista.RemoverIndice(len(lista)-1), lista)

}
