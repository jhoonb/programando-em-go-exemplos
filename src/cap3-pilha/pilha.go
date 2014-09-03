// pilha.go
// exemplo 5 livro: Programando em Go
/*
exemplo de uso (terminal):

    $ go run pilha.go
*/

package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	values []interface{}
}

func (s Stack) Length() int {

	return len(s.values)
}

func (s Stack) Empty() bool {
	return s.Length() == 0
}

func (s *Stack) Add(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("Stack Empty!")
	}

	value := s.values[s.Length()-1]
	s.values = s.values[:s.Length()-1]
	return value, nil
}

func main() {

	stack := Stack{}

	fmt.Println("pilha criada com tamanho:", stack.Length())
	fmt.Println("Vazia?", stack.Empty())

	stack.Add("Go")
	stack.Add(2009)
	stack.Add(3.14)
	stack.Add("final")

	fmt.Println("tamanho: ", stack.Length())
	fmt.Println("Vazia?: ", stack.Empty())

	for !stack.Empty() {

		i, _ := stack.Pop()
		fmt.Println("pop: ", i)
		fmt.Println("length: ", stack.Length())
		fmt.Println("Vazia?: ", stack.Empty())
	}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("erro", err)
	}
}
