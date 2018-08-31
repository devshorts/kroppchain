package kroppchain

import (
	"math/rand"
	"strconv"
)

type Miner interface {
	Mine() ProofOfWork
	Verify(nonce string) error
}

type NoOpMiner struct{}

func (x NoOpMiner) Mine() ProofOfWork {
	return ProofOfWork{Nonce: strconv.Itoa(rand.Int())}
}

func (x NoOpMiner) Verify(nonce string) error {
	return nil
}

