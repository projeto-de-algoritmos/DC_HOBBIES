package main

import (
	"DC/HOBBIES/dchobbies"
	"fmt"
)

func main() {

	fmt.Println(`
			Hobbies influenciam o QI ?

Segundo a associação de pesquisa e achismo do google, um estudo sugere 
que o nível de inteligência de uma pessoa pode ser definido com a ordem dos 
hobbies que ela prefere segundo a lista abaixo. Onde quanto menor o nível, maior
é a inteligência.
				`)

	fmt.Println(`
1 - Programação
2 - Esportes
3 - Podcasts
4 - Séries e Filmes
5 - Video games
6 - Ler
7 - Escrever
				`)

	var one_player []int

	fmt.Print("Jogador escolha a ordem que você prefere dos hobbies:")
	fmt.Print("\n")

	for i := 0; i < 7; i++ {
		var number int
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			fmt.Println(err)
		}

		one_player = append(one_player, number)
	}

	_, totalInv := dchobbies.CountSliceInversions(one_player)

	fmt.Printf("Seu nível de inteligência é %d!\n", totalInv)
}
