package main

import "fmt"

type Card struct {
	tries      int
	isBlocked  bool
	maxTries   int
	password   string
}

func (card *Card) verifyTriesLimit() bool {
	remainingAttempts := card.maxTries - card.tries

	if remainingAttempts <= 0 {
		card.isBlocked = true
		fmt.Println("Você está bloqueado de tentar efetuar tentativas!")
		return true
	}

	fmt.Printf("Você ainda tem %d tentativas.\n", remainingAttempts)
	return false
}

func (card *Card) verifyPasswordInput(tentativa string) bool {
	if tentativa != card.password {
		card.tries++
		fmt.Println("\nSenha incorreta! Digite novamente.")
		return false
	}

	fmt.Println("Senha confirmada!")
	return true
}

func start(card *Card) {
	for !card.isBlocked {
		var passwordInput string
		fmt.Println("Insira a senha do seu cartão de crédito:")
		fmt.Scanln(&passwordInput)

		if card.verifyPasswordInput(passwordInput) {
			break
		}

		if card.verifyTriesLimit() {
			break
		}
	}
}

func main() {
	card := Card{
		password: "picos18",
		maxTries: 3,
	}

	start(&card)
}
