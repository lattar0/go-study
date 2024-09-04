package main

import "fmt"

type Card struct {
	tries       int
	isBlocked   bool
	isInserted  bool
	password    string
}

func (card *Card) insertCard() {
	card.isInserted = true
}

func (card *Card) verifyIfCardIsInserted() {
	if !card.isInserted {
		fmt.Println("Insira o cartão antes de digitar a senha!")
		return
	}
}

func (card *Card) verifyTriesLimit() {
	if card.tries > 2 {
		card.isBlocked = true
		fmt.Println("Você está bloqueado de tentar efetuar tentativas!")
	}
}

func (card *Card) verifyPasswordInput(tentativa string) {
	if tentativa != card.password {
		card.tries += 1
		fmt.Println("Senha incorreta! \nDigite novamente.")
		return
	}

	fmt.Println("Senha confirmada!")
}

func main() {
	card := Card{
		password: "picos18",
	}

	card.verifyIfCardIsInserted()
	card.insertCard()
	card.verifyTriesLimit()
	card.verifyPasswordInput("senha")
	card.verifyTriesLimit()
	card.verifyPasswordInput("senha")
	card.verifyTriesLimit()
	card.verifyPasswordInput("picos18")
}
