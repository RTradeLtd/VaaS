package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

type EthereumGenerator struct {
	SearchPrefix     string
	RunTimeInSeconds int64
}

type Account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

// InitializeEthereumGenerator is used to initialize our ethereum address generation service
func InitializeEthereumGenerator(searchPrefix string, runTimeInSeconds int64) *EthereumGenerator {
	return &EthereumGenerator{
		SearchPrefix:     searchPrefix,
		RunTimeInSeconds: runTimeInSeconds}
}

// Run is used to execute our ethereum address generation service
func (eg *EthereumGenerator) Run(c *gin.Context) {
	for {
		acct, err := eg.CreateAccount()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		matched := eg.Match(acct)
		if matched {
			encodedKey := fmt.Sprintf("0x%s", hex.EncodeToString(crypto.FromECDSA(acct.PrivateKey)))
			address := acct.Address.String()
			c.JSON(http.StatusOK, gin.H{
				"private_key": encodedKey,
				"address":     address,
			})
			return
		}
	}
}

// CreateAccount is used to create an ethereum account
func (eg *EthereumGenerator) CreateAccount() (*Account, error) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(pk.PublicKey)
	return &Account{
		PrivateKey: pk,
		Address:    address,
	}, nil
}

// Match is used to check if the provided ethereum account matches our search
func (eg *EthereumGenerator) Match(acct *Account) bool {
	charactersToMatch := len(eg.SearchPrefix)
	trimmedAddress := strings.TrimPrefix(acct.Address.String(), "0x")
	fmt.Println(trimmedAddress)
	fmt.Println("search prefix", eg.SearchPrefix)
	partToMatch := trimmedAddress[0:charactersToMatch]
	fmt.Println("part to match", partToMatch)
	if partToMatch == eg.SearchPrefix {
		return true
	}
	return false
}
