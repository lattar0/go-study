package main

import "fmt"

type Card struct {
	tries      int
	isBlocked  bool
	maxTries   int
	password   string
}

func (card *Card) verifyTriesLimit() bool {
	maxTries := 3

	remainingAttempts := maxTries - card.tries

	if card.tries > 2 {
		card.isBlocked = true
		fmt.Println("Você está bloqueado de tentar efetuar tentativas!")
		return true
	}

	fmt.Printf("Você ainda tem: %d tentativas.\n", remainingAttempts)

	return false
}

func (card *Card) verifyPasswordInput(tentativa string) {
	if tentativa != card.password {
		card.tries += 1
		fmt.Println("\nSenha incorreta! \nDigite novamente.")
		return
	}

	fmt.Println("Senha confirmada!")
}

func start(card *Card) {
	for true {
		passwordTrie := ""
		fmt.Println("Insira a senha do seu cartão de crédito:")
		fmt.Scanln(&passwordTrie)
		
		card.verifyPasswordInput(passwordTrie)
		if card.verifyTriesLimit() {
			break
		}
	}
}

func main() {
	card := Card{
		password: "picos18",
	}

	start(&card)
}
