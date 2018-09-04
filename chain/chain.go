package chain

import (
	"github.com/pkg/errors"
	"time"
)

type Hash string
type Nonce int64

type Transaction struct {
	From   string
	To     string
	Amount int
}

func (t Transaction) String() string {
	return ""
}

type Block struct {
	Timestamp    time.Time
	Hash         Hash
	Transactions []Transaction
	Nonce        Nonce
	Previous     *Block
}

type BlockChain struct {
	Miner Miner
}

func NewKroppChain() BlockChain {
	return BlockChain{
		Miner: NoOpMiner{},
	}
}

func rewardForMining(identity string) Transaction {
	return Transaction{
		To:     identity,
		Amount: 50,
	}
}

func (b BlockChain) NewBlock(transactions []Transaction, identity string) *Block {
	return b.AddBlock(transactions, identity, nil)
}

func (b BlockChain) AddBlock(transactions []Transaction, identity string, chain *Block) *Block {
	next := Block{
		Timestamp:    time.Now(),
		Transactions: transactions,
		Previous:     chain,
	}

	next.Nonce = b.Miner.Mine(&next)

	next.Transactions = append(next.Transactions, rewardForMining(identity))

	next.Hash = Sha256(&next)

	return &next
}

func (b BlockChain) Transfer(from string, to string, amount int) Transaction {
	return Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
}

func (b BlockChain) VerifyBlock(chain *Block) error {
	curr := chain

	for curr != nil {
		if Sha256(curr) == curr.Hash {
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
