package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops - For")
	// i := 0
	// for i < 10 {
	// 	fmt.Println("incrementando I")
	// 	time.Sleep(time.Second)
	// 	i++
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("incrementando J", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := []string{"Diego", "Dayana", "Douglas"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Diego",
		"sobrenome": "Simon",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
