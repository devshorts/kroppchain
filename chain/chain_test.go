package chain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const identity = "test"

func FakeTransaction() []Transaction {
	return []Transaction{{}}
}

func TestVerifyValid(t *testing.T) {
	chain := NewKroppChain()

	root := chain.NewBlock(FakeTransaction(), identity)

	root = chain.AddBlock(FakeTransaction(), identity, root)

	err := chain.VerifyBlock(root)

	assert.NoError(t, err)
}

func TestChainLength(t *testing.T) {
	chain := NewKroppChain()

	root := chain.NewBlock(FakeTransaction(), identity)

	root = chain.AddBlock(FakeTransaction(), identity, root)

	chainLength := LengthOf(root)

	assert.Equal(t, 2, chainLength)
}

func TestVerifyNotValid(t *testing.T) {
	chain := NewKroppChain()

	root := chain.NewBlock(FakeTransaction(), identity)

	root = chain.AddBlock(FakeTransaction(), identity, root)

	root.Previous.Hash = Hash("0" + root.Previous.Hash)

	err := chain.VerifyBlock(root)

	assert.Error(t, err)
}

func TestReconcileMultiChains(t *testing.T) {
	chain := NewKroppChain()

	root1 := chain.NewBlock(FakeTransaction(), identity)
	root1 = chain.AddBlock(FakeTransaction(), identity, root1)

	root2 := chain.NewBlock(FakeTransaction(), identity)
	root2 = chain.AddBlock(FakeTransaction(), identity, root2)
	root2 = chain.AddBlock(FakeTransaction(), identity, root2)

	resultingChain := Reconcile(root1, root2)

	assert.Equal(t, root2, resultingChain)
}
