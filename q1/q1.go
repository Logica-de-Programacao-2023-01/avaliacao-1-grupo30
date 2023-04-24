package q1

//Em um dia quente de verão, Pete e seu amigo Billy decidiram comprar uma melancia. Eles escolheram a maior e mais
//saborosa, na opinião deles, e, em seguida, pesaram a fruta nas balanças, obtendo seu peso em quilos. Morrendo de sede,
//correram para casa com a melancia e decidiram dividi-la. No entanto, enfrentaram um problema difícil.
//
//Como grandes fãs de números pares, Pete e Billy querem dividir a melancia de tal forma que cada uma das duas partes pese
//um número par de quilos, sem que as partes necessariamente tenham pesos iguais. Mas os meninos estão extremamente
//cansados e querem começar a refeição o mais rápido possível. Portanto, precisam de ajuda para descobrir se é possível
//dividir a melancia da maneira que desejam. É importante destacar que cada um deles deve receber uma parte de peso
//positivo.
//
//A função deve retornar um valor booleano, indicando se é possível ou não dividir a melancia da forma desejada. Se o peso
//da melancia for menor ou igual a 0, a função deve retornar um erro.

package q4

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

package main

import (
	"fmt"
)

func DivideWatermelon(PesoMelancia int) (bool, error) {
	if PesoMelancia <= 0 {
		return false, fmt.Errorf("peso inválido")
	} else if PesoMelancia%2 != 0 {
		return false, fmt.Errorf("peso inválido")
	} else if PesoMelancia == 2 {
		return false, nil
	}
	return true, nil
}
func main() {

	possible, err := DivideWatermelon(10)
	fmt.Printf("Q1:\tpossible:%v, \terr:: %v\n", possible, err)
}

