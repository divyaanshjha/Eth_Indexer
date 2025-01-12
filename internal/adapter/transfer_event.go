package adapter

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

// TransferEvent represents the event data of an ERC-20 transfer
type TransferEvent struct {
	From  common.Address // Address sending the tokens
	To    common.Address // Address receiving the tokens
	Value *big.Int       // Amount of tokens transferred
	TxHash common.Hash   // The transaction hash that this event belongs to
}