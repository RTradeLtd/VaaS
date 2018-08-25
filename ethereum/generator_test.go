package ethereum_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/VaaS/ethereum"
)

func TestEthereumGenerator(t *testing.T) {
	generator := ethereum.InitializeEthereumGenerator(
		"D",
		10000000,
	)
	for {
		acct, err := generator.CreateAccount()
		if err != nil {
			t.Fatal(err)
		}
		matched := generator.Match(acct)
		if matched {
			fmt.Println("Match found!")
			fmt.Println("Address: ", acct.Address.String())
			return
		}
	}
}
