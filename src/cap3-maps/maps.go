// maps.go
// exemplo 4 livro: Programando em Go
/*
exemplo de uso (terminal):

    $ go run maps.go teste casa cadeira casamento teste s olhos s abismo rua rua lua sua 
            palavra:  casa  | ocorrencia:  1
            palavra:  olhos  | ocorrencia:  1
            palavra:  rua  | ocorrencia:  2
            palavra:  lua  | ocorrencia:  1
            palavra:  teste  | ocorrencia:  2
            palavra:  cadeira  | ocorrencia:  1
            palavra:  casamento  | ocorrencia:  1
            palavra:  s  | ocorrencia:  2
            palavra:  abismo  | ocorrencia:  1
            palavra:  sua  | ocorrencia:  1
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func allLower(str []string) []string {

	for i, v := range str {

		str[i] = strings.ToLower(v)

	}

	return str
}

func gerarEstatistica(p []string) map[string]int {

	// cria um map (hash table)
	dict := make(map[string]int)

	for _, j := range p {

		//retorna o valor, um erro se não existir a chave
		_, fin := dict[j]

		if fin {
			dict[j] += 1
		} else {
			dict[j] = 1
		}
	}

	return dict
}

func imprime(dict map[string]int) {

	for i, j := range dict {
		fmt.Println("palavra: ", i, " | ocorrencia: ", j)
	}
}

func main() {

	// armazena as palavras
	p := os.Args[1:]

	// converte tudo para minusculo
	p = allLower(p)

	imprime(gerarEstatistica(p))

}
