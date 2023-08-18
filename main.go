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

func (pokemon Pokemon) ApresentarPokemon() {
	// ! Remover depois para testar
	if pokemon.Id == 0 {
		return
	}

	fmt.Println(fmt.Sprintf("%v - %v", pokemon.Id, pokemon.Nome))
	fmt.Println(fmt.Sprintf("HP: %v, ATAQUE: %v, RECUPERAÇÃO DE VIDA:%v", pokemon.Vida, pokemon.Ataque, pokemon.RecuperaVida))
	fmt.Println()
	Delay()
}

func (pokemon Pokemon) ApresentarDadosPokemon(isUser bool) {
	var tipoString = "pokemon adversário"

	if isUser {
		tipoString = "seu pokemon"
	}
	fmt.Println()
	fmt.Println(fmt.Sprintf("O %v é: %v", tipoString, pokemon.Nome))
	fmt.Print(fmt.Sprintf("HP: %v, ATAQUE: %v, RECUPERAÇÃO DE VIDA:%v\n", pokemon.Vida, pokemon.Ataque, pokemon.RecuperaVida))
}

func RegistrarAcao() (interacao, subacao int) {
	fmt.Println("1- Para atacar")
	Delay()
	fmt.Println()

	fmt.Println("2- Para Recuperar sua vida")
	Delay()
	fmt.Println()

	fmt.Print("Escolha uma opção:")
	fmt.Println()

	fmt.Scan(&interacao)

	switch interacao {
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

		fmt.Scan(&subacao)
	}

	return
}

func (pokemon *Pokemon) ExecutarAcao(pokemonInimigo *Pokemon, acao, subacao int) {
	if acao == 1 {
		switch subacao {
		case 1:
			pokemonInimigo.Vida -= pokemon.Ataque
			fmt.Printf("O pokemon %v atacou com %v de dano.\n", pokemon.Nome, pokemon.Ataque)
			fmt.Println()
			Delay()
		case 2:
			pokemonInimigo.Vida -= pokemon.Ataque + 15
			fmt.Printf("O pokemon %v super atacou com %v de dano.\n", pokemon.Nome, pokemon.Ataque+15)
			fmt.Println()
			Delay()
		}
	} else {
		if pokemon.Vida+pokemon.RecuperaVida <= pokemon.VidaMax {
			pokemon.Vida += pokemon.RecuperaVida
			fmt.Printf("o pokemon %v recuperou %v de vida.\n", pokemon.Nome, pokemon.Vida)
			fmt.Println()
			Delay()
		} else {
			fmt.Println("Você não pode recuperar mais vida, rodada perdida")
			fmt.Println()
			Delay()
		}
	}
}

func Delay() {
	time.Sleep(1 * time.Second)
}

func MontarPokemons() (listaPokemons []Pokemon) {
	charizard := Pokemon{
		Id:           1,
		Nome:         "Charizard",
		Vida:         70,
		VidaMax:      70,
		Ataque:       12,
		RecuperaVida: 10,
	}

	blastoise := Pokemon{
		Id:           2,
		Nome:         "Blastoise",
		Vida:         100,
		VidaMax:      100,
		Ataque:       8,
		RecuperaVida: 13,
	}

	venusaur := Pokemon{
		Id:           3,
		Nome:         "Venusaur",
		Vida:         110,
		VidaMax:      110,
		Ataque:       9,
		RecuperaVida: 12,
	}

	mew := Pokemon{
		Nome:         "Mew",
		Vida:         150,
		VidaMax:      150,
		Ataque:       10,
		RecuperaVida: 70,
	}

	mewtwo := Pokemon{
		Nome:         "Mewtwo",
		Vida:         100,
		VidaMax:      100,
		Ataque:       20,
		RecuperaVida: 15,
	}

	suicune := Pokemon{
		Nome:         "Suicune",
		Vida:         100,
		VidaMax:      100,
		Ataque:       15,
		RecuperaVida: 40,
	}

	listaPokemons = make([]Pokemon, 0)

	listaPokemons = append(listaPokemons, charizard, blastoise, venusaur, mew, mewtwo, suicune)
	return
}

func PokemonsValidosIDs(listaPokemons []Pokemon) (res []int) {
	for _, pokemon := range listaPokemons {
		if pokemon.Id != 0 {
			res = append(res, pokemon.Id)
		}
	}

	return
}

func BuscarPokemon(idPokemon int, listaPokemons []Pokemon) (res Pokemon) {
	for _, pokemon := range listaPokemons {
		if pokemon.Id == idPokemon {
			return pokemon
		}
	}
	return
}

func IsInInt(base int, values ...int) (isIn bool) {
	for _, value := range values {
		if value == base {
			return true
		}
	}
	return
}

func IniciarBatalha() {
	var pokemonUsuario, pokemonPC Pokemon
	pokemons := MontarPokemons()

	// Aleatorizando o Pokemon Inimigo
	pokemonPC = pokemons[rand.Intn(3)+3]

	//ESCOLHENDO O POKEMON

	fmt.Println("BEM VINDO A BATALHA POKEMÓN")
	Delay()

	for _, pokemon := range pokemons {
		pokemon.ApresentarPokemon()
	}

	fmt.Println("Escolha o seu pokemon:")
	var idPokemonUsuario int
	fmt.Scan(&idPokemonUsuario)

	//VALIDANDO A ESCOLHA
	for !IsInInt(idPokemonUsuario, PokemonsValidosIDs(pokemons)...) {
		fmt.Println("Escolha uma opção válida")
		fmt.Scan(&idPokemonUsuario)
	}

	//MOSTRANDO A ESCOLHA DO USUÁRIO
	pokemonUsuario = BuscarPokemon(idPokemonUsuario, pokemons)
	pokemonUsuario.ApresentarDadosPokemon(true)
	Delay()

	//MOSTRANDO A ESCOLHA DO COMPUTADOR
	pokemonPC.ApresentarDadosPokemon(false)

	Delay()
	fmt.Println("A batalha vai começaaaaar")
	fmt.Println()

	Delay()
	fmt.Println(fmt.Sprintf("%s X %s", pokemonUsuario.Nome, pokemonPC.Nome))
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------")

	for pokemonPC.Vida > 0 && pokemonUsuario.Vida > 0 {
		fmt.Printf("VIDA DO SEU POKEMON: %d\n", pokemonUsuario.Vida)
		fmt.Println()
		Delay()

		Delay()
		fmt.Printf("VIDA DO POKEMON INIMIGO: %d\n", pokemonPC.Vida)
		fmt.Println()
		Delay()

		// Responsável por registrar a ação escolhida pelo usuário
		acaoUsuario, subacaoUsuario := RegistrarAcao()

		acaoPC, subacaoPC := rand.Intn(2)+1, rand.Intn(2)+1

		// Responsável por executar uma ação definida anteriormente
		pokemonUsuario.ExecutarAcao(&pokemonPC, acaoUsuario, subacaoUsuario)
		// Responsável por executar uma ação definida de forma automágica
		pokemonPC.ExecutarAcao(&pokemonUsuario, acaoPC, subacaoPC)

		if pokemonUsuario.Vida < 0 {
			fmt.Println(fmt.Sprintf("você foi derrotado pelo %s", pokemonPC.Nome))
		} else if pokemonPC.Vida < 0 {
			fmt.Println(fmt.Sprintf("você derrotou o %s", pokemonPC.Nome))
		}

	}
}

func main() {
	IniciarBatalha()
}
