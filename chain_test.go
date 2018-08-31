package kroppchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyValid(t *testing.T) {
	chain := NewKroppChain()

	root := chain.AddBlock(Metadata("root"), nil)

	root = chain.AddBlock(Metadata("next"), root)

	err := chain.VerifyBlock(root)

	assert.NoError(t, err)
}

func TestChainLength(t *testing.T) {
	chain := NewKroppChain()

	root := chain.AddBlock(Metadata("root"), nil)

	root = chain.AddBlock(Metadata("next"), root)

	chainLength := LengthOf(root)

	assert.Equal(t, 2, chainLength)
}

func TestVerifyNotValid(t *testing.T) {
	chain := NewKroppChain()

	root := chain.AddBlock(Metadata("root"), nil)

	root = chain.AddBlock(Metadata("next"), root)

	root.Previous.Hash = Hash("0" + root.Previous.Hash)

	err := chain.VerifyBlock(root)

	assert.Error(t, err)
}

func TestReconcileMultiChains(t *testing.T) {
	chain := NewKroppChain()

	root1 := chain.AddBlock(Metadata("root1"), nil)
	root1 = chain.AddBlock(Metadata("next"), root1)

	root2 := chain.AddBlock(Metadata("root1"), nil)
	root2 = chain.AddBlock(Metadata("next"), root2)
	root2 = chain.AddBlock(Metadata("next again"), root2)

	resultingChain := Reconcile(root1, root2)

	assert.Equal(t, root2, resultingChain)
}
