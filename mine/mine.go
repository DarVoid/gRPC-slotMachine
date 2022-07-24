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
	//creates a block
	block := bloco{
		"Jorge",
		time.Now(),
		0,
	}
	//creates buffered channel only takes 1 value will serve as sync mechanism
	blocos := make(chan bloco, 1)

	//initializes go routines each mining their own random solution
	for i := 0; i < 100; i++ {
		go mine(block, blocos)

	}
	//this line blocks until 1 result is found
	validateBloco(<-blocos)
}

//mines block
func mine(block bloco, blocks chan bloco) {
	blockNew := bloco{
		Name:  block.Name,
		Date:  block.Date,
		Nonce: 0,
	}
	for {
		nBig, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			fmt.Println(err)
		}
		blockNew.Nonce = nBig.Int64()
		blocojson, err := json.Marshal(blockNew)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(blocojson))
		if validateBloco(blockNew) {

			blocks <- blockNew
		}
	}

}

//validates if block obeys dificulty
func validateBloco(block bloco) bool {

	blocojson, err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(blocojson))
	hash := sha256.Sum256([]byte(blocojson))
	fmt.Println(base64.StdEncoding.EncodeToString((hash[:]))) //just to be human readable
	return strings.HasPrefix(base64.StdEncoding.EncodeToString((hash[:])), "000")
}

type bloco struct {
	Name  string    "json:Name"
	Date  time.Time "json:Date"
	Nonce int64     "json:Nonce"
}
