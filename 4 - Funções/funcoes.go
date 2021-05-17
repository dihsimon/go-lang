package main

import "fmt"

func main() {
	fmt.Println(somar(2, 2))

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	fmt.Println(f("texto da função f"))

	resSoma, resSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resSoma, resSubtracao)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
