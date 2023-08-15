package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Pokemon struct {
	Id           int
	Nome         string
	VidaMax      int
	Vida         int
	Ataque       int
	RecuperaVida int
}

func Delay() {
	time.Sleep(1 * time.Second)
}

func main() {

	// POKEMONS ULTILIZÁVEIS

	Charizard := Pokemon{
		Id:           1,
		Nome:         "Charizard",
		Vida:         70,
		VidaMax:      70,
		Ataque:       12,
		RecuperaVida: 10,
	}

	Blastoise := Pokemon{
		Id:           2,
		Nome:         "Blastoise",
		Vida:         100,
		VidaMax:      100,
		Ataque:       8,
		RecuperaVida: 13,
	}

	Venusaur := Pokemon{
		Id:           3,
		Nome:         "Venusaur",
		Vida:         110,
		VidaMax:      110,
		Ataque:       9,
		RecuperaVida: 12,
	}

	//POKEMONS INIMIGOS

	Mew := Pokemon{
		Nome:         "Mew",
		Vida:         150,
		VidaMax:      150,
		Ataque:       10,
		RecuperaVida: 70,
	}

	Mewtwo := Pokemon{
		Nome:         "Mewtwo",
		Vida:         100,
		VidaMax:      100,
		Ataque:       20,
		RecuperaVida: 15,
	}
	Suicune := Pokemon{
		Nome:         "Suicune",
		Vida:         100,
		VidaMax:      100,
		Ataque:       15,
		RecuperaVida: 40,
	}

	// Aleatorizando o Pokemon Inimigo
	PokemonInimigo := rand.Intn(3) + 1

	//ESCOLHENDO O POKEMON

	fmt.Println("BEM VINDO A BATALHA POKEMÓN")
	Delay()

	fmt.Println("1-Charizard")
	fmt.Println("HP: 70, ATAQUE: 12, RECUPERAÇÃO DE VIDA:10")
	fmt.Println()
	Delay()

	fmt.Println("2-Blastoise")
	fmt.Println("HP: 100, ATAQUE: 8, RECUPERAÇÃO DE VIDA:13")
	fmt.Println()
	Delay()

	fmt.Println("3-Venusaur")
	fmt.Println("HP: 110, ATAQUE: 9, RECUPERAÇÃO DE VIDA:7")
	fmt.Println()
	Delay()

	fmt.Print("Escolha o seu pokemon:")

	var Escolher int

	fmt.Scan(&Escolher)

	//VALIDANDO A ESCOLHA

	for Escolher != Charizard.Id && Escolher != Blastoise.Id && Escolher != Venusaur.Id {
		fmt.Println("Escolha uma opção válida entre 1 a 3")
		fmt.Scan(&Escolher)
	}

	//MOSTRANDO A ESCOLHA DO USUÁRIO
	Delay()
	switch Escolher {
	case 1:
		fmt.Println("Você escolheu o", Charizard.Nome)
	case 2:
		fmt.Println("Você escolheu o", Blastoise.Nome)
	case 3:
		fmt.Println("Você escolheu o", Venusaur.Nome)

	}
	fmt.Println()

	//MOSTRANDO A ESCOLHA DO COMPUTADOR
	Delay()
	switch PokemonInimigo {
	case 1:
		fmt.Println("Seu adversário  o:", Mew.Nome)
		fmt.Print("HP: 150, ATAQUE: 10, RECUPERAÇÃO DE VIDA:70\n")
	case 2:
		fmt.Println("Seu adversário  o:", Mewtwo.Nome)
		fmt.Print("HP: 100, ATAQUE: 20, RECUPERAÇÃO DE VIDA:15\n")
	case 3:
		fmt.Println("Seu adversário  o:", Suicune.Nome)
		fmt.Print("HP: 100, ATAQUE: 15, RECUPERAÇÃO DE VIDA:40\n")

	}
	fmt.Println()

	var PokemonUsuario Pokemon

	switch Escolher {
	case 1:
		PokemonUsuario = Charizard
	case 2:
		PokemonUsuario = Blastoise
	case 3:
		PokemonUsuario = Venusaur
	}

	var PokemonPC Pokemon

	switch PokemonInimigo {
	case 1:
		PokemonPC = Mew
	case 2:
		PokemonPC = Mewtwo
	case 3:
		PokemonPC = Suicune
	}

	Delay()
	fmt.Println("A batalha vai começaaaaar")
	fmt.Println()

	Delay()
	fmt.Println(fmt.Sprintf("%s X %s", PokemonUsuario.Nome, PokemonPC.Nome))
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------")

	for PokemonPC.Vida > 0 && PokemonUsuario.Vida > 0 {

		fmt.Printf("VIDA DO SEU POKEMON: %d\n", PokemonUsuario.Vida)
		fmt.Println()
		Delay()

		Delay()
		fmt.Printf("VIDA DO POKEMON INIMIGO: %d\n", PokemonPC.Vida)
		fmt.Println()
		Delay()

		fmt.Println("1- Para atacar")
		Delay()
		fmt.Println()

		fmt.Println("2- Para Recuperar sua vida")
		Delay()
		fmt.Println()

		fmt.Print("Escolha uma opção:")
		fmt.Println()

		var Interacao int

		fmt.Scan(&Interacao)

		switch Interacao {

		//ATAQUE DO USUÁRO
		case 1:
			fmt.Println("Escolha o tipo de ataque")
			Delay()
			fmt.Println()

			fmt.Println("1-Ataque normal")
			fmt.Println()
			Delay()

			fmt.Println("2-Ataque Especial")
			fmt.Println()
			Delay()

			fmt.Print("Escolha uma opção de ataque:")
			fmt.Println()
			Delay()

			var Ataque int

			fmt.Scan(&Ataque)

			var UsuarioAtaca int
			if Ataque == 1 {
				UsuarioAtaca = PokemonUsuario.Ataque
			} else if Ataque == 2 {
				UsuarioAtaca = PokemonUsuario.Ataque + 20
			} else {
				fmt.Println("Você errou o ataque, na próxima escolha uma opção válida")
				fmt.Println()
				Delay()

			}

			PokemonPC.Vida -= UsuarioAtaca
			Delay()
			fmt.Printf("Voce atacou o %s com %d de dano\n ", PokemonPC.Nome, UsuarioAtaca)
			fmt.Println()
			Delay()

		case 2:
			//RECUPERAR VIDA DO POKEMON USUARIO
			PokemonUsuario.Vida += PokemonUsuario.RecuperaVida

			if PokemonUsuario.Vida > PokemonUsuario.VidaMax {
				Delay()

				fmt.Printf("Seu pokemon já está com vida máxima, perdeu o ataque\n")
				fmt.Println()
				Delay()

			} else {
				Delay()
				fmt.Printf("você recuperou %d de vida do seu pokemon\n", PokemonUsuario.RecuperaVida)
				fmt.Println()

			}
		default:
			// Opção inválida
			fmt.Println("Opção inválida. Você perdeu a rodada.")
			fmt.Println()
			Delay()
		}

		TurnoPC := rand.Intn(2) + 1
		if TurnoPC == 1 {
			// O computador ataca

			poderAtaque := PokemonPC.Ataque
			PokemonUsuario.Vida -= poderAtaque
			Delay()

			fmt.Printf("Seu adversário %s atacou com %d de dano.\n", PokemonPC.Nome, poderAtaque)
			fmt.Println()
			Delay()

		} else {
			// O computador recupera vida

			if PokemonPC.Vida >= PokemonPC.VidaMax {
				fmt.Printf("O pokemon  %s tentou recuperar  vida, mas já estava com a vida máxima\n", PokemonPC.Nome)
				fmt.Println()
				Delay()

			} else {
				fmt.Printf("Seu adversário %s recuperou %d de vida.\n", PokemonPC.Nome, PokemonPC.RecuperaVida)
				fmt.Println()
				Delay()

				if PokemonPC.Vida+PokemonPC.RecuperaVida > PokemonPC.VidaMax {
					PokemonPC.Vida = PokemonPC.VidaMax
				} else {
					PokemonPC.Vida += PokemonPC.RecuperaVida
				}
			}
			Delay()

		}

		if PokemonUsuario.Vida < 0 {
			fmt.Println(fmt.Sprintf("você foi derrotado pelo %s", PokemonPC.Nome))
		} else if PokemonPC.Vida < 0 {
			fmt.Println(fmt.Sprintf("você derrotou o %s", PokemonPC.Nome))
		}

	}
}
