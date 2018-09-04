package chain

type Miner interface {
	Mine(block *Block) Nonce
	Verify(nonce string) error
}

type NoOpMiner struct{}

func (x NoOpMiner) Mine(block *Block) Nonce {
	return Nonce(0)
}

func (x NoOpMiner) Verify(nonce string) error {
	return nil
}

