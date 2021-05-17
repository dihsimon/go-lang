package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Diego",
		"sobrenome": "Simon",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Dayana",
			"Ultimo":   "Rizzato",
		},
		"curso": {
			"nome":   "Engenharia2",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "gemeos",
	}
	fmt.Println(usuario2)
}
