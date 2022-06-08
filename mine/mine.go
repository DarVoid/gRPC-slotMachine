package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"
)

func main() {
	block := bloco{
		"Jorge",
		time.Now(),
		0,
	}
	blocos := make(chan bloco, 1)
	for i := 0; i < 100; i++ {
		go mine(block, blocos)

	}

	validateBloco(<-blocos)
}

func mine(block bloco, blocks chan bloco) {

	for {
		nBig, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			fmt.Println(err)
		}
		block.Nonce = nBig.Int64()
		fmt.Println(block)
		if validateBloco(block) {
			blocks <- block
		}
	}

}

func validateBloco(block bloco) bool {
	val, err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	hash := sha256.Sum256([]byte(val))
	fmt.Println(base64.StdEncoding.EncodeToString((hash[:])))
	return strings.HasPrefix(base64.StdEncoding.EncodeToString((hash[:])), "00")
}

type bloco struct {
	name  string    "json:Name"
	date  time.Time "json:Date"
	Nonce int64     "json:Nonce"
}
