package main

import (
	"fmt"
)

func main() {

	var (
		wiki = []string{"ane", "ama", "mar", "claro", "calma", "credo", "clama"}
	)

	len_word := 5
	p_letter := 3
	letter := "a" //transferir para valor unicode int32

	for i := 0; i < len(wiki); i++ {
		// verifica se a string tem o tamanho exigido
		if len(wiki[i]) == len_word {
			// itera sobre a string
			for x := 0; x < len(wiki[i]); x++ {
				//verifica se x esta na posicao da letra, e se o código UTF8 é o mesmo
				if x == p_letter-1 && wiki[i][x] == letter[0] {
					fmt.Println(wiki[i])
				}
			}
		}
	}
}
