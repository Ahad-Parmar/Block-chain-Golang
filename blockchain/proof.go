// Proof of work algoritm
// below are steps to see how it works in blockchain

// taking a data from the block
// creating a counter (nonce) which starts at 0
// creating a hash of the data plus the counter
// checking the hash to see if it meets a below set of requirements

// requirements:
// the first few bytes must contain 0s

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block   // blocks in blockchain
	Target *big.Int // no. that represent the reuirement of difficulty
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte { // nonce is the counter. because we have to combine the data from the block with counter
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash, // i have taken prevHash and Data like in the derivedHash function
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{}, //then i have combine the another byte struct using bytes join function to create cohesive set of byte
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce) //preparing the data
		hash = sha256.Sum256(data)  // taking all that data and hash it into a sha256 format
		fmt.Printf("\r %x", hash)
		intHash.SetBytes(hash[:]) // to change my hash in BigInt

		if intHash.Cmp(pow.Target) == -1 { // for comparing the Proof Of Work Target
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool { // validation method
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num) // binary.BigEndian - signifies how i want to organize my bytes
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
