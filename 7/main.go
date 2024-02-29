package main

import "fmt"

func main() {
	salarios := map[string]int{"Vinicius": 1000, "João": 2000, "Jr": 5000}
	fmt.Println(salarios["Vinicius"])
	delete(salarios, "Jr")
	salarios["Maria"] = 1100
	println(salarios)

	precos := make(map[string]int)
	precos2 := map[string]int{}
	precos2["Head"] = 1000
	precos["Head"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
