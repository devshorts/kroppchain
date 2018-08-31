package kroppchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"io"
)

type Hash string
type Metadata string

type Block struct {
	Hash     Hash
	Proof    ProofOfWork
	Metadata Metadata
	Previous *Block
}

type ProofOfWork struct {
	Nonce string
}

func RandomWork(metadata Metadata) ProofOfWork {
	return ProofOfWork{Nonce: string(metadata)}
}

type BlockChain struct {
	Miner func(Metadata) ProofOfWork
}

func NewKroppChain() BlockChain {
	return BlockChain{Miner: RandomWork}
}

func (b BlockChain) AddBlock(metadata Metadata, chain *Block) *Block {
	next := Block{
		Proof:    b.Miner(metadata),
		Metadata: metadata,
		Previous: chain,
	}

	next.Hash = hash(&next)

	return &next
}

func (b BlockChain) VerifyBlock(chain *Block) error {
	curr := chain

	for curr != nil {
		if hash(curr) == curr.Hash {
			curr = curr.Previous
		} else {
			return errors.New("chain hash invalid")
		}
	}

	return nil
}

func Reconcile(chain1 *Block, chain2 *Block) *Block {
	if LengthOf(chain1) >= LengthOf(chain2) {
		return chain1
	}

	return chain2
}

func LengthOf(chain *Block) int {
	length := 0
	for chain != nil {
		chain = chain.Previous
		length += 1
	}

	return length
}

func hash(block *Block) Hash {
	s := sha256.New()

	byteWriter := bytes.NewBuffer([]byte{})

	byteWriter.WriteString(string(block.Proof.Nonce))
	byteWriter.WriteString(string(block.Metadata))

	if block.Previous != nil {
		byteWriter.WriteString(string(block.Previous.Hash))
	}

	io.Copy(s, byteWriter)

	return Hash(fmt.Sprintf("%x", s.Sum(nil)))
}
