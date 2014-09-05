// estados.go
// exemplo 8 livro: Programando em Go
/*
exemplo de uso (terminal):

    $ go run estados.go
*/

package main

import "fmt"

type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
    
    //cria um map onde a chave é do tipo string e o valor é a struct Estado de tam. máx.: 6
	estados := make(map[string]Estado, 6)

	// insere os dados no map
	estados["GO"] = Estado{"Goias", 6434052, "Goiania"}
	estados["PB"] = Estado{"Paraiba", 3914418, "Joao pessoa"}
	estados["PR"] = Estado{"Parana", 10997462, "Curitiba"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
	estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}

	//imprime o map completo
	fmt.Println(estados)

	//acesso um valor pela chave
	fmt.Println(estados["PR"])

	//verifica se existe a chave no map
	valor, ok := estados["GO"]

	if ok {
		fmt.Println("existe")
		fmt.Println(valor)
	}

	// outro exemplo
	valor, ok = estados["GK"]

	if !ok {
		fmt.Println("chave inexistente, ", ok)
		fmt.Println(valor)
	}
    
    //acessando atributo da struct
    fmt.Println("paraiba\n populacao: ", estados["PB"].populacao)
    
    
}
